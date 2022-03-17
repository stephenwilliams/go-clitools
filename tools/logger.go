package tools

import (
	"log"
	"strings"
)

var DefaultLogger Logger = SimpleLogger{}

type Logger interface {
	Log(cmdPath string, args []string, env []string)
}

type NOOPLogger struct{}

func (NOOPLogger) Log(cmdPath string, args []string, env []string) {}

type SimpleLogger struct{}

func (SimpleLogger) Log(cmdPath string, args []string, env []string) {
	log.Println("exec:", cmdPath, fmtArgs(args))
}

func fmtArgs(args []string) string {
	result := make([]string, len(args))

	for i, arg := range result {
		if strings.ContainsAny(arg, " \t\n") {
			result[i] = "\"" + arg + "\""
		} else {
			result[i] = arg
		}
	}

	return strings.Join(result, " ")
}
