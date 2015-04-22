package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

const (
	prefix = "oui_test"
)

var (
	tdataPath  = "testdata/oui.txt"
	dummyURL   = "http://localhost:8080"
	macAddr    = "08:00:27:9d:7c:80"
	otherMac   = "08-00-27-9d-7c-80"
	invalidMac = "gg:00:00:00:00:00"
	noConvMac  = "0800279d7c85"
)

func TestReadOui(t *testing.T) {
	c := &ouiData{}
	c.readOui(tdataPath)
	if len(c.ouis) != 20507 {
		t.Fatal("Fail reading oui.txt")
	}

	if c.search(macAddr).hex != "08-00-27" {
		t.Fatal("Failed to search OUI")
	}
}

func TestGetOuiData(t *testing.T) {
	body, err := ioutil.ReadFile(tdataPath)
	if err != nil {
		t.Fatal(err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request) {
		fmt.Fprint(w, string(body))
	}))
	defer ts.Close()

	f, err := ioutil.TempDir("", prefix)
	if err != nil {
		t.Fatal(err)
	}

	if err := getOuiData(filepath.Join(f, "oui.txt"), dummyURL); err == nil {
		t.Fatal(err)
	}

	if err := getOuiData(filepath.Join(f, "oui.txt"), ts.URL); err != nil {
		t.Fatal(err)
	}
	os.RemoveAll(f)
}

func TestConvOUI(t *testing.T) {
	oui, rc := convOUI(macAddr)
	if !rc || oui != "080027" {
		t.Fatal("Failed to convert MAC address to OUI.")
	}

	if _, rc := convOUI(invalidMac); rc {
		t.Fatal("Failed to detect MAC address")
	}

	oui, rc = convOUI(otherMac)
	if !rc || oui != "080027" {
		t.Fatal("Failed to convert MAC address to OUI.")
	}

	oui, rc = convOUI(noConvMac)
	if !rc || oui != "080027" {
		t.Fatal("Failed to extract OUI.")
	}
}
