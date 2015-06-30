// Command neighbor provides neighbor detection akin to 'ip neighbor show'.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mdlayher/neighbor"
)

var (
	ipFlag = flag.String("ip", "", "IP address to lookup specific neighbor")

	ip4Flag = flag.Bool("4", false, "IPv4 only?")
	ip6Flag = flag.Bool("6", false, "IPv6 only?")
)

func main() {
	flag.Parse()

	// Attempt to discover a specific neighbor by IP address
	if *ipFlag != "" {
		n, err := neighbor.Lookup(net.ParseIP(*ipFlag))
		if err != nil {
			log.Fatal(err)
		}

		printNeighbor(n)
		return
	}

	// Determine network address family by flags
	var family neighbor.Family
	switch {
	case *ip4Flag:
		family = neighbor.FamilyIPv4
	case *ip6Flag:
		family = neighbor.FamilyIPv6
	default:
		family = neighbor.FamilyUnspec
	}

	// Attempt to show neighbors with matching family
	neighbors, err := neighbor.Show(family)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range neighbors {
		printNeighbor(n)
	}
}

// printNeighbor prints a neighbor akin to 'ip neighbor show'.
func printNeighbor(n *neighbor.Neighbor) {
	// Ignore neighbors with StateNoARP, such as loopback interface,
	// tunnels, etc.
	if n.State&neighbor.StateNoARP != 0 {
		return
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s dev %s lladdr %s",
		n.IP,
		n.Interface.Name,
		n.HardwareAddr,
	))

	buf.WriteString(" " + n.State.String())

	if n.Flags != 0 {
		buf.WriteString(" " + n.Flags.String())
	}

	fmt.Println(buf.String())
}
