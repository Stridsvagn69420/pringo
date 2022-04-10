package pringo

import (
	"fmt"
)

func (p *Printer) Print(s string, c Color) {
	p.stdout.Write([]byte(string(c) + s + string(reset)))
}

func (p *Printer) Printf(format string, c Color, args ...interface{}) {
	p.stdout.Write([]byte(string(c) + fmt.Sprintf(format, args...) + string(reset)))
}

func (p *Printer) Println(s string, c Color) {
	p.stdout.Write([]byte(string(c) + s + string(reset) + "\n"))
}
