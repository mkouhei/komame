package main

import (
	"fmt"
	"net"
	"strings"
)

func localNetworks() ([]net.IPNet, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var nws []net.IPNet
	for _, addr := range addrs {
		if !strings.HasPrefix(addr.String(), "127.0.0.") && addr.String() != "::1/128" {
			_, nw, err := net.ParseCIDR(addr.String())
			if err != nil {
				fmt.Println(err)
			}
			nws = append(nws, *nw)
		}
	}
	return nws, nil
}
