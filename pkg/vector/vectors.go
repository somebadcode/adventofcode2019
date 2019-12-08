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

type Point struct {
	X, Y int64
}

type Vector struct {
	P, Q Point
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

func Parse(r io.Reader) ([]Vector, error) {
	var vectors []Vector

	scanner := bufio.NewScanner(r)
	scanner.Split(scanVectors)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		// Convert token to upper case for easier comparison
		token := strings.ToUpper(scanner.Text())

		// Get prefix (direction).
		prefix, width := utf8.DecodeRuneInString(token)

		// Decode direction from the prefix.
		direction, err := getDirection(prefix)
		if err != nil {
			return nil, ErrInvalidDirection
		}

		// Parse magnitude from what remains of the token..
		magnitude, err := strconv.ParseInt(token[width:], 10, 64)
		if err != nil {
			return nil, ErrInvalidMagnitude
		}

		// If there are no vectors, simply put the first one into the slice.
		if len(vectors) == 0 {
			vectors = []Vector{attachVector(Vector{}, direction, magnitude)}
			continue
		}

		// Append the vector to the last vector.
		vectors = append(vectors, attachVector(vectors[len(vectors)-1], direction, magnitude))
	}

	return vectors, nil
}

func attachVector(vector Vector, d direction, m int64) Vector {
	v := Vector{P: vector.Q, Q: vector.Q}
	switch d {
	case Left:
		v.Q.X -= m
	case Right:
		v.Q.X += m
	case Up:
		v.Q.Y += m
	case Down:
		v.Q.Y -= m
	}
	return v
}

func ManhattanDistance(v1, v2 Vector) int64 {
	return abs(v1.Q.X-v2.Q.X) + abs(v1.Q.Y-v2.Q.Y)
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
