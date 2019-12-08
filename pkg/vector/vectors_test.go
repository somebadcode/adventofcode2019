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
				{P: Point{0, 0}, Q: Point{75, 0}},
				{P: Point{75, 0}, Q: Point{75, -30}},
				{P: Point{75, -30}, Q: Point{158, -30}},
				{P: Point{158, -30}, Q: Point{158, 53}},
				{P: Point{158, 53}, Q: Point{146, 53}},
				{P: Point{146, 53}, Q: Point{146, 4}},
				{P: Point{146, 4}, Q: Point{217, 4}},
				{P: Point{217, 4}, Q: Point{217, 11}},
				{P: Point{217, 11}, Q: Point{145, 11}},
			},
		},
		{
			args: args{
				r: strings.NewReader("U62,R66,U55,R34,D71,R55,D58,R83"),
			},
			wantErr: false,
			want: []Vector{
				{P: Point{0, 0}, Q: Point{0, 62}},
				{P: Point{0, 62}, Q: Point{66, 62}},
				{P: Point{66, 62}, Q: Point{66, 117}},
				{P: Point{66, 117}, Q: Point{100, 117}},
				{P: Point{100, 117}, Q: Point{100, 46}},
				{P: Point{100, 46}, Q: Point{155, 46}},
				{P: Point{155, 46}, Q: Point{155, -12}},
				{P: Point{155, -12}, Q: Point{238, -12}},
			},
		},
		{
			args: args{
				r: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			},
			wantErr: false,
			want: []Vector{
				{P: Point{0, 0}, Q: Point{98, 0}},
				{P: Point{98, 0}, Q: Point{98, 47}},
				{P: Point{98, 47}, Q: Point{124, 47}},
				{P: Point{124, 47}, Q: Point{124, -16}},
				{P: Point{124, -16}, Q: Point{157, -16}},
				{P: Point{157, -16}, Q: Point{157, 71}},
				{P: Point{157, 71}, Q: Point{95, 71}},
				{P: Point{95, 71}, Q: Point{95, 51}},
				{P: Point{95, 51}, Q: Point{128, 51}},
				{P: Point{128, 51}, Q: Point{128, 104}},
				{P: Point{128, 104}, Q: Point{179, 104}},
			},
		},
		{
			args: args{
				r: strings.NewReader("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			},
			want: []Vector{
				{P: Point{0, 0}, Q: Point{0, 98}},
				{P: Point{0, 98}, Q: Point{91, 98}},
				{P: Point{91, 98}, Q: Point{91, 78}},
				{P: Point{91, 78}, Q: Point{107, 78}},
				{P: Point{107, 78}, Q: Point{107, 11}},
				{P: Point{107, 11}, Q: Point{147, 11}},
				{P: Point{147, 11}, Q: Point{147, 18}},
				{P: Point{147, 18}, Q: Point{162, 18}},
				{P: Point{162, 18}, Q: Point{162, 24}},
				{P: Point{162, 24}, Q: Point{169, 24}},
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
		want    direction
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
			want:    InvalidDirection,
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
		d direction
		m int64
	}
	tests := []struct {
		name string
		args args
		want Vector
	}{
		{
			args: args{
				v: Vector{P: Point{0, 0}, Q: Point{10, 0}},
				d: Right,
				m: 10,
			},
			want: Vector{P: Point{10, 0}, Q: Point{20, 0}},
		},
		{
			args: args{
				v: Vector{P: Point{5, -10}, Q: Point{20, 5}},
				d: Up,
				m: 50,
			},
			want: Vector{P: Point{20, 5}, Q: Point{20, 55}},
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
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				v1: Vector{P: Point{0, 0}, Q: Point{0, 0}},
				v2: Vector{P: Point{10, 30}, Q: Point{20, 50}},
			},
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ManhattanDistance(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("ManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		x int64
	}
	tests := []struct {
		name string
		args args
		want int64
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
