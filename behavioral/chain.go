package behavioral

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (l *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)
	if l.NextChain != nil {
		l.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (l *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if l.NextChain != nil {
			l.NextChain.Next(s)
		}

		return
	}
	fmt.Printf("Finishing in second logging\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("WriterLogger: " + s))
	}
	if w.NextChain != nil {
		w.NextChain.Next(s)
	}
}
