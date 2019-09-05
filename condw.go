package condw

import (
	"io"
	"strings"
)

var _ io.Writer = (*condWriter)(nil)

// condw.CondWriter creates a writer that output to one writer of provided writers,
// whose key is corresponding with the prefix of the output buffer.
func CondWriter(writers map[string]io.Writer) io.Writer {
	return &condWriter{writers: writers}
}

type condWriter struct {
	writers map[string]io.Writer
}

func (t *condWriter) Write(p []byte) (n int, err error) {
	s := string(p)
	for k, w := range t.writers {
		if strings.HasPrefix(s, k) {
			n, err = w.Write(p)
			if err != nil {
				return n, err
			}
			if n != len(p) {
				return n, io.ErrShortWrite
			}
		}
	}
	return len(p), nil
}
