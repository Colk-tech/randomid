package arguments

import (
	"github.com/Colk-tech/randomid/config"
	flags "github.com/jessevdk/go-flags"
)

func Parse(opts *Options) (err error) {
	parser := getParser(opts)
	_, err = parser.Parse()

	return err
}

func getParser(opts *Options) *flags.Parser {
	parser := flags.NewParser(opts, flags.Default)
	parser.Name = config.ProgramName
	parser.Usage = config.UsageDescription

	return parser
}
