package sizechunk

import (
	"io"

	chunk "github.com/jbenet/go-chunk"
)

type Splitter struct {
	R    io.Reader
	Size int
}

func (s *Splitter) Next() ([]byte, error) {
	buf := make([]byte, s.Size)
	n, err := io.ReadFull(s.R, buf)
	buf = buf[:n]
	return buf, err
}
