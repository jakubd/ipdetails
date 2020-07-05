package ipdetails

import (
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

func OpenMaxmindDb(givenDbName string, givenDirectory ...string) (*geoip2.Reader, error){
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

func Lookup(givenIpStr string) (IPInfo, error) {

	parseFailed := IPInfo{
		IPStr:       givenIpStr,
		IP:          nil,
		ASNum:       -1,
		ASNumStr:    "AS0",
		ASName:      "",
		CountryCode: "",
		CountryName: "",
	}

	asnDb, err := OpenMaxmindDb("GeoLite2-ASN.mmdb")
	if err != nil {
		return parseFailed, err
	}

	countryDb, err := OpenMaxmindDb("GeoLite2-Country.mmdb")
	if err != nil {
		return parseFailed, err
	}

	defer asnDb.Close()
	defer countryDb.Close()

	ip := net.ParseIP(givenIpStr)

	asnRecord, err := asnDb.ASN(ip)
	if err != nil {
		return parseFailed, err
	}

	countryRecord, err := countryDb.Country(ip)
	if err != nil {
		return parseFailed, err
	}

	return IPInfo{
		IPStr:       givenIpStr,
		IP:          ip,
		ASNum:       int(asnRecord.AutonomousSystemNumber),
		ASNumStr:    "AS" + strconv.Itoa(int(asnRecord.AutonomousSystemNumber)),
		ASName:      asnRecord.AutonomousSystemOrganization,
		CountryCode: countryRecord.Country.IsoCode,
		CountryName: countryRecord.Country.Names["en"],
	}, nil

}