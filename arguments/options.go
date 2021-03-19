package arguments

type Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Ununique bool `short:"n" long:"ununique" description:"Generate ID which may not unique on this machine"`
	Length int `short:"l" long:"length" description:"Determine number of digits of ID" default:"4"`
	Seed string `short:"s" long:"seed" description:"Choose seed used while generating ID."`
	Clear bool `long:"clear" description:"Clear the history of generated IDs"`
}
