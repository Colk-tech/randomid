package main

import (
	"github.com/Colk-tech/randomid/arguments"
	"github.com/Colk-tech/randomid/handler"
	"os"
)

var (
	opts arguments.Options
)

func main() {
	err := arguments.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	handler.Handle(&opts)
}
