package structural

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(p.Writer, "%s", msg)
	return nil
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *NormalPrinter) Print() error {
	p.Printer.PrintMessage(p.Msg)
	return nil
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *PacktPrinter) Print() error {
	p.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", p.Msg))
	return nil
}
