package pringo

// Directly reads from the command line.
func (p *Printer) Read() string {
	text, _ := p.reader.ReadString('\n')
	return text
}

// Promts the user in the same line and returns the input.
func (p *Printer) Promt(s string, c Color) string {
	p.Print(s, c)
	return p.Read()
}

// The same as Promt() but with newline.
func (p *Printer) Promtln(s string, c Color) string {
	p.Println(s, c)
	return p.Read()
}

// The same as Promt() but as a template string.
func (p *Printer) Promtf(s string, c Color, args ...interface{}) string {
	p.Printf(s, c, args...)
	return p.Read()
}
