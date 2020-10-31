package printer

import "context"

type Interface interface {
	Fmt

	// Err returns a printer printing to stderr.
	Err() Fmt
	// H1 returns a printer printing level 1 header messages to stdout.
	H1() Fmt
	// H2 returns a printer printing level 2 header messages to stdout.
	H2() Fmt

	// Success prints success message to stdout.
	Success()

	// WaitIndicator prints a wait indicator if TTY is set. It needs to be
	// stopped. After stopping it disappears from the screen.
	WaitIndicator(ctx context.Context) WaitIndicator
}

type Fmt interface {
	Print(a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
}

type WaitIndicator interface {
	Stop()
}
