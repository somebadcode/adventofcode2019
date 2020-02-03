package day2

import (
	"fmt"
	"github.com/somebadcode/adventofcode2019/internal/intcodemachine"
	"github.com/somebadcode/adventofcode2019/internal/solver"
	"github.com/spf13/viper"
	"io"
	"strconv"
)

type Solver struct {
	config *viper.Viper
}

func New(config *viper.Viper) solver.Parts {
	return &Solver{
		config: config,
	}
}

func (s Solver) PartOne(r io.ReadSeeker) string {
	// Load machine.
	m, err := intcodemachine.New(r)
	if err != nil {
		return err.Error()
	}

	// Get input parameters.
	input := s.config.GetIntSlice("part1.input")

	// Set input parameters.
	m.SetInput(input[0], input[1])

	err = m.Run()
	if err != nil {
		return err.Error()
	}
	return strconv.Itoa(m.Output())
}

func (s *Solver) PartTwo(r io.ReadSeeker) string {
	// Load machine.
	m, err := intcodemachine.New(r)
	if err != nil {
		return err.Error()
	}

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb <= noun; verb++ {
			// Set input parameters.
			m.SetInput(noun, verb)

			// Run!
			if err := m.Run(); err != nil {
				return err.Error()
			}

			// Get the integer we should be looking for to find the noun and verb.
			input := s.config.GetInt("part2.input")

			// Check if we got the right output.
			if m.Output() == input {
				return fmt.Sprintf("100 \u2715 %d + %d = %d", noun, verb, 100*noun+verb)
			}

			// Reset the machine.
			if err = m.Reset(); err != nil {
				return err.Error()
			}
		}
	}
	return "unexpected end of part two"
}
