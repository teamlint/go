package templatex

import (
	"regexp"
	"strings"

	"github.com/teamlint/gox/stringx"
)

var striptagsRegexp = regexp.MustCompile("<[^>]*?>")

func striptags(s string) string {
	return strings.TrimSpace(striptagsRegexp.ReplaceAllString(s, ""))
}

func truncate(n int, value string) string {
	return stringx.Left(value, n, "")
}
func join(arg string, value []string) string {
	defer recovery()

	return strings.Join(value, arg)
}
