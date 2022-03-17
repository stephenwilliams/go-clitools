package versionproviders

import (
	"fmt"
	"strings"
)

func SplitLine(s string, i int) string {
	lines := strings.Split(s, "\n")

	if len(lines) <= i {
		panic(fmt.Errorf("version provider splitLines does not have lines past index '%d'", i))
	}

	return lines[i]
}

func SplitString(s, sep string, i int) string {
	items := strings.Split(s, sep)

	if len(items) <= i {
		panic(fmt.Errorf("version provider Split does not have items past index '%d' with sep '%s', only len '%d', string '%s'", i, sep, len(items), s))
	}

	return items[i]
}

func EqualsReplace(s, old, new string) string {
	if strings.TrimSpace(s) == old {
		return new
	}

	return s
}
