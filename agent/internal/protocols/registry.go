package protocols

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Registry manages a set of Protocol adapters, providing the discover→load→action pipeline.
// It mirrors Moss's Registry class: protocols are registered, then queried/actioned by name.
type Registry struct {
	protocols map[string]Protocol
}

// NewRegistry creates an empty Registry.
func NewRegistry() *Registry {
	return &Registry{protocols: make(map[string]Protocol)}
}

// Register adds a protocol to the registry. Rejects duplicates.
func (r *Registry) Register(p Protocol) error {
	name := p.Name()
	if _, exists := r.protocols[name]; exists {
		return fmt.Errorf("protocol %q is already registered", name)
	}
	r.protocols[name] = p
	return nil
}

// Protocol returns a registered protocol by name.
func (r *Registry) Protocol(name string) (Protocol, error) {
	p, ok := r.protocols[name]
	if !ok {
		names := make([]string, 0, len(r.protocols))
		for n := range r.protocols {
			names = append(names, n)
		}
		return nil, fmt.Errorf("unknown protocol %q (registered: %s)", name, strings.Join(names, ", "))
	}
	return p, nil
}

// List returns the names of all registered protocols.
func (r *Registry) List() []string {
	names := make([]string, 0, len(r.protocols))
	for n := range r.protocols {
		names = append(names, n)
	}
	return names
}

// Discover returns coordinates for all registered protocols, filtered by the optional criteria.
func (r *Registry) Discover(filter DiscoverFilter) []Coordinate {
	var found []Coordinate
	for _, p := range r.protocols {
		if filter.Protocol != "" && filter.Protocol != p.Name() {
			continue
		}
		if filter.Category != "" && filter.Category != p.Category() {
			continue
		}
		for _, c := range p.Discover() {
			if filter.Verb != "" && filter.Verb != c.Verb {
				continue
			}
			found = append(found, c)
		}
	}
	return found
}

// Load returns stubs for the requested (protocol, method) pairs.
func (r *Registry) Load(items []struct{ Protocol, Method string }) []Stub {
	var stubs []Stub
	for _, item := range items {
		p, err := r.Protocol(item.Protocol)
		if err != nil {
			continue
		}
		stub, err := p.Load(item.Method)
		if err != nil {
			continue
		}
		stubs = append(stubs, *stub)
	}
	return stubs
}

// Action executes a query or builds a capability on the named protocol.
func (r *Registry) Action(protocol, method string, account common.Address, params map[string]interface{}) (*ActionNode, error) {
	p, err := r.Protocol(protocol)
	if err != nil {
		return nil, err
	}
	return p.Action(method, account, params)
}
