package day3

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/somebadcode/adventofcode2019/internal/solver"
	"github.com/somebadcode/adventofcode2019/pkg/vector"
	"github.com/spf13/viper"
	"io"
	"math"
)

type Solver struct {
	config *viper.Viper
}

func New(config *viper.Viper) solver.Solver {
	return &Solver{
		config: config,
	}
}

func (s *Solver) Solve(r io.ReadSeeker) []string {
	s1 := s.PartOne(r)

	_, err := r.Seek(0, io.SeekStart)
	if err != nil {
		return []string{err.Error(), ""}
	}

	return []string{s1, s.PartTwo(r)}
}

func (s *Solver) PartOne(r io.ReadSeeker) string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var wires [][]vector.Vector
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err.Error()
		}
		wire, err := vector.Parse(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			return err.Error()
		}
		wires = append(wires, wire)
	}

	if len(wires) != 2 {
		return "expected 2 wires"
	}

	nearest := math.MaxFloat64
	for _, w1 := range wires[0] {
		for _, w2 := range wires[1] {
			if p := vector.GetIntersection(w1, w2); p != nil {
				if p.X == 0 && p.Y == 0 {
					continue
				}
				distance := vector.ManhattanDistance(vector.Point{}, *p)
				if distance < nearest {
					nearest = distance
				}
			}
		}
	}

	return fmt.Sprintf("%.0f", nearest)
}

func (s *Solver) PartTwo(r io.ReadSeeker) string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var wires [][]vector.Vector
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err.Error()
		}
		wire, err := vector.Parse(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			return err.Error()
		}
		wires = append(wires, wire)
	}

	if len(wires) != 2 {
		return "expected 2 wires"
	}

	intersections := make(map[vector.Point]float64)
	for _, w1 := range wires[0] {
		for _, w2 := range wires[1] {
			if p := vector.GetIntersection(w1, w2); p != nil {
				if p.X == 0 && p.Y == 0 {
					continue
				}
				intersections[*p] = math.SmallestNonzeroFloat64
			}
		}
	}

	for p := range intersections {
		for _, wire := range wires {
			for _, w := range wire {
				if vector.IsInSegment(p, w, 0.01) {
					intersections[p] += vector.Distance(w.P, p)
					break
				} else {
					intersections[p] += w.M
				}
			}
		}
	}

	distance := math.MaxFloat64
	for _, v := range intersections {
		if v < distance {
			distance = v
		}
	}

	return fmt.Sprintf("%.0f", distance)
}
