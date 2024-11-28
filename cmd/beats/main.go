package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/peterhellberg/beats"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		if arg == "-h" || arg == "--help" || arg == "help" {
			help(filepath.Base(os.Args[0]))
			os.Exit(0)
		}
	}

	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Error: too many arguments")
		os.Exit(2)
	}

	if len(args) == 0 {
		fmt.Println(beats.NowString())
		os.Exit(0)
	}

	// We have one argument. Parse it as a Unix timestamp.
	timestamp, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: can't parse Unix timestamp")
		os.Exit(2)
	}

	t := time.Unix(timestamp, 0)
	beatTime := beats.FromTime(t)
	fmt.Println(beats.String(beatTime))
}

func help(me string) {
	fmt.Printf(
		`Usage: %[1]s [<timestamp>]

Arguments:
  [<timestamp>]
          Unix timestamp to convert to .beat time

Examples:

$ %[1]s
@456

$ %[1]s 1732500000
@125
`,
		me,
	)
}
