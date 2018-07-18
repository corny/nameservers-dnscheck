package main

import (
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

var geoDbPath string

func location(address string) (isocode string, city string) {
	db, err := geoip2.Open(geoDbPath)
	if err != nil {
		log.Fatalf("cannot open geoip database %q: %v", geoDbPath, err)
	}
	defer db.Close()

	record, err := db.City(net.ParseIP(address))
	if err != nil {
		log.Fatalf("cannot resolve IP address: %v", err)
	}

	return record.Country.IsoCode, record.City.Names["en"]
}
