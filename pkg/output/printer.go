package output

import (
	"context"
	"fmt"
	"io"
	"os"
)

type PrinterConfig struct {
	Err io.Writer
	Out io.Writer
}

type Printer struct {
	*decoratedPrinter

	err io.Writer
	out io.Writer
}

func NewPrinter(config PrinterConfig) (*Printer, error) {
	if config.Err == nil {
		config.Err = os.Stderr
	}
	if config.Out == nil {
		config.Out = os.Stdout
	}

	p := &Printer{
		decoratedPrinter: &decoratedPrinter{
			writer: config.Out,
		},

		err: config.Err,
		out: config.Out,
	}

	return p, nil
}

func (p *Printer) Err() Fmt {
	return &decoratedPrinter{
		prefix: errPrefix,
		suffix: errSuffix,
		writer: p.err,
	}
}

func (p *Printer) H1() Fmt {
	return &decoratedPrinter{
		prefix: h1Prefix,
		suffix: h1Suffix,
		writer: p.out,
	}
}

func (p *Printer) H2() Fmt {
	return &decoratedPrinter{
		prefix: h2Prefix,
		suffix: h2Suffix,
		writer: p.out,
	}
}

func (p *Printer) Success() {
	fmt.Fprintln(p.out, successPrefix+"Success!"+successSuffix)
}

func (p *Printer) WaitIndicator(ctx context.Context) WaitIndicator {
	w := newWaitIndicator(p.out)
	go w.Start(ctx)
	return w
}
