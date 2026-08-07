// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC721URIStorage, ERC721} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {IERC1271} from "@openzeppelin/contracts/interfaces/IERC1271.sol";

/// @title ERC-8004 Identity Registry (Minimal Implementation)
/// @notice Compatible with the ERC-8004 Trustless Agents identity layer.
/// Each agent/operator mints an ERC-721 NFT as its on-chain passport.
contract ERC8004IdentityRegistry is ERC721URIStorage, Ownable {
    using ECDSA for bytes32;

    // ── Storage ──────────────────────────────────────────────

    uint256 private _nextAgentId = 1;

    // agentId → agentWallet (optional receiving wallet for the agent)
    mapping(uint256 => address) private _agentWallets;

    // agentId → metadata key → value
    mapping(uint256 => mapping(string => bytes)) private _metadata;

    // ── Events ────────────────────────────────────────────────

    event AgentRegistered(uint256 indexed agentId, address indexed owner, string agentURI);
    event AgentWalletSet(uint256 indexed agentId, address indexed newWallet);
    event AgentWalletUnset(uint256 indexed agentId);
    event MetadataUpdated(uint256 indexed agentId, string key, bytes value);

    // ── Constructor ──────────────────────────────────────────

    constructor() ERC721("ERC-8004 Identity", "E8004-ID") Ownable(msg.sender) {}

    // ── Registration ─────────────────────────────────────────

    /// @notice Register a new agent with metadata URI (e.g. IPFS agent card).
    /// @return agentId The minted NFT token ID.
    function register(string calldata agentURI) external returns (uint256 agentId) {
        agentId = _nextAgentId++;
        _safeMint(msg.sender, agentId);
        _setTokenURI(agentId, agentURI);

        emit AgentRegistered(agentId, msg.sender, agentURI);
    }

    /// @notice Register a new agent with metadata URI and initial agent wallet.
    function register(string calldata agentURI, address agentWallet) external returns (uint256 agentId) {
        // Inline the registration logic — Solidity can't call overloads of the same name internally
        agentId = _nextAgentId++;
        _safeMint(msg.sender, agentId);
        _setTokenURI(agentId, agentURI);
        emit AgentRegistered(agentId, msg.sender, agentURI);
        _agentWallets[agentId] = agentWallet;
        emit AgentWalletSet(agentId, agentWallet);
    }

    // ── URI Management ───────────────────────────────────────

    /// @notice Update the agent's metadata URI.
    function setAgentURI(uint256 agentId, string calldata newURI) external {
        _requireOwnership(agentId);
        _setTokenURI(agentId, newURI);
    }

    // ── Agent Wallet ─────────────────────────────────────────

    /// @notice Set the agent's receiving wallet. Proof of ownership via EIP-712 or EIP-1271.
    /// @param agentId The agent's NFT ID.
    /// @param newWallet The new wallet address.
    /// @param deadline Expiry timestamp for the signature.
    /// @param signature EIP-712 signature from the new wallet proving ownership.
    function setAgentWallet(
        uint256 agentId,
        address newWallet,
        uint256 deadline,
        bytes calldata signature
    ) external {
        _requireOwnership(agentId);

        require(deadline >= block.timestamp, "ERC8004: signature expired");

        // Verify the new wallet signed an EIP-712 proof of control
        bytes32 structHash = keccak256(
            abi.encode(
                keccak256("SetAgentWallet(uint256 agentId,address newWallet,uint256 deadline)"),
                agentId,
                newWallet,
                deadline
            )
        );

        bytes32 domainSeparator = keccak256(
            abi.encode(
                keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
                keccak256(bytes("ERC-8004 Identity Registry")),
                keccak256(bytes("1")),
                block.chainid,
                address(this)
            )
        );

        bytes32 digest = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
        address signer = ECDSA.recover(digest, signature);

        require(signer == newWallet || _isValidERC1271Signature(newWallet, digest, signature),
                "ERC8004: invalid wallet signature");

        _agentWallets[agentId] = newWallet;
        emit AgentWalletSet(agentId, newWallet);
    }

    /// @notice Remove the agent's linked wallet.
    function unsetAgentWallet(uint256 agentId) external {
        _requireOwnership(agentId);
        delete _agentWallets[agentId];
        emit AgentWalletUnset(agentId);
    }

    /// @notice Get the agent's linked wallet (0x0 if none).
    function getAgentWallet(uint256 agentId) external view returns (address) {
        return _agentWallets[agentId];
    }

    // ── On-Chain Metadata ───────────────────────────────────

    /// @notice Set a key-value metadata entry for the agent.
    function setMetadata(uint256 agentId, string calldata key, bytes calldata value) external {
        _requireOwnership(agentId);
        _metadata[agentId][key] = value;
        emit MetadataUpdated(agentId, key, value);
    }

    /// @notice Get a metadata value by key.
    function getMetadata(uint256 agentId, string calldata key) external view returns (bytes memory) {
        return _metadata[agentId][key];
    }

    /// @notice Returns the current token counter (next agentId to be minted).
    function nextAgentId() external view returns (uint256) {
        return _nextAgentId;
    }

    // ── ERC-721 Overrides ───────────────────────────────────

    /// @dev Blocks transfers unless the recipient explicitly accepts.
    /// In ERC-8004, ownership represents the agent's controller — transfers are allowed
    /// but the agentWallet is cleared on transfer.
    function _update(address to, uint256 tokenId, address auth)
        internal
        override
        returns (address)
    {
        address from = _ownerOf(tokenId);
        // Clear agent wallet on transfer
        if (from != address(0) && to != address(0)) {
            delete _agentWallets[tokenId];
        }
        return super._update(to, tokenId, auth);
    }

    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        return super.tokenURI(tokenId);
    }

    // ── Helpers ──────────────────────────────────────────────

    function _requireOwnership(uint256 agentId) internal view {
        require(ownerOf(agentId) == msg.sender, "ERC8004: not agent owner");
    }

    function _isValidERC1271Signature(address wallet, bytes32 digest, bytes memory signature)
        internal
        view
        returns (bool)
    {
        if (wallet.code.length == 0) return false;
        try IERC1271(wallet).isValidSignature(digest, signature) returns (bytes4 magic) {
            return magic == IERC1271.isValidSignature.selector;
        } catch {
            return false;
        }
    }
}
