package printer

import (
	"os"

	"github.com/mattn/go-isatty"
)

// The variables below are initialized in init().
var (
	errPrefix     string
	errSuffix     string
	h1Prefix      string
	h1Suffix      string
	h2Prefix      string
	h2Suffix      string
	successPrefix string
	successSuffix string
)

func init() {
	terminal := isTerminal()

	errPrefix = "Error: "
	if terminal {
		errPrefix = escape(escapeBold, escapeRed) + errPrefix + escape(escapeReset, escapeRed)
		errSuffix = escape(escapeReset)
	}

	h1Prefix = "====> "
	if terminal {
		h1Prefix = escape(escapeBold) + h1Prefix
		h1Suffix = escape(escapeReset)
	}

	h2Prefix = "----> "
	if terminal {
		h2Prefix = escape(escapeBold) + h2Prefix
		h2Suffix = escape(escapeReset)
	}

	if terminal {
		successPrefix = escape(escapeBold, escapeGreen)
		successSuffix = escape(escapeReset)
	}
}

func isTerminal() bool {
	// Copied and inverted from:
	// https://github.com/fatih/color/blob/daf2830f2741ebb735b21709a520c5f37d642d85/color.go#L20.
	return os.Getenv("TERM") != "dumb" && (isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()))
}
