package main

import (
	"flag"

	"github.com/a8uhnf/suich/cmd"
)

func main() {
	flag.Parse()

	c := cmd.RootCmd()
	c.Execute()
}