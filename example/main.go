package main

import (
	"bufio"
	"fmt"
	"github.com/jakubd/ipdetails"
	"io"
	"os"
	"strings"
)

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

	var thisLine string
	for j := 0; j < len(runeInput); j++ {
		thisLine = thisLine + string(runeInput[j])
		if runeInput[j] == '\n' {
			if len(thisLine) > 1 {
				thisLine = strings.TrimSuffix(thisLine, "\n")
				ipinfo, _ := ipdetails.Lookup(thisLine)
				record := []string{
					ipinfo.IPStr,
					ipinfo.CountryCode,
					ipinfo.ASName,
					ipinfo.ASNumStr,
				}
				fmt.Println(strings.Join(record, ","))
				thisLine = ""
			}
		}
	}
}
