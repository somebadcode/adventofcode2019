package solver

import "io"

type Parts interface {
	PartOne(r io.ReadSeeker) string
	PartTwo(r io.ReadSeeker) string
}
