package pringo

import (
	"bufio"
	"os"
)

type Printer struct {
	stdin  *os.File
	stdout *os.File
	stderr *os.File
	reader *bufio.Reader
}

// Creates a new Printer instance with set stdin, stdout and stderr.
func New(stdout *os.File, stderr *os.File, stdin *os.File) *Printer {
	if stdout == nil {
		stdout = os.Stdout
	}
	if stderr == nil {
		stderr = os.Stderr
	}
	if stdin == nil {
		stdin = os.Stdin
	}
	return &Printer{
		stdin:  stdin,
		stdout: stdout,
		stderr: stderr,
		reader: bufio.NewReader(stdin),
	}
}

// The same as New() but with os.Stdout, os.Stderr and os.Stdin as default.
func NewDefault() *Printer {
	return New(os.Stdout, os.Stderr, os.Stdin)
}
