package main

import (
	"fmt"
	"github.com/jakubd/ipdetails"
)

func main() {
	fmt.Println(ipdetails.Lookup("81.2.69.142"))
}