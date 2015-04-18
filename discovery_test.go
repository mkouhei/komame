package main

import (
	"fmt"
	"testing"
)

func TestLocalNetworks(t *testing.T) {
	nws, err := localNetworks()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(nws)
}
