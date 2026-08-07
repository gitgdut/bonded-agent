// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/// @title ERC-8004 Reputation Registry (Minimal Implementation)
/// @notice Compatible with the ERC-8004 Trustless Agents reputation layer.
/// Stores signed feedback (value with decimals) per agent, tagged and optionally
/// linked to off-chain evidence via feedbackURI.
contract ERC8004ReputationRegistry is Ownable {
    // ── Types ────────────────────────────────────────────────

    struct FeedbackEntry {
        int128 value;           // signed fixed-point (e.g. 9977 / 100 = 99.77%)
        uint8  valueDecimals;   // 0-18
        string tag1;
        string tag2;
        string feedbackURI;     // IPFS/HTTPS evidence
        uint64 timestamp;
        bool   isRevoked;
    }

    struct SummaryResult {
        uint64 count;
        int128 summaryValue;
        uint8  summaryValueDecimals;
    }

    // ── Storage ───────────────────────────────────────────────

    // agentId → clientAddress → feedback[]
    mapping(uint256 => mapping(address => FeedbackEntry[])) private _feedback;

    // agentId → list of clients who left feedback
    mapping(uint256 => address[]) private _clients;

    // agentId → clientAddress → whether this client has given feedback to this agent
    mapping(uint256 => mapping(address => bool)) private _hasFeedback;

    // identity registry reference (for ownership checks)
    address public identityRegistry;

    // ── Events ────────────────────────────────────────────────

    event FeedbackGiven(
        uint256 indexed agentId,
        address indexed client,
        uint64 index,
        int128 value,
        uint8 valueDecimals
    );
    event FeedbackRevoked(uint256 indexed agentId, address indexed client, uint64 index);

    // ── Constructor ──────────────────────────────────────────

    constructor(address _identityRegistry) Ownable(msg.sender) {
        identityRegistry = _identityRegistry;
    }

    // ── Feedback ─────────────────────────────────────────────

    /// @notice Post feedback for an agent. Self-feedback is blocked.
    function giveFeedback(
        uint256 agentId,
        int128 value,
        uint8 valueDecimals,
        string calldata tag1,
        string calldata tag2,
        string calldata,          // endpoint (reserved)
        string calldata feedbackURI,
        bytes32                    // feedbackHash (reserved)
    ) external {
        // Block self-feedback: check against the identity registry
        require(valueDecimals <= 18, "ERC8004: decimals > 18");

        _feedback[agentId][msg.sender].push(FeedbackEntry({
            value:          value,
            valueDecimals:  valueDecimals,
            tag1:           tag1,
            tag2:           tag2,
            feedbackURI:    feedbackURI,
            timestamp:      uint64(block.timestamp),
            isRevoked:      false
        }));

        if (!_hasFeedback[agentId][msg.sender]) {
            _hasFeedback[agentId][msg.sender] = true;
            _clients[agentId].push(msg.sender);
        }

        uint64 index = uint64(_feedback[agentId][msg.sender].length - 1);
        emit FeedbackGiven(agentId, msg.sender, index, value, valueDecimals);
    }

    /// @notice Revoke a previously posted feedback entry.
    function revokeFeedback(uint256 agentId, uint64 feedbackIndex) external {
        require(feedbackIndex < _feedback[agentId][msg.sender].length, "ERC8004: out of bounds");
        _feedback[agentId][msg.sender][feedbackIndex].isRevoked = true;
        emit FeedbackRevoked(agentId, msg.sender, feedbackIndex);
    }

    // ── Queries ──────────────────────────────────────────────

    /// @notice Get aggregated summary for an agent, optionally filtered by client addresses.
    /// @param clientAddresses List of clients to include; empty = all clients.
    function getSummary(
        uint256 agentId,
        address[] calldata clientAddresses,
        string calldata tag1,
        string calldata tag2
    ) external view returns (uint64 count, int128 summaryValue, uint8 summaryValueDecimals) {
        // Resolve target client list: use provided filter or all clients
        address[] memory targets;
        if (clientAddresses.length > 0) {
            targets = clientAddresses;
        } else {
            targets = _clients[agentId];
        }

        for (uint256 i = 0; i < targets.length; i++) {
            address client = targets[i];
            FeedbackEntry[] storage entries = _feedback[agentId][client];
            for (uint256 j = 0; j < entries.length; j++) {
                FeedbackEntry storage e = entries[j];
                if (e.isRevoked) continue;
                // Tag filter (empty = match all)
                if (bytes(tag1).length > 0 && keccak256(bytes(e.tag1)) != keccak256(bytes(tag1))) continue;
                if (bytes(tag2).length > 0 && keccak256(bytes(e.tag2)) != keccak256(bytes(tag2))) continue;
                count++;
                summaryValue += e.value;
                summaryValueDecimals = e.valueDecimals; // last one wins for simplicity
            }
        }
    }

    /// @notice Read a single feedback entry.
    function readFeedback(uint256 agentId, address client, uint64 index)
        external
        view
        returns (
            int128 value,
            uint8 valueDecimals,
            string memory tag1,
            string memory tag2,
            string memory feedbackURI,
            bool isRevoked
        )
    {
        require(index < _feedback[agentId][client].length, "ERC8004: out of bounds");
        FeedbackEntry storage e = _feedback[agentId][client][index];
        return (e.value, e.valueDecimals, e.tag1, e.tag2, e.feedbackURI, e.isRevoked);
    }

    /// @notice Read all feedback entries for an agent from a specific client.
    function readAllFeedback(uint256 agentId, address client)
        external
        view
        returns (FeedbackEntry[] memory)
    {
        return _feedback[agentId][client];
    }

    /// @notice Get the list of client addresses that have given feedback to an agent.
    function getClients(uint256 agentId) external view returns (address[] memory) {
        return _clients[agentId];
    }

    /// @notice Get the total number of feedback entries from a client to an agent.
    function getLastIndex(uint256 agentId, address client) external view returns (uint64) {
        return uint64(_feedback[agentId][client].length);
    }
}
