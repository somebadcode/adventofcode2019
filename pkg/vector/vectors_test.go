package vector

import (
	"github.com/somebadcode/adventofcode2019/pkg/badreadseeker"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []Vector
		wantErr bool
	}{
		{
			args: args{
				r: strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
			},
			wantErr: false,
			want: []Vector{
				{P: Point{0, 0}, Q: Point{75, 0}, D: Right, M: 75},
				{P: Point{75, 0}, Q: Point{75, -30}, D: Down, M: 30},
				{P: Point{75, -30}, Q: Point{158, -30}, D: Right, M: 83},
				{P: Point{158, -30}, Q: Point{158, 53}, D: Up, M: 83},
				{P: Point{158, 53}, Q: Point{146, 53}, D: Left, M: 12},
				{P: Point{146, 53}, Q: Point{146, 4}, D: Down, M: 49},
				{P: Point{146, 4}, Q: Point{217, 4}, D: Right, M: 71},
				{P: Point{217, 4}, Q: Point{217, 11}, D: Up, M: 7},
				{P: Point{217, 11}, Q: Point{145, 11}, D: Left, M: 72},
			},
		},
		{
			args: args{
				r: strings.NewReader("U62,R66,U55,R34,D71,R55,D58,R83"),
			},
			wantErr: false,
			want: []Vector{
				{P: Point{0, 0}, Q: Point{0, 62}, D: Up, M: 62},
				{P: Point{0, 62}, Q: Point{66, 62}, D: Right, M: 66},
				{P: Point{66, 62}, Q: Point{66, 117}, D: Up, M: 55},
				{P: Point{66, 117}, Q: Point{100, 117}, D: Right, M: 34},
				{P: Point{100, 117}, Q: Point{100, 46}, D: Down, M: 71},
				{P: Point{100, 46}, Q: Point{155, 46}, D: Right, M: 55},
				{P: Point{155, 46}, Q: Point{155, -12}, D: Down, M: 58},
				{P: Point{155, -12}, Q: Point{238, -12}, D: Right, M: 83},
			},
		},
		{
			args: args{
				r: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			},
			wantErr: false,
			want: []Vector{
				{P: Point{0, 0}, Q: Point{98, 0}, D: Right, M: 98},
				{P: Point{98, 0}, Q: Point{98, 47}, D: Up, M: 47},
				{P: Point{98, 47}, Q: Point{124, 47}, D: Right, M: 26},
				{P: Point{124, 47}, Q: Point{124, -16}, D: Down, M: 63},
				{P: Point{124, -16}, Q: Point{157, -16}, D: Right, M: 33},
				{P: Point{157, -16}, Q: Point{157, 71}, D: Up, M: 87},
				{P: Point{157, 71}, Q: Point{95, 71}, D: Left, M: 62},
				{P: Point{95, 71}, Q: Point{95, 51}, D: Down, M: 20},
				{P: Point{95, 51}, Q: Point{128, 51}, D: Right, M: 33},
				{P: Point{128, 51}, Q: Point{128, 104}, D: Up, M: 53},
				{P: Point{128, 104}, Q: Point{179, 104}, D: Right, M: 51},
			},
		},
		{
			args: args{
				r: strings.NewReader("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			},
			want: []Vector{
				{P: Point{0, 0}, Q: Point{0, 98}, D: Up, M: 98},
				{P: Point{0, 98}, Q: Point{91, 98}, D: Right, M: 91},
				{P: Point{91, 98}, Q: Point{91, 78}, D: Down, M: 20},
				{P: Point{91, 78}, Q: Point{107, 78}, D: Right, M: 16},
				{P: Point{107, 78}, Q: Point{107, 11}, D: Down, M: 67},
				{P: Point{107, 11}, Q: Point{147, 11}, D: Right, M: 40},
				{P: Point{147, 11}, Q: Point{147, 18}, D: Up, M: 7},
				{P: Point{147, 18}, Q: Point{162, 18}, D: Right, M: 15},
				{P: Point{162, 18}, Q: Point{162, 24}, D: Up, M: 6},
				{P: Point{162, 24}, Q: Point{169, 24}, D: Right, M: 7},
			},
		},
		{
			args: args{
				r: badreadseeker.New(strings.NewReader("L20,U90"), io.ErrUnexpectedEOF, badreadseeker.Read),
			},
			wantErr: true,
		},
		{
			args: args{
				r: strings.NewReader("V300,U90"),
			},
			wantErr: true,
		},
		{
			args: args{
				r: strings.NewReader("R30m,U90"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDirection(t *testing.T) {
	type args struct {
		direction rune
	}
	tests := []struct {
		name    string
		args    args
		want    Angle
		wantErr bool
	}{
		{
			name: "up",
			args: args{
				direction: 'U',
			},
			want:    Up,
			wantErr: false,
		},
		{
			name: "down",
			args: args{
				direction: 'D',
			},
			want:    Down,
			wantErr: false,
		},
		{
			name: "left",
			args: args{
				direction: 'L',
			},
			want:    Left,
			wantErr: false,
		},
		{
			name: "right",
			args: args{
				direction: 'R',
			},
			want:    Right,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				direction: 'X',
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDirection(tt.args.direction)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDirection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getDirection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_attachVector(t *testing.T) {
	type args struct {
		v Vector
		d Angle
		m float64
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{
			args: args{
				v: Vector{P: Point{0, 0}, Q: Point{10, 0}, D: Right, M: 10},
				d: Right,
				m: 10,
			},
			want: Vector{P: Point{10, 0}, Q: Point{20, 0}, D: Right, M: 10},
		},
		{
			args: args{
				v: Vector{P: Point{5, 5}, Q: Point{20, 5}, D: Up, M: 50},
				d: Up,
				m: 50,
			},
			want: Vector{P: Point{20, 5}, Q: Point{20, 55}, D: Up, M: 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := attachVector(tt.args.v, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attachVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManhattanDistance(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				p1: Point{0, 0},
				p2: Point{3, 3},
			},
			want: 6,
		},
		{
			args: args{
				p1: Point{0, 0},
				p2: Point{10, 30},
			},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ManhattanDistance(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("ManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				x: 200,
			},
			want: 200,
		},
		{
			args: args{
				x: -36871,
			},
			want: 36871,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetIntersection(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name string
		args args
		want *Point
	}{
		{
			name: "plus",
			args: args{
				v1: Vector{P: Point{30, 0}, Q: Point{30, 60}, D: Up, M: 30},
				v2: Vector{P: Point{0, 30}, Q: Point{60, 30}, D: Right, M: 30},
			},
			want: &Point{30, 30},
		},
		{
			name: "T",
			args: args{
				v1: Vector{P: Point{30, 0}, Q: Point{30, 60}, D: Up, M: 60},
				v2: Vector{P: Point{0, 60}, Q: Point{60, 60}, D: Right, M: 30},
			},
			want: &Point{30, 60},
		},
		{
			name: "T upside down",
			args: args{
				v1: Vector{P: Point{30, 0}, Q: Point{30, 60}, D: Up, M: 60},
				v2: Vector{P: Point{0, 0}, Q: Point{60, 0}, D: Right, M: 60},
			},
			want: &Point{30, 0},
		},
		{
			name: "T gap",
			args: args{
				v1: Vector{P: Point{30, 0}, Q: Point{30, 60}, D: Up, M: 30},
				v2: Vector{P: Point{0, 61}, Q: Point{60, 61}, D: Right, M: 30},
			},
			want: nil,
		},
		{
			name: "parallel",
			args: args{
				v1: Vector{P: Point{30, 10}, Q: Point{60, 10}, D: Up, M: 30},
				v2: Vector{P: Point{40, 10}, Q: Point{70, 10}, D: Up, M: 30},
			},
			want: nil,
		},
		{
			name: "collinear",
			args: args{
				v1: Vector{P: Point{0, 10}, Q: Point{60, 10}, D: Right, M: 60},
				v2: Vector{P: Point{60, 10}, Q: Point{90, 10}, D: Right, M: 60},
			},
			want: nil,
		},
		{
			name: "flat",
			args: args{
				v1: Vector{P: Point{0, 0}, Q: Point{100, 0}, D: Right, M: 100},
				v2: Vector{P: Point{50, 50}, Q: Point{50, -50}, D: Down, M: 100},
			},
			want: &Point{50, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIntersection(tt.args.v1, tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				p1: Point{0, 3},
				p2: Point{0, 9},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IsInSegment(t *testing.T) {
	type args struct {
		p       Point
		v       Vector
		epsilon float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				p: Point{10, 10},
				v: Vector{P: Point{10, 0}, Q: Point{10, 30}, D: Up, M: 30},
			},
			want: true,
		},
		{
			args: args{
				p: Point{10, 10},
				v: Vector{P: Point{0, 0}, Q: Point{0, 20}, D: Up, M: 20},
			},
			want: false,
		},
		{
			args: args{
				p: Point{10, 10},
				v: Vector{P: Point{-200, 5}, Q: Point{200, 9}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInSegment(tt.args.p, tt.args.v, tt.args.epsilon); got != tt.want {
				t.Errorf("IsInSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}
