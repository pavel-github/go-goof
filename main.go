package main

import (
	"fmt"

	"github.com/mholt/archiver"
)

func main() {
	fmt.Println("Hello")

	err := archiver.Archive([]string{"Gopkg.lock", "Gopkg.toml"}, ".Gopkg.tar.gz")
	if err != nil {
		fmt.Println(err)
	}
}
