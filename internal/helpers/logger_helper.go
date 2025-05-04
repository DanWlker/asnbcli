package helpers

import (
	"io"
	"log"
	"os"
)

var DebugLogger, VerboseLogger *log.Logger

// Note this will be usd for prompts as well
var (
	StdErrLogger = log.New(os.Stderr, "error: ", 0)
)

func InitVerboseAndDebugLogger(verbose, debug bool) {
	verboseWriter := io.Discard
	if verbose {
		verboseWriter = os.Stderr
	}
	VerboseLogger = log.New(verboseWriter, "verbose: ", 0)

	debugWriter := io.Discard
	if debug {
		debugWriter = os.Stderr
	}
	DebugLogger = log.New(debugWriter, "debug: ", 0)
}
