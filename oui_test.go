package main

import (
	"testing"
)

var tdataPath = "testdata/oui.txt"

func TestReadOui(t *testing.T) {
	c := &ouiData{}
	c.readOui(tdataPath)
	if len(c.ouis) != 20507 {
		t.Fatal("Fail reading oui.txt")
	}
}

func TestGetOuiData(t *testing.T) {
	if err := getOuiData(tdataPath); err != nil {
		t.Fatal(err)
	}
}
