package vector

import (
	"bufio"
	"errors"
	"io"
	"math"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Point struct {
	X, Y float64
}

type Vector struct {
	P, Q Point
	D    Angle
	M    float64
}

type Angle float64

const (
	Up    Angle = 0.0
	Right Angle = 90.0
	Down  Angle = 180.0
	Left  Angle = 270.0
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
		magnitude, err := strconv.ParseFloat(token[width:], 64)
		if err != nil {
			return nil, ErrInvalidMagnitude
		}

		// If there are no vectors in the slice, simply put the first one into the slice.
		if len(vectors) == 0 {
			vectors = []Vector{attachVector(Vector{}, direction, magnitude)}
			continue
		}

		// Append the vector to the last vector.
		vectors = append(vectors, attachVector(vectors[len(vectors)-1], direction, magnitude))
	}
	return vectors, nil
}

func attachVector(vector Vector, d Angle, m float64) Vector {
	v := Vector{P: vector.Q, Q: vector.Q, D: d, M: m}
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

func GetIntersection(v1, v2 Vector) *Point {
	// Get the intersection of the line segments of intersecting lines created
	// by the vectors. Return nil if there's no intersection.
	// http://www.cs.swan.ac.uk/~cssimon/line_intersection.html
	// http://www-cs.ccny.cuny.edu/~wolberg/capstone/intersection/Intersection%20point%20of%20two%20lines.html

	x1, x2 := v1.P.X, v1.Q.X
	x3, x4 := v2.P.X, v2.Q.X
	y1, y2 := v1.P.Y, v1.Q.Y
	y3, y4 := v2.P.Y, v2.Q.Y

	// Denominator.
	da := (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)

	// Denominator b.
	db := (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)

	// If either denominator is zero, the segments do not intersect.
	if da == 0 {
		return nil
	}

	// Solve the two unknowns (a and b)
	ta := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)) / da
	tb := ((x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)) / db

	// If 0 ≤ ta ≤ 1 and 0 ≤ tb ≤ 1 then they intersect.
	if (0 <= ta && ta <= 1) && (0 <= tb && tb <= 1) {
		x := x1 + ta*(x2-x1)
		y := y1 + ta*(y2-y1)
		return &Point{x, y}
	}
	return nil
}

func ManhattanDistance(p1, p2 Point) float64 {
	return math.Abs(p1.X-p2.X) + math.Abs(p1.Y-p2.Y)
}

func getDirection(direction rune) (Angle, error) {
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
	return 0.0, ErrInvalidDirection
}

func IsInSegment(p Point, v Vector, epsilon float64) bool {
	crossProduct := (p.Y-v.P.Y)*(v.Q.X-v.P.X) - (p.X-v.P.X)*(v.Q.Y-v.P.Y)
	if math.Abs(crossProduct) > epsilon {
		return false
	}

	dotProduct := (p.X-v.P.X)*(v.Q.X-v.P.X) + (p.Y-v.P.Y)*(v.Q.Y-v.P.Y)
	if dotProduct < 0 {
		return false
	}

	squaredLength := (v.Q.X-v.P.X)*(v.Q.X-v.P.X) + (v.Q.Y-v.P.Y)*(v.Q.Y-v.P.Y)
	return squaredLength >= dotProduct
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}
