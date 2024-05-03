package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i, val := range b[:n] {
		if val >= 'a' && val <= 'z' {
			b[i] = (val-'a'+13)%26 + 'a'
		} else if val >= 'A' && val <= 'Z' {
			b[i] = (val-'A'+13)%26 + 'A'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
