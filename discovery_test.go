package main

import (
	"fmt"
	"testing"
)

func TestlocalNetworks(t *testing.T) {
	nws, err := localNetworks()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(nws)
}
