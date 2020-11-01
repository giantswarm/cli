package output

import (
	"context"
	"fmt"
	"io"
	"time"
)

type waitIndicator struct {
	stopCh chan struct{}
	writer io.Writer
}

func newWaitIndicator(w io.Writer) *waitIndicator {
	return &waitIndicator{
		stopCh: make(chan struct{}),
		writer: w,
	}
}

func (w *waitIndicator) Start(ctx context.Context) {
	if !isTerminal() {
		return
	}

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	transitions := []string{"|", "/", "-", "\\"}

	// Print leading space for first backspace.
	fmt.Fprint(w.writer, " ")

	i := 0
	for {
		i++
		i = i % len(transitions)

		fmt.Fprint(w.writer, "\b")
		select {
		case <-ctx.Done():
			return
		case <-w.stopCh:
			return
		case <-ticker.C:
			fmt.Fprint(w.writer, transitions[i])
		}
	}

}

func (w *waitIndicator) Stop() {
	close(w.stopCh)
}
