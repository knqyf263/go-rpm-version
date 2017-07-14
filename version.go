package version

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	// alphanumPattern is a regular expression to match all sequences of numeric
	// characters or alphanumeric characters.
	alphanumPattern = regexp.MustCompile("([a-zA-Z]+)|([0-9]+)|(~)")
	allowedSymbols  = []rune{'.', '-', '+', '~', ':', '_'}
)

// Version represents a package version.
type Version struct {
	epoch   int
	version string
	release string
}

// NewVersion returns a parsed version
func NewVersion(ver string) (version Version) {
	var err error

	// Parse epoch
	splitted := strings.SplitN(ver, ":", 2)
	if len(splitted) == 1 {
		version.epoch = 0
		ver = splitted[0]
	} else {
		// Trim left space
		epoch := strings.TrimLeftFunc(splitted[0], unicode.IsSpace)

		version.epoch, err = strconv.Atoi(epoch)
		if err != nil {
			version.epoch = 0
		}

		ver = splitted[1]
	}

	// Parse version and release
	index := strings.Index(ver, "-")
	if index >= 0 {
		version.version = ver[:index]
		version.release = ver[index+1:]

	} else {
		version.version = ver
	}

	return version
}
