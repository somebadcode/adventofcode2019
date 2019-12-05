package day1

import (
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
		want *Solver
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
				r: strings.NewReader("garbage"),
			},
			want: "strconv.ParseInt: parsing \"garbage\": invalid syntax",
		},
		{
			args: args{
				r: badreadseeker.New(strings.NewReader("666"), io.ErrUnexpectedEOF, badreadseeker.Read),
			},
			want: io.ErrUnexpectedEOF.Error(),
		},
		{
			args: args{
				r: strings.NewReader("12"),
			},
			want: "2",
		},
		{
			args: args{
				r: strings.NewReader("14"),
			},
			want: "2",
		},
		{
			args: args{
				r: strings.NewReader("1969"),
			},
			want: "654",
		},
		{
			args: args{
				r: strings.NewReader("100756"),
			},
			want: "33583",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solver{
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
				r: strings.NewReader("garbage"),
			},
			want: "strconv.ParseInt: parsing \"garbage\": invalid syntax",
		},
		{
			args: args{
				r: badreadseeker.New(strings.NewReader("666"), io.ErrUnexpectedEOF, badreadseeker.Read),
			},
			want: io.ErrUnexpectedEOF.Error(),
		},
		{
			args: args{
				r: strings.NewReader("14"),
			},
			want: "2",
		},
		{
			args: args{
				r: strings.NewReader("1969"),
			},
			want: "966",
		},
		{
			args: args{
				r: strings.NewReader("100756"),
			},
			want: "50346",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solver{
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
				r: strings.NewReader("100756"),
			},
			want: []string{"33583", "50346"},
		},
		{
			args: args{
				r: badreadseeker.New(strings.NewReader("100756"), io.ErrShortBuffer, badreadseeker.Seek),
			},
			want: []string{io.ErrShortBuffer.Error(), ""},
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
