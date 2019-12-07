package solver

import "io"

type Solver interface {
	Solve(r io.ReadSeeker) []string
}
