package helpers

import (
	"fmt"
	"io"
	"log"
	"os"
)

var DebugLogger, VerboseLogger *log.Logger

// Note this will be usd for prompts as well
var StdErrPrintf = func(format string, a ...any) (int, error) {
	return fmt.Fprintf(os.Stderr, format, a)
}

var StdErrPrintln = func(a ...any) (int, error) {
	return fmt.Fprintln(os.Stderr, a)
}

var StdErrPrint = func(a ...any) (int, error) {
	return fmt.Fprint(os.Stderr, a)
}

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
