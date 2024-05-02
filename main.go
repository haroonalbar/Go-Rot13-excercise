package main

import (
	"fmt"
	"io"
	// "os"
	"strings"
)

type rot13Reader struct {
	io.Reader
}

// first two case to check if it's form a,A to z,Z -13
// and the second two case is to check if z,Z -13 to z,Z
func rot13(x byte) byte {
	switch {
	case x >= 65 && x <= 77:
		fallthrough
	case x >= 97 && x <= 109:
		return x + 13
	case x >= 78 && x <= 90:
		fallthrough
	case x >= 110 && x <= 122:
		return x - 13
	default:
		return x
	}
}

func rot13convert(b []byte) []byte {
	for i, val := range b {
		x := rot13(val)
		b[i] = x
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	b, _ := io.ReadAll(r.Reader)
	x := rot13convert(b)
	// io.Copy(os.Stdout, &r)
	fmt.Println(string(x))
}
