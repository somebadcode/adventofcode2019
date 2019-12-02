package badreader

import "io"

// BadStringReader implements io.Reader and is meant to return an error when the
// internal buffer has been copied or the receiving buffer is full.
type BadStringReader struct {
	io.Reader
	buffer []byte
	error  error
}

// NewBadReader returns a reader that will copy the specified string to the caller's
// buffer and return the specified error when done regardless of buffer size and/or
// string length. The string is buffered until Read() is called, which will result in
// an error.
func NewBadStringReader(s string, e error) *BadStringReader {
	return &BadStringReader{
		buffer: []byte(s),
		error:  e,
	}
}

// Read copies the reader's buffer to the specified buffer but will always return the
// error that was specified when the reader was created.
func (r *BadStringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.buffer)
	return n, r.error
}
