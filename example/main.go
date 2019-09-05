package main

import (
	"io"
	"os"

	"github.com/kazufusa/condw"
)

func main() {
	w := condw.CondWriter(map[string]io.Writer{
		"STDOUT": os.Stdout,
		"STDERR": os.Stderr,
	})
	w.Write([]byte("STDOUT: hello stdout\n"))
	w.Write([]byte("STDERR: hello stderr\n"))
}
