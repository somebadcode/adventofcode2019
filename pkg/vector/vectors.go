package vector

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Line struct {
	x0, x1, y0, y1 int64
}

type direction uint8

const (
	Left direction = iota
	Right
	Up
	Down
	InvalidDirection
)

var (
	ErrInvalidDirection = errors.New("invalid direction")
	ErrInvalidMagnitude = errors.New("invalid magnitude")
)

func Parse(r io.Reader) ([]Line, error) {
	var lines []Line

	scanner := bufio.NewScanner(r)
	scanner.Split(scanVectors)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return lines, err
		}

		// Convert token to upper case for easier comparison
		token := strings.ToUpper(scanner.Text())

		// Get prefix (direction).
		prefix, width := utf8.DecodeRuneInString(token)
		if prefix == utf8.RuneError {
			return lines, ErrInvalidDirection
		}

		// Decode direction from the prefix.
		direction, err := getDirection(prefix)
		if err != nil {
			return lines, ErrInvalidDirection
		}

		// Parse magnitude from what remains of the token..
		magnitude, err := strconv.ParseInt(token[width:], 10, 64)
		if err != nil {
			return lines, ErrInvalidMagnitude
		}

		if len(lines) == 0 {
			// First vector, append to empty slice.
			vector := attachVector(Line{}, direction, magnitude)
			lines = append(lines, vector)
		} else {
			// If slice is not empty then join the last vector with the new vector
			// to form the next vector.
			lines = append(lines, attachVector(lines[len(lines)-1], direction, magnitude))
		}
	}

	return lines, nil
}

func attachVector(line Line, d direction, m int64) Line {
    newLine := Line{
    	x0: line.x1,
    	x1: line.x1,
    	y0: line.y1,
    	y1: line.y1,
	}
	switch d {
	case Left:
		newLine.x1 -= m
	case Right:
		newLine.x1 += m
	case Up:
		newLine.y1 += m
	case Down:
		newLine.y1 -= m
	}
	return newLine
}

func ManhattanDistance(l1, l2 Line) int64 {
	return abs(l1.x1 - l2.x1) + abs(l1.y1 - l2.y1)
	// |x1 - x2| + |y1 - y2|
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func getDirection(direction rune) (direction, error) {
	switch unicode.ToUpper(direction) {
	case 'L':
		return Left, nil
	case 'R':
		return Right, nil
	case 'U':
		return Up, nil
	case 'D':
		return Down, nil
	}
	return InvalidDirection, ErrInvalidDirection
}
