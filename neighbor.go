// Package neighbor enables network neighbor detection using operating system
// specific facilities.
package neighbor

import (
	"bytes"
	"errors"
	"net"
	"syscall"
)

//go:generate stringer -output=string.go -type=Family,Flags,State

// Family is a network address family, derived from operating system-specific
// constants.
type Family uint8

// Family constants which can be used to filter neighbors using various
// network address families.
const (
	FamilyUnspec Family = syscall.AF_UNSPEC
	FamilyIPv4   Family = syscall.AF_INET
	FamilyIPv6   Family = syscall.AF_INET6
)

// Flags is a set of flags indicating if a neighbor is a proxy, router, etc.
type Flags uint8

// Flags constants which can be used to determine special properties of a
// neighbor.
const (
	FlagsUse    Flags = 0x01
	FlagsSelf   Flags = 0x02
	FlagsMaster Flags = 0x04
	FlagsProxy  Flags = 0x08
	FlagsRouter Flags = 0x80
)

// State is a neighbor's state, indicating if the neighbor is reachable,
// unreachable, currently being contacted, etc.
type State uint16

// State constants which can be used to determine a neighbor's state.
const (
	StateNone       State = 0x00
	StateIncomplete State = 0x01
	StateReachable  State = 0x02
	StateStale      State = 0x04
	StateDelay      State = 0x08
	StateProbe      State = 0x10
	StateFailed     State = 0x20
	StateNoARP      State = 0x40
	StatePermanent  State = 0x80
)

var (
	// ErrInvalidIP is returned when an IP address cannot be parsed as a valid
	// IPv4 or IPv6 address.
	ErrInvalidIP = errors.New("invalid IP address")

	// ErrNeighborNotFound is returned when Lookup cannot find a requested
	// neighbor by its IP address.
	ErrNeighborNotFound = errors.New("neighbor not found")

	// ErrNotImplemented is returned when neighbor lookup functionality is
	// not implemented on the host operating system.
	ErrNotImplemented = errors.New("not implemented")
)

// Neighbor is a neighbor residing on the same link as the host operating
// system.
type Neighbor struct {
	// Interface specifies the host network interface which was used to
	// contact a neighbor.
	Interface *net.Interface

	// IP specifies the neighbor's IPv4 or IPv6 address.
	IP net.IP

	// HardwareAddr specifies the neighbor's MAC address.
	HardwareAddr net.HardwareAddr

	// Flags specifies any special flags indicated by the neighbor, such as
	// if the neighbor is a router or proxy.
	Flags Flags

	// State specifies the current state of the neighbor, according to the
	// host operating system, such as reachable, unreachable, etc.
	State State
}

// Show returns all available Neighbors for a given Family.
func Show(family Family) ([]*Neighbor, error) {
	return getNeighbors(family)
}

// Lookup attempts to look up a specific neighbor by its IPv4 or IPv6
// address.
//
// If a Neighbor with the matching IP address cannot be found,
// ErrNeighborNotFound is returned.
//
// If ip is an invalid IPv4 or IPv6 address, ErrInvalidIP is returned.
func Lookup(ip net.IP) (*Neighbor, error) {
	// Determine ip's family and convert to IPv4 or IPv6 form
	fIP, family, err := ipFamily(ip)
	if err != nil {
		return nil, err
	}

	// Retrieve all neighbors with the designated address family
	nn, err := Show(family)
	if err != nil {
		return nil, err
	}

	// Attempt to find a matching neighbor with the given IP address,
	// and return it if found.
	for _, n := range nn {
		if bytes.Equal(n.IP, fIP) {
			return n, nil
		}
	}

	// Neighbor could not be found
	return nil, ErrNeighborNotFound
}

// ipFamily convert's an input ip to IPv4 or IPv6 form, and returns the
// matching Family.
//
// If IP is not an IPv4 or IPv6 address, ErrInvalidIP is returned.
func ipFamily(ip net.IP) (net.IP, Family, error) {
	if ip := ip.To4(); ip != nil {
		return ip, FamilyIPv4, nil
	}
	if ip := ip.To16(); ip != nil {
		return ip, FamilyIPv6, nil
	}

	return nil, 0, ErrInvalidIP
}
