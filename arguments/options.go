package arguments

type Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Length uint16 `short:"l" long:"length" description:"Determine number of digits of ID" default:"4"`
}
