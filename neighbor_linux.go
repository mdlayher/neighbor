// +build linux

package neighbor

// Copyright (c) 2012 The Go Authors. All rights reserved.
// Source code in this file is based on src/net/interface_linux.go,
// from the Go standard library.  The Go license can be found here:
// https://golang.org/LICENSE.

import (
	"net"
	"os"
	"syscall"
	"unsafe"
)

// Constants taken from <linux/neighbour.h>, for use in parsing netlink
// route attribute types specific to neighbor discovery.
const (
	ndaDST    = 1
	ndaLLADDR = 2
)

// ndmsg is a raw ndmsg struct, used to create a Neighbor.
type ndmsg struct {
	Family  byte
	Pad     uint8
	IfIndex int32
	State   uint16
	Flags   uint16
	Type    uint16
}

// getNeighbors sends a request to netlink to retrieve all neighbors using
// the specified address family.
func getNeighbors(family Family) ([]*Neighbor, error) {
	// Request neighbors belonging to a specific family from netlink
	tab, err := syscall.NetlinkRIB(syscall.RTM_GETNEIGH, int(family))
	if err != nil {
		return nil, os.NewSyscallError("netlink rib", err)
	}

	// Parse netlink information into individual messages
	msgs, err := syscall.ParseNetlinkMessage(tab)
	if err != nil {
		return nil, os.NewSyscallError("netlink message", err)
	}

	// Check messages for information
	var nn []*Neighbor
	for _, m := range msgs {
		// Ignore any messages which don't indicate a new neighbor
		if m.Header.Type != syscall.RTM_NEWNEIGH {
			continue
		}

		// Attempt to parse an individual neighbor from a message
		n, err := newNeighbor(&m)
		if err != nil {
			return nil, err
		}

		nn = append(nn, n)
	}

	return nn, nil
}

// newNeighbor attempts to parse a netlink message into a Neighbor.
func newNeighbor(m *syscall.NetlinkMessage) (*Neighbor, error) {
	// Gather netlink route attributes from message
	attrs, err := parseNetlinkRouteAttr(m)
	if err != nil {
		return nil, err
	}

	// Iterate attributes in search of values for MAC address and IP
	// address; other values can be inferred by ndmsg struct directly
	var ip net.IP
	var mac net.HardwareAddr
	for _, a := range attrs {
		if a.Attr.Type == ndaDST {
			ip = net.IP(a.Value)
		}

		if a.Attr.Type == ndaLLADDR {
			mac = net.HardwareAddr(a.Value)
		}
	}

	// Cast raw data into a ndmsg struct to obtain more information for
	// a Neighbor
	ndm := (*ndmsg)(unsafe.Pointer(&m.Data[0]))

	// Retrieve local interface designated by index, indicating which interface
	// contacted this neighbor
	ifi, err := net.InterfaceByIndex(int(ndm.IfIndex))
	if err != nil {
		return nil, err
	}

	return &Neighbor{
		Interface:    ifi,
		IP:           ip,
		HardwareAddr: mac,
		Flags:        Flags(ndm.Flags),
		State:        State(ndm.State),
	}, nil
}

// Required for parseNetlinkRouteAttr, taken from C.
const sizeofNdmsg = 12

// This code is copied from src/syscall/netlink_linux.go in the Go standard
// library.
//
// This code is required because syscall.ParseNetlinkRouteAttr cannot handle
// syscall.RTM_NEWNEIGH messages.  An effort will be made to add the change
// upstream, but it may not be possible due to the syscall package freeze.

// parseNetlinkRouteAttr parses m's payload as an array of netlink
// route attributes and returns the slice containing the
// NetlinkRouteAttr structures.
func parseNetlinkRouteAttr(m *syscall.NetlinkMessage) ([]syscall.NetlinkRouteAttr, error) {
	var b []byte
	switch m.Header.Type {
	// Added specifically for use with this package;
	// other cases have been removed to ensure that no other message
	// types can be processed by this function
	case syscall.RTM_NEWNEIGH:
		b = m.Data[sizeofNdmsg:]
	default:
		return nil, syscall.EINVAL
	}
	var attrs []syscall.NetlinkRouteAttr
	for len(b) >= syscall.SizeofRtAttr {
		a, vbuf, alen, err := netlinkRouteAttrAndValue(b)
		if err != nil {
			return nil, err
		}
		ra := syscall.NetlinkRouteAttr{Attr: *a, Value: vbuf[:int(a.Len)-syscall.SizeofRtAttr]}
		attrs = append(attrs, ra)
		b = b[alen:]
	}
	return attrs, nil
}

func netlinkRouteAttrAndValue(b []byte) (*syscall.RtAttr, []byte, int, error) {
	a := (*syscall.RtAttr)(unsafe.Pointer(&b[0]))
	if int(a.Len) < syscall.SizeofRtAttr || int(a.Len) > len(b) {
		return nil, nil, 0, syscall.EINVAL
	}
	return a, b[syscall.SizeofRtAttr:], rtaAlignOf(int(a.Len)), nil
}

// Round the length of a netlink route attribute up to align it
// properly.
func rtaAlignOf(attrlen int) int {
	return (attrlen + syscall.RTA_ALIGNTO - 1) & ^(syscall.RTA_ALIGNTO - 1)
}
