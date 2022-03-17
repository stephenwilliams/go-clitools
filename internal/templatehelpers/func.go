package templatehelpers

import "strings"

var Funcs = map[string]interface{}{
	"trimPrefix": func(a, b string) string { return strings.TrimPrefix(b, a) },
}
