package printer

import "strings"

// https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	escapeReset = "0"
	escapeBold  = "1"

	escapeGreen = "32"
	escapeRed   = "31"
)

func escape(codes ...string) string {
	return "\033[" + strings.Join(codes, ";") + "m"
}
