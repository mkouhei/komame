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
	tdataPath = "testdata/oui.txt"
	dummyURL  = "http://localhost"
)

func TestReadOui(t *testing.T) {
	c := &ouiData{}
	c.readOui(tdataPath)
	if len(c.ouis) != 20507 {
		t.Fatal("Fail reading oui.txt")
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
