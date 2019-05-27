package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	fmt.Println("Hello")

	err := dns.ListenAndServe("localhost", "", nil)
	if err != nil {
		fmt.Println(err)
	}
}
