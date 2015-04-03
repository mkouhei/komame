package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type ouiData struct {
	utime time.Time
	ouis  []oui
}

type oui struct {
	name    string
	hex     string
	base16  string
	address string
}

func getOuiData(p, url string) error {
	_, err := os.Stat(p)
	if err != nil {
		if url == "" {
			url = ouiURL
		}
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		f.Write(body)
	}
	return nil
}

func (c *ouiData) readOui(p string) error {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}
	payload := strings.Split(string(b), "\n\n")

	// generated time of oui.txt
	u, err := time.Parse(time.RFC1123Z, strings.Split(payload[0], "Generated: ")[1])
	if err != nil {
		return err
	}
	c.utime = u

	// oui data
	// The First record is description
	// The last record is IEEE registration authority
	data := payload[2 : len(payload)-1]
	o := &oui{}
	for _, v := range data {
		o.parseOui(v)
		c.ouis = append(c.ouis, *o)
	}
	return nil
}

func (o *oui) parseOui(oui string) {
	o.address = ""
	for i, s := range strings.Split(oui, "\n") {
		switch {
		case i == 0:
			o.hex = strings.Fields(s)[0]
			o.name = strings.TrimSpace(strings.Split(s, "(hex)")[1])
		case i == 1:
			o.base16 = strings.Fields(s)[0]
		default:
			if strings.TrimSpace(s) != "" {
				o.address += strings.TrimSpace(s) + " "
			}
		}
	}
	o.address = strings.TrimSpace(o.address)
}

func convOUI(ma string) (string, bool) {
	// 08:00:27:9d:7c:85
	pat0 := regexp.MustCompile(`\A(([0-9a-fA-F]{2}\:){5})([0-9a-fA-F]{2})\z`)
	// 08-00-27-9d-7c-85
	pat1 := regexp.MustCompile(`\A(([0-9a-fA-F]{2}\-){5})([0-9a-fA-F]{2})\z`)
	// 0800279d7c85
	pat2 := regexp.MustCompile(`\A[0-9a-fA-F]{12}\z`)
	if pat0.MatchString(ma) {
		return strings.Replace(ma, ":", "", -1)[:6], true
	}
	if pat1.MatchString(ma) {
		return strings.Replace(ma, "-", "", -1)[:6], true
	}
	if pat2.MatchString(ma) {
		return ma[:6], true
	}
	return "", false
}
