package day2

import (
	"github.com/somebadcode/adventofcode2019/internal/intcodemachine"
	"io"
	"strconv"
)

func PartOne(r io.ReadSeeker) string {
	// Load machine.
	m, err := intcodemachine.New(r)
	if err != nil {
		return err.Error()
	}

	// Set input parameters.
	m.SetInput(12, 2)
	err = m.Run()
	if err != nil {
		return err.Error()
	}
	return strconv.Itoa(m.Output())
}

func PartTwo(r io.ReadSeeker) string {
	// Load machine.
	m, err := intcodemachine.New(r)
	if err != nil {
		return err.Error()
	}

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			// Set input parameters.
			m.SetInput(noun, verb)

			// Run!
			if err := m.Run(); err != nil {
				return err.Error()
			}
			// Check if we got the right output.
			if m.Output() == 19690720 {
				return strconv.Itoa(100*noun + verb)
			}
			// Reset the machine.
			if err = m.Reset(); err != nil {
				return err.Error()
			}
		}
	}
	return "unexpected end of part two"
}
