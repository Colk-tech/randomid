package handler

import (
	"fmt"
	"github.com/Colk-tech/randomid/arguments"
	"github.com/Colk-tech/randomid/logic"
	"github.com/Colk-tech/randomid/util"
	"time"
)

func Handle(opts *arguments.Options) {
	seed := time.Now().UnixNano()

	if opts.Verbose {
		fmt.Printf("INFO: Creating id with length %d from the seed %d...\n", opts.Length, seed)
	}

	util.SetRandSeed(seed)

	generated_id := logic.GenerateId(opts.Length)
	fmt.Println(generated_id)
}
