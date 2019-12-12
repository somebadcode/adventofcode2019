package day3

import (
	"github.com/somebadcode/adventofcode2019/internal/solver"
	"github.com/somebadcode/adventofcode2019/internal/testdatafromfile"
	"github.com/somebadcode/adventofcode2019/pkg/badreadseeker"
	"github.com/spf13/viper"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		config *viper.Viper
	}
	tests := []struct {
		name string
		args args
		want solver.Solver
	}{
		{
			args: args{
				config: viper.GetViper(),
			},
			want: &Solver{
				config: viper.GetViper(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolver_PartOne(t *testing.T) {
	type fields struct {
		config *viper.Viper
	}
	type args struct {
		r io.ReadSeeker
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			args: args{
				r: strings.NewReader("R8,U5,L5,D3\nU7,R6,D4,L4"),
			},
			want: "6",
		},
		{
			args: args{
				r: strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"),
			},
			want: "159",
		},
		{
			args: args{
				r: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			},
			want: "135",
		},
		{
			args: args{
				r: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			},
			want: "expected 2 wires",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{
				config: tt.fields.config,
			}
			if got := s.PartOne(tt.args.r); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolver_PartTwo(t *testing.T) {
	type fields struct {
		config *viper.Viper
	}
	type args struct {
		r io.ReadSeeker
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			args: args{
				r: strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"),
			},
			want: "610",
		},
		{
			args: args{
				r: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			},
			want: "410",
		},
		{
			args: args{
				r: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			},
			want: "expected 2 wires",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{
				config: tt.fields.config,
			}
			if got := s.PartTwo(tt.args.r); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolver_Solve(t *testing.T) {
	type fields struct {
		config *viper.Viper
	}
	type args struct {
		r io.ReadSeeker
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			args: args{
				r: strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"),
			},
			want: []string{"159", "610"},
		},
		{
			args: args{
				r: strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			},
			want: []string{"135", "410"},
		},
		{
			args: args{
				r: testdatafromfile.From("day3.txt"),
			},
			want: []string{"1431", "48012"},
		},
		{
			args: args{
				r: badreadseeker.New(strings.NewReader("L90,U40"), io.ErrUnexpectedEOF, badreadseeker.Read),
			},
			want: []string{io.ErrUnexpectedEOF.Error(), io.ErrUnexpectedEOF.Error()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{
				config: tt.fields.config,
			}
			if got := s.Solve(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
