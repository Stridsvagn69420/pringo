package pringo

import (
	"fmt"
)

func (p *Printer) Write(s string) {
	p.stdout.Write([]byte(s))
}

func (p *Printer) Writef(format string, args ...interface{}) {
	p.stdout.Write([]byte(fmt.Sprintf(format, args...)))
}

func (p *Printer) Writeln(s string) {
	p.stdout.Write([]byte(s + "\n"))
}
