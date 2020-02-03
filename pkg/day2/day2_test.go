package day2

import (
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
				r: strings.NewReader("1,12,2,0,99,0,0,0,0,0,0,0,1"),
			},
			want: "3",
		},
		{
			args: args{
				r: strings.NewReader("2,12,2,3,2,1,4,0,99,0,0,0,0"),
			},
			want: "24",
		},
		{
			args: args{
				r: strings.NewReader("2,12,2,0,2,12,8,0,99,0,0,0,99"),
			},
			want: "9801",
		},
		{
			args: args{
				r: strings.NewReader("1,12,2,0,99,5,6,0,99,0,0,0,28"),
			},
			want: "30",
		},
		{
			args: args{
				r: strings.NewReader("1,12,2,0,0,0,0,0,0,0,0,0,0"),
			},
			want: "invalid instruction: 0",
		},
		{
			args: args{
				r: strings.NewReader("1,12,2,0,mio,99,0,0,0,0,0,0,0,0,0,0"),
			},
			want: "strconv.Atoi: parsing \"mio\": invalid syntax",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solver{
				config: tt.fields.config,
			}
			s.config = viper.New()
			s.config.Set("part1.input", []int{12, 2})
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
		input  int
		want   string
	}{
		{
			fields: fields{
				config: viper.New(),
			},
			args: args{
				r: strings.NewReader("2,0,0,0,99,3"),
			},
			input: 9,
			want:  "100 \u2715 5 + 5 = 505",
		},
		{
			fields: fields{
				config: viper.New(),
			},
			args: args{
				r: testdatafromfile.From("day2.txt"),
			},
			input: 1000,
			want:  "unexpected end of part two",
		},
		{
			fields: fields{
				config: viper.New(),
			},
			args: args{
				r: badreadseeker.New(testdatafromfile.From("day2.txt"), io.ErrUnexpectedEOF, badreadseeker.Seek),
			},
			input: 1000,
			want:  "unexpected EOF",
		},
		{
			fields: fields{
				config: viper.New(),
			},
			args: args{
				r: testdatafromfile.From("day2.txt"),
			},
			input: 19690720,
			want:  "100 \u2715 49 + 25 = 4925",
		},
		{
			fields: fields{
				config: viper.New(),
			},
			args: args{
				r: strings.NewReader("1,12,2,0,0,0,0,0,0,0,0,0,0"),
			},
			input: 0,
			want: "invalid instruction: 0",
		},
		{
			fields: fields{
				config: viper.New(),
			},
			args: args{
				r: strings.NewReader("1,12,2,0,mio,99,0,0,0,0,0,0,0,0,0,0"),
			},
			input: 0,
			want: "strconv.Atoi: parsing \"mio\": invalid syntax",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solver{
				config: tt.fields.config,
			}
			s.config.Set("part2.input", tt.input)
			if got := s.PartTwo(tt.args.r); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
