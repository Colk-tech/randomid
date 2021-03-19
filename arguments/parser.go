package arguments

import (
	"github.com/Colk-tech/randomid/config"
	flags "github.com/jessevdk/go-flags"
)

func MakeParser(opts *Options) *flags.Parser {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = config.MyName
	parser.Usage = "[OPTIONS]"

	return parser
}
