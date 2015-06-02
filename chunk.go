package chunk

// Splitter returns a sequence of data chunks from a source.
// Construct a Splitter with the New function.
// (I can't decide! Chunker or Splitter?!)
type Splitter interface {
	// Next returns the next split chunk. Next returns io.EOF after all
	// chunks have been split.
	Next() ([]byte, error)
}

// FullSplit splits an io.Reader in full. It will consume all of the
// reader. The error will be nil if the reader finishes (io.EOF),
// but all other errors are forwarded.
func FullSplit(s Splitter) ([][]byte, error) {
	all := make([][]byte)
	for {
		next, err := s.Next()
		if err == io.EOF {
			return all, nil // normal exit condition
		}
		if err != nil {
			return all, err
		}
	}
}

// Chan takes a Splitter, returns an unbuffered channel. A goroutine
// will continue splitting until EOF. It returns an error channel
// and will send an error through it if chunking fails.
func Chan(s Splitter) (<-chan []byte, <-chan error) {
	outCh := make(chan []byte)
	errCh := make(chan error, 1) // buffered so error can be safely ignored
	go func() {
		defer close(outCh)
		defer close(errCh)

		for {
			next, err := s.Split(r)
			if err != nil {
				err <- err
				return
			}

			outCh <- next
		}
	}()
	return outCh
}
