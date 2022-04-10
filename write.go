package pringo

// Prints to stdout but without newline and color.
func (p *Printer) Write(s string) {
	p.Print(s, None)
}

// Prints to stdout as a format string but without color.
func (p *Printer) Writef(format string, args ...interface{}) {
	p.Printf(format, None, args...)
}

// Prints to stdout with newline but without color.
func (p *Printer) Writeln(s string) {
	p.Println(s, None)
}
