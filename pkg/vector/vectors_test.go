package vector

import (
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
		want    []Line
		wantErr bool
	}{
		{
			args: args{
				r: strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
			},
			wantErr: false,
			want: []Line{
				{
					x0: 0,
					x1: 75,
					y0: 0,
					y1: 0,
				},
				{
					x0: 75,
					x1: 75,
					y0: 0,
					y1: -30,
				},
				{
					x0: 75,
					x1: 158,
					y0: -30,
					y1: -30,
				},
				{
					x0: 158,
					x1: 158,
					y0: -30,
					y1: 53,
				},
				{
					x0: 158,
					x1: 146,
					y0: 53,
					y1: 53,
				},
				{
					x0: 146,
					x1: 146,
					y0: 53,
					y1: 4,
				},
				{
					x0: 146,
					x1: 217,
					y0: 4,
					y1: 4,
				},
				{
					x0: 217,
					x1: 217,
					y0: 4,
					y1: 11,
				},
				{
					x0: 217,
					x1: 145,
					y0: 11,
					y1: 11,
				},
			},
		},
		{
			args: args{
				r: strings.NewReader("U62,R66,U55,R34,D71,R55,D58,R83"),
			},
			wantErr: false,
			want: []Line{
				{
					x0: 0,
					x1: 0,
					y0: 0,
					y1: 62,
				},
				{
					x0: 0,
					x1: 66,
					y0: 62,
					y1: 62,
				},
				{
					x0: 66,
					x1: 66,
					y0: 62,
					y1: 117,
				},
				{
					x0: 66,
					x1: 100,
					y0: 117,
					y1: 117,
				},
				{
					x0: 100,
					x1: 100,
					y0: 117,
					y1: 46,
				},
				{
					x0: 100,
					x1: 155,
					y0: 46,
					y1: 46,
				},
				{
					x0: 155,
					x1: 155,
					y0: 46,
					y1: -12,
				},
				{
					x0: 155,
					x1: 238,
					y0: -12,
					y1: -12,
				},
			},
		},
		{
			args: args{
				r: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			},
			wantErr: false,
			want: []Line{
				{
					x0: 0,
					x1: 98,
					y0: 0,
					y1: 0,
				},
				{
					x0: 98,
					x1: 98,
					y0: 0,
					y1: 47,
				},
				{
					x0: 98,
					x1: 124,
					y0: 47,
					y1: 47,
				},
				{
					x0: 124,
					x1: 124,
					y0: 47,
					y1: -16,
				},
				{
					x0: 124,
					x1: 157,
					y0: -16,
					y1: -16,
				},
				{
					x0: 157,
					x1: 157,
					y0: -16,
					y1: 71,
				},
				{
					x0: 157,
					x1: 95,
					y0: 71,
					y1: 71,
				},
				{
					x0: 95,
					x1: 95,
					y0: 71,
					y1: 51,
				},
				{
					x0: 95,
					x1: 128,
					y0: 51,
					y1: 51,
				},
				{
					x0: 128,
					x1: 128,
					y0: 51,
					y1: 104,
				},
				{
					x0: 128,
					x1: 179,
					y0: 104,
					y1: 104,
				},
			},
		},
		{
			args: args{
				r: strings.NewReader("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			},
			want: []Line{
				{
					x0: 0,
					x1: 0,
					y0: 0,
					y1: 98,
				},
				{
					x0: 0,
					x1: 91,
					y0: 98,
					y1: 98,
				},
				{
					x0: 91,
					x1: 91,
					y0: 98,
					y1: 78,
				},
				{
					x0: 91,
					x1: 107,
					y0: 78,
					y1: 78,
				},
				{
					x0: 107,
					x1: 107,
					y0: 78,
					y1: 11,
				},
				{
					x0: 107,
					x1: 147,
					y0: 11,
					y1: 11,
				},
				{
					x0: 147,
					x1: 147,
					y0: 11,
					y1: 18,
				},
				{
					x0: 147,
					x1: 162,
					y0: 18,
					y1: 18,
				},
				{
					x0: 162,
					x1: 162,
					y0: 18,
					y1: 24,
				},
				{
					x0: 162,
					x1: 169,
					y0: 24,
					y1: 24,
				},
			},
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
		v Line
		d direction
		m int64
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		{
			args: args{
				v: Line{
					x0: 0,
					x1: 10,
					y0: 0,
					y1: 0,
				},
				d: Right,
				m: 10,
			},
			want: Line{
				x0: 10,
				x1: 20,
				y0: 0,
				y1: 0,
			},
		},
		{
			args: args{
				v: Line{
					x0: 5,
					x1: 20,
					y0: -10,
					y1: 5,
				},
				d: Up,
				m: 50,
			},
			want: Line{
				x0: 20,
				x1: 20,
				y0: 5,
				y1: 55,
			},
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
		l1 Line
		l2 Line
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				l1: Line{},
				l2: Line{
					x0: 10,
					x1: 20,
					y0: 30,
					y1: 50,
				},
			},
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ManhattanDistance(tt.args.l1, tt.args.l2); got != tt.want {
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