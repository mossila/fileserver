package main

import (
	"log"
	"net"

	"github.com/mossila/fileserver"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				log.Printf("IPNet %s %s \n", i.Name, ip)
			case *net.IPAddr:
				ip = v.IP
				// log.Printf("IPAddr %s\n", ip)
			}
			// log.Println(ip)
		}
	}
	fileserver.Serve(".")
}
