// Command neighborhttp provides a simple HTTP server which retrieves neighbor
// information for any device which makes a request to it.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/mdlayher/neighbor"
)

var (
	hostFlag = flag.String("host", ":8080", "host:port pair to listen on")
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve host only for neighbor lookup
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err)
			return
		}

		// Attempt to find neighbor
		n, err := neighbor.Lookup(net.ParseIP(host))
		if err != nil {
			log.Println(err)
			return
		}

		// Log neighbor and return information to client
		msg := fmt.Sprintf("%s -> %s [%s] %s",
			n.IP,
			n.HardwareAddr,
			n.State,
			n.Flags,
		)
		log.Println(msg)
		w.Write([]byte(msg))
	})

	// Serve HTTP on specified host
	if err := http.ListenAndServe(*hostFlag, nil); err != nil {
		log.Fatal(err)
	}
}
