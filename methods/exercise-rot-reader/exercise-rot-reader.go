package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	const ALPHABET_COUNT = 26
	const OFFSET = 13

	n, e := r.r.Read(b)
	if e == nil {
		for i, v := range b {
			switch {
			case v >= 'A' && v <= 'Z':
				b[i] = (v-'A'+OFFSET)%ALPHABET_COUNT + 'A'
			case v >= 'a' && v <= 'z':
				b[i] = (v-'a'+OFFSET)%ALPHABET_COUNT + 'a'
			}
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
