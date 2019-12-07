package day1

import (
	"bufio"
	"github.com/somebadcode/adventofcode2019/internal/solver"
	"github.com/spf13/viper"
	"io"
	"strconv"
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
	solution := s.PartOne(r)

	_, err := r.Seek(0, io.SeekStart)
	if err != nil {
		return []string{err.Error(), ""}
	}

	return []string{solution, s.PartTwo(r)}
}

func (s Solver) PartOne(r io.ReadSeeker) string {
	var fuelQuantity int64

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err.Error()
		}

		mass, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return err.Error()
		}

		fuelQuantity += (mass / 3) - 2
	}

	return strconv.FormatInt(fuelQuantity, 10)
}

func (s Solver) PartTwo(r io.ReadSeeker) string {
	var fuelQuantity int64

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var totalFuelRequirement int64

		if err := scanner.Err(); err != nil {
			return err.Error()
		}

		mass, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return err.Error()
		}

		baseFuelRequirement := (mass / 3) - 2
		q := baseFuelRequirement
		for {
			totalFuelRequirement += q
			q = (q / 3) - 2
			if q <= 0 {
				break
			}
		}

		fuelQuantity += totalFuelRequirement
	}

	return strconv.FormatInt(fuelQuantity, 10)
}
