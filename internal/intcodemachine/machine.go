package intcodemachine

import (
	"fmt"
	"io"
)

type Machine struct {
	memory []int
	tape   io.ReadSeeker
	ip     int
	err    error
}

func New(tape io.ReadSeeker) (*Machine, error) {
	var err error
	m := &Machine{
		tape: tape,
	}

	if err = m.LoadProgram(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Machine) Reset() error {
	m.ip = 0
	if _, err := m.tape.Seek(0, io.SeekStart); err != nil {
		return err
	}
	if err := m.LoadProgram(); err != nil {
		return err
	}
	return nil
}

func (m *Machine) SetInput(d0, d1 int) {
	m.memory[1], m.memory[2] = d0, d1
}

func (m *Machine) Run() error {
loop:
	for {
		switch m.memory[m.ip] {
		case 1:
			m.add()

		case 2:
			m.mul()

		case 99:
			break loop

		default:
			// Invalid instruction.
			m.err = fmt.Errorf("invalid instruction: %d", m.memory[m.ip])
		}
		if m.err != nil {
			return m.err
		}
	}
	return m.err
}

func (m *Machine) Output() int {
	return m.memory[0]
}

func (m *Machine) add() {
	m.memory[m.memory[m.ip+3]] = m.memory[m.memory[m.ip+1]] + m.memory[m.memory[m.ip+2]]
	m.ip += 4
}

func (m *Machine) mul() {
	m.memory[m.memory[m.ip+3]] = m.memory[m.memory[m.ip+1]] * m.memory[m.memory[m.ip+2]]
	m.ip += 4
}
