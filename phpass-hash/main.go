package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mhemmings/phpass"
)

var pw = ""

func main() {
	flag.StringVar(&pw, "pw", "", "Password to hash")
	flag.Parse()
	if pw == "" {
		pw = strings.Join(flag.Args(), " ")
	}
	var h = phpass.New(nil)
	hash, _ := h.Hash([]byte(pw))
	fmt.Println(string(hash))
}
