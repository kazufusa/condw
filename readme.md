# condw

condw.CondWriter creates a writer that output to one writer of provided writers,
whose key is corresponding with the prefix of the output buffer.

## usage of selectw package

```example.go
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
```

```
$ go run example.go 2> /dev/null
STDOUT: hello stdout
$ go run example.go > /dev/null
STDERR: hello stderr
```
