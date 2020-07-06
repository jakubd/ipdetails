package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jakubd/ipdetails"
	"io"
	"os"
	"strings"
)

func outputLookup(givenInput string, intel bool) {
	ipinfo, err := ipdetails.Lookup(givenInput)
	status := "good_ip"
	if err != nil {
		status = "bad_ip"
	}
	var record []string
	record = []string{ipinfo.IPStr, ipinfo.CountryCode, ipinfo.ASName, ipinfo.ASNumStr, status}

	if intel {
		intelrecord := []string{
			" https://censys.io/ipv4/" + ipinfo.IPStr + " ",
			" https://www.shodan.io/host/" + ipinfo.IPStr + " ",
			" https://bgp.he.net/" + ipinfo.ASNumStr + " ",
		}
		record = append(record, intelrecord...)
	}
	fmt.Println(strings.Join(record, ","))
}

func main() {
	_, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	var runeInput []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		runeInput = append(runeInput, input)
	}

	var ipintel bool
	flag.BoolVar(&ipintel, "ipintel", false, "add urls for ipintel")
	flag.Parse()

	if ipintel {
		fmt.Println("yes to intel!")
	}

	var thisLine string
	for j := 0; j < len(runeInput); j++ {
		thisLine = thisLine + string(runeInput[j])
		if runeInput[j] == '\n' {
			if len(thisLine) > 1 {
				thisLine = strings.TrimSuffix(thisLine, "\n")
				outputLookup(thisLine, ipintel)
				thisLine = ""
			}
		}
	}
}
