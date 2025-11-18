package clicmds

import (
	"flag"
	"fmt"
)

func Add(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	var x, y int
	fs.IntVar(&x, "val1", 0, "val 1")
	fs.IntVar(&y, "val2", 0, "val 1")
	fs.Parse(args)

	fmt.Println(x + y)
}
