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

type Error struct {
        Message string
}

func (e *Error) Error() string {
        return e.Message
}

func die(format string, args ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf(format, args...) + "\n")
	os.Exit(1)
}

func readFortuneFile(fortuneFile string) ([]string, error) {
	content, err := ioutil.ReadFile(fortuneFile)
	var fortunes []string = nil
	if err == nil {
		fortunes = strings.Split(string(content), "\n%\n")
	}
	return fortunes, err
}

func findAndPrint(fortuneFile string) error {
	fortunes, err := readFortuneFile(fortuneFile)
	if err == nil {
		rand.Seed(time.Now().UTC().UnixNano())
		i := rand.Int() % len(fortunes)
		fmt.Println(fortunes[i])
	}
	return err
}

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
