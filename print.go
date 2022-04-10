package pringo

import (
	"fmt"
)

// Prints to stdout with color but without a newline.
func (p *Printer) Print(s string, c Color) {
	p.stdout.Write([]byte(string(c) + s + string(reset)))
}

// Printf is the same as Print() but with a format string.
func (p *Printer) Printf(format string, c Color, args ...interface{}) {
	p.stdout.Write([]byte(string(c) + fmt.Sprintf(format, args...) + string(reset)))
}

// Println is the same as Print() but with a newline at the end.
func (p *Printer) Println(s string, c Color) {
	p.stdout.Write([]byte(string(c) + s + string(reset) + "\n"))
}
