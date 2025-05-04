package helpers

import (
	"log"
	"os"
)

var DebugLogger, VerboseLogger *log.Logger

// Note this will be usd for prompts as well
var StdErrLogger = log.New(os.Stderr, "error: ", 0)

func InitVerboseAndDebugLogger(verbose, debug bool) {
	if verbose {
		VerboseLogger = log.New(os.Stderr, "verbose: ", 0)
	}

	if debug {
		DebugLogger = log.New(os.Stderr, "debug: ", 0)
	}
}
