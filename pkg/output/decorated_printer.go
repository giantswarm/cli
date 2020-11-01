package output

import (
	"fmt"
	"io"
)

type decoratedPrinter struct {
	prefix string
	suffix string
	writer io.Writer
}

func (p *decoratedPrinter) Print(a ...interface{}) (n int, err error) {
	s := fmt.Sprint(a...)
	return fmt.Fprint(p.writer, p.decorated(s))
}

func (p *decoratedPrinter) Printf(format string, a ...interface{}) (n int, err error) {
	s := fmt.Sprintf(format, a...)
	return fmt.Fprint(p.writer, p.decorated(s))
}

func (p *decoratedPrinter) Println(a ...interface{}) (n int, err error) {
	s := fmt.Sprint(a...)
	return fmt.Fprintln(p.writer, p.decorated(s))
}

func (p *decoratedPrinter) decorated(s string) string {
	return p.prefix + s + p.suffix
}
