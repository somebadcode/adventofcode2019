package badreadseeker

import (
	"io"
)

type when int8

const (
	Read when = 1 << iota
	Seek
)

// BadReadSeeker implements io.Reader and is meant to return an error when the
// internal buffer has been copied or the receiving buffer is full.
type BadReadSeeker struct {
	readSeeker io.ReadSeeker
	when       when
	error      error
}

// NewBadReader returns a reader that will copy the specified string to the caller's
// buffer and return the specified error when done regardless of buffer size and/or
// string length. The string is buffered until Read() is called, which will result in
// an error.
func New(r io.ReadSeeker, e error, when when) *BadReadSeeker {
	return &BadReadSeeker{
		readSeeker: r,
		when:       when,
		error:      e,
	}
}

// Read will read from the ReadSeeker nad return an error if specified.
// If the underlying reader returns an error then Read will return the same error.
func (r *BadReadSeeker) Read(p []byte) (n int, err error) {
	n, err = r.readSeeker.Read(p)
	if err != nil {
		return n, err
	}
	if r.when&Read != 0 {
		return n, r.error
	}
	return n, nil
}

// Seek will seek the underlying ReadSeeker and return an error if specified.
// If the underlying seeker returns an error then Seek will return the same error.
func (r *BadReadSeeker) Seek(offset int64, whence int) (int64, error) {
	n, err := r.readSeeker.Seek(offset, whence)
	if err != nil {
		return n, err
	}
	if r.when&Seek != 0 {
		return n, r.error
	}
	return n, nil
}
