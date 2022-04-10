package pringo

import (
	"fmt"
)

// Prints to stderr with color but without a newline.
func (p *Printer) Error(s string, c Color) {
	p.Stderr.Write([]byte(string(c) + s + string(reset)))
}

// Errorf is the same as Error() but with a format string.
func (p *Printer) Errorf(format string, c Color, args ...interface{}) {
	p.Stderr.Write([]byte(string(c) + fmt.Sprintf(format, args...) + string(reset)))
}

// Errorln is the same as Error() but with a newline at the end.
func (p *Printer) Errorln(s string, c Color) {
	p.Stderr.Write([]byte(string(c) + s + string(reset) + "\n"))
}
