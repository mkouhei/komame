package main

import "net"

type nic struct {
	name     net.Interface
	macAddr  net.HardwareAddr
	ipv4Addr net.IPAddr
	v4Mask   net.IPMask
	ipv6Addr net.IPAddr
	v6Mask   net.IPMask
	status   bool // true: up; false: down
}

type node struct {
	name string
	nics []nic
}
