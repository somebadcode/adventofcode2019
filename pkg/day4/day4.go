package day4

import (
	"errors"
	"github.com/somebadcode/adventofcode2019/internal/solver"
	"github.com/spf13/viper"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
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
	pair, err := parseRange(r)
	if err != nil {
		return err.Error()
	}

	var counter uint64
	for i := pair[0]; i <= pair[1]; i++ {
		s := strconv.FormatInt(i, 10)
		if validatePasswordOne(s) {
			counter++
		}
	}

	return strconv.FormatUint(counter, 10)
}

func (s Solver) PartTwo(r io.ReadSeeker) string {
	pair, err := parseRange(r)
	if err != nil {
		return err.Error()
	}

	var counter uint64
	for i := pair[0]; i <= pair[1]; i++ {
		s := strconv.FormatInt(i, 10)
		if validatePasswordTwo(s) {
			counter++
		}
	}

	return strconv.FormatUint(counter, 10)
}

func parseRange(r io.ReadSeeker) ([]int64, error) {
	b := make([]byte, 50)
	size, err := r.Read(b)
	if err != nil {
		return nil, err
	} else if size <= 3 {
		return nil, errors.New("input too short")
	}

	input := string(b[:size])

	if valid := utf8.Valid(b[:size]); !valid {
		return nil, errors.New("input is not a valid string")
	}

	pair := strings.Split(input, "-")
	if len(pair) != 2 {
		return nil, errors.New("input is not a string with one separator '-'")
	}

	if len(pair[0]) < 2 || len(pair[1]) < 2 {
		return nil, errors.New("input pairs are too short")
	}

	intPair := make([]int64, 2)
	intPair[0], err = strconv.ParseInt(pair[0], 10, 64)
	if err != nil {
		return nil, err
	}
	intPair[1], err = strconv.ParseInt(pair[1], 10, 64)
	if err != nil {
		return nil, err
	}

	return intPair, nil
}

func validatePasswordOne(s string) bool {
	var hasDouble bool
	var width int
	var r, prev rune
	for i := 0; i < len(s); i++ {
		r, width = utf8.DecodeRuneInString(s[i:])
		if r < prev {
			return false
		}
		if repeats := runeRepeats(s[i:], r); repeats >= 2 {
			hasDouble = true
			i += repeats*width - width*1
		}
		prev = r
	}
	return hasDouble
}

func validatePasswordTwo(s string) bool {
	var hasDouble bool
	var width int
	var r, prev rune
	for i := 0; i < len(s); i++ {
		r, width = utf8.DecodeRuneInString(s[i:])
		if r < prev {
			return false
		}
		if repeats := runeRepeats(s[i:], r); repeats == 2 {
			hasDouble = true
			i += repeats*width - 1*width
		} else if repeats > 2 {
			i += repeats*width - 1*width
		}
		prev = r
	}
	return hasDouble
}

func runeRepeats(s string, r rune) int {
	var count int
	for _, c := range s {
		if c != r {
			break
		}
		count++
	}
	return count
}
