package main

import (
	"github.com/Colk-tech/randomid/arguments"
)

var (
	opts arguments.Options
)

func main() {
	parser := arguments.MakeParser(&opts)
	_, err := parser.Parse()

	if err != nil {
	}
}
