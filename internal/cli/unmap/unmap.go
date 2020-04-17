package unmap

import (
	"fmt"
	"log"
	"os"
	"testing/iotest"

	"github.com/bensallen/rbdmap/krbd"
	flag "github.com/spf13/pflag"
)

const usageHeader = `unmap - Unmap RBD Image

Usage:
  unmap

Flags:
`

var (
	flags = flag.NewFlagSet("unmap", flag.ContinueOnError)
	devid = flags.IntP("devid", "d", 0, "RBD Device Index (default 0)")
	force = flags.BoolP("force", "f", false, "Optional force argument will wait for running requests and then unmap the image")
)

// Usage of the unmap subcommand
func Usage() {
	fmt.Fprintf(os.Stderr, usageHeader)
	fmt.Fprintf(os.Stderr, flags.FlagUsagesWrapped(0)+"\n")
}

// Run the unmap subcommand
func Run(args []string, verbose bool, noop bool) error {
	flags.ParseErrorsWhitelist.UnknownFlags = true
	if err := flags.Parse(args); err != nil {
		Usage()
		fmt.Printf("Error: %v\n\n", err)
		os.Exit(2)
	}

	w, err := krbd.RBDBusRemoveWriter()
	if err != nil {
		return err
	}

	if verbose {
		w = iotest.NewWriteLogger("unmap", w)
	}

	i := krbd.Image{
		DevID: *devid,
		Options: &krbd.Options{
			Force: *force,
		},
	}

	if noop {
		log.Printf("%s", i)
	} else {
		return i.Unmap(w)
	}
	return nil
}