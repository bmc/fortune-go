// fortune is a stripped-down implementation of the classic BSD Unix
// fortune command. Unlike the BSD fortune command (or my own Python version,
// at https://github.com/bmc/fortune), this version does not use an index file.
// We have loads of memory these days, and fortunes files aren't that big, so
// it's feasible to load the whole text file in memory, parse it on the fly,
// and randomly choose a resulting fortune.
//
// See the accompanying README for more information.

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

const VERSION = "2.0"

// ---------------------------------------------------------------------------
// Local error object that conforms to the Go error interface.
// ---------------------------------------------------------------------------

type Error struct {
        Message string
}

func (e *Error) Error() string {
        return e.Message
}

// ---------------------------------------------------------------------------
// Local functions
// ---------------------------------------------------------------------------

// Convenience function to print a message (printf style) to standard error
// and exit with a non-zero exit code.
func die(format string, args ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf(format, args...) + "\n")
	os.Exit(1)
}

// Given a path representing a fortune file, load the file, parse it,
// an return an array of fortune strings.
func readFortuneFile(fortuneFile string) ([]string, error) {
	content, err := ioutil.ReadFile(fortuneFile)
	var fortunes []string = nil
	if err == nil {
		fortunes = strings.Split(string(content), "\n%\n")
	}
	return fortunes, err
}

// Given a path representing a fortune file, load the file, parse it,
// choose a random fortune, and display it to standard output. Returns
// a Go error object on error or nil on success.
func findAndPrint(fortuneFile string) error {
	fortunes, err := readFortuneFile(fortuneFile)
	if err == nil {
		rand.Seed(time.Now().UTC().UnixNano())
		i := rand.Int() % len(fortunes)
		fmt.Println(fortunes[i])
	}
	return err
}

// Parse the command line arguments. For now, this is simple, because this
// program requires very few arguments. If something more complicated is
// needed, consider the Go "flag" module or github.com/docopt/docopt-go
func parseArgs() (string, error) {
        prog := path.Base(os.Args[0])
        usage := fmt.Sprintf(`%s, version %s

Usage:
  %s [/path/to/fortune/cookie/file]
  %s -h|--help

If the fortune cookie file path is omitted, the contents of environment
variable FORTUNE_FILE will be used. If neither is available, fortune will
abort.`, prog, VERSION, prog, prog)

        var fortuneFile string
        var err error

        switch len(os.Args) {
        case 1:
                fortuneFile = os.Getenv("FORTUNE_FILE")
        case 2:
                {
                        if (os.Args[1] == "-h") || (os.Args[1] == "--help") {
                                err = &Error{usage}
                        } else {
                                fortuneFile = os.Args[1]
                        }
                }
        default:
                err = &Error{usage}
        }

        if (err == nil) && (fortuneFile == "") {
                err = &Error{"No fortunes parameter and no FORTUNE_FILE " +
                                 "environment variable"}
        }

        return fortuneFile, err
}

// ---------------------------------------------------------------------------
// Main program
// ---------------------------------------------------------------------------

func main() {
        fortuneFile, err := parseArgs()
        if err != nil {
                die(err.Error())
        }

	err = findAndPrint(fortuneFile)
	if err != nil {
		die(err.Error())
	}
}
