package main

import (
	"flag"
	"strings"

	"github.com/brianfoshee/itcreporter"
)

func main() {
	r := itcreporter.New()

	flag.Parse()

	// TODO handle flags overriding properties file
	r.Command(strings.Join(flag.Args(), " "))
}
