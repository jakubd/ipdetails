package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
	"path"
	"strconv"
)

type IPInfo struct {
	IPStr string
	IP net.IP
	ASNum int
	ASNumStr string
	ASName string
	CountryCode string
	CountryName string
}

func openMaxmindDb(givenDbName string, givenDirectory ...string) (*geoip2.Reader, error){
	var maxmindDirectory string
	if len(givenDirectory) == 0 {
		maxmindDirectory = "/var/lib/GeoIP/"
	} else {
		maxmindDirectory = givenDirectory[0]
	}

	maxmindDb, err := geoip2.Open(path.Join(maxmindDirectory, givenDbName))
	if err != nil {
		return nil, err
	}

	return maxmindDb, nil
}

func main() {
	// If you are using strings that may be invalid, check that ip is not nil
	asnDb, _ := openMaxmindDb("GeoLite2-ASN.mmdb")
	countryDb, _ := openMaxmindDb("GeoLite2-Country.mmdb")
	defer asnDb.Close()
	defer countryDb.Close()

	givenIP := "81.2.69.142"
	ip := net.ParseIP(givenIP)

	asnRecord, err := asnDb.ASN(ip)
	if err != nil {
		panic(err)
	}

	countryRecord, err := countryDb.Country(ip)
	if err != nil {
		panic(err)
	}

	var ipinfo =  IPInfo{
		givenIP,
		ip,
		int(asnRecord.AutonomousSystemNumber),
		"AS" + strconv.Itoa(int(asnRecord.AutonomousSystemNumber)),
		asnRecord.AutonomousSystemOrganization,
		countryRecord.Country.IsoCode,
		countryRecord.Country.Names["en"],
	}

	fmt.Println(ipinfo)
}