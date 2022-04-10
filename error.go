package pringo

// Prints to stderr but without newline and color.
func (p *Printer) Err(s string) {
	p.Error(s, None)
}

// Prints to stderr as a format string but without color.
func (p *Printer) Errf(format string, args ...interface{}) {
	p.Errorf(format, None, args...)
}

// Prints to stderr with newline but without color.
func (p *Printer) Errln(s string) {
	p.Errorln(s, None)
}
