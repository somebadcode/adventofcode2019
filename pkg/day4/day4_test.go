package day4

import (
	"bytes"
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
		r io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			args: args{
				r: strings.NewReader("111110-111111"),
			},
			want: "1",
		},
		{
			args: args{
				r: strings.NewReader("223450-223450"),
			},
			want: "0",
		},
		{
			args: args{
				r: strings.NewReader("123789-123790"),
			},
			want: "0",
		},
		{
			args: args{
				r: badreadseeker.New(strings.NewReader("100-900"), io.ErrShortBuffer, badreadseeker.Read),
			},
			want: io.ErrShortBuffer.Error(),
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
		r io.Reader
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			args: args{
				r: strings.NewReader("112345-112347"),
			},
			want: "3",
		},
		{
			args: args{
				r: strings.NewReader("123458-123458"),
			},
			want: "0",
		},
		{
			args: args{
				r: strings.NewReader("123789-126703"),
			},
			want: "66",
		},
		{
			args: args{
				r: badreadseeker.New(strings.NewReader("100-900"), io.ErrShortBuffer, badreadseeker.Read),
			},
			want: io.ErrShortBuffer.Error(),
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
				r: strings.NewReader("12240-12999"),
			},
			want: []string{"70", "63"},
		},
		{
			args: args{
				r: strings.NewReader("1-3"),
			},
			want: []string{"input too short", "input too short"},
		},
		{
			args: args{
				r: bytes.NewReader([]byte{0xff, 0xfe, 0xfd, 0xfc, 0xfb}),
			},
			want: []string{"input is not a valid string", "input is not a valid string"},
		},
		{
			args: args{
				r: strings.NewReader("weeee"),
			},
			want: []string{"input is not a string with one separator '-'", "input is not a string with one separator '-'"},
		},
		{
			args: args{
				r: strings.NewReader("1-333"),
			},
			want: []string{"input pairs are too short", "input pairs are too short"},
		},
		{
			args: args{
				r: badreadseeker.New(strings.NewReader("100-900"), io.ErrShortBuffer, badreadseeker.Read),
			},
			want: []string{io.ErrShortBuffer.Error(), io.ErrShortBuffer.Error()},
		},
		{
			args: args{
				r: badreadseeker.New(strings.NewReader("100-900"), io.ErrShortBuffer, badreadseeker.Seek),
			},
			want: []string{io.ErrShortBuffer.Error(), ""},
		},
		{
			args: args{
				r: testdatafromfile.From("day4.txt"),
			},
			want: []string{"544", "334"},
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

func Test_validatePasswordOne(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "1111111",
			},
			want: true,
		},
		{
			args: args{
				s: "123456789",
			},
			want: false,
		},
		{
			args: args{
				s: "123345678",
			},
			want: true,
		},
		{
			args: args{
				s: "123454321",
			},
			want: false,
		},
		{
			args: args{
				s: "123445321",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePasswordOne(tt.args.s); got != tt.want {
				t.Errorf("validatePasswordOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatePasswordTwo(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "123345678",
			},
			want: true,
		},
		{
			args: args{
				s: "122234567",
			},
			want: false,
		},
		{
			args: args{
				s: "122210567",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePasswordTwo(tt.args.s); got != tt.want {
				t.Errorf("validatePasswordTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRange(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantErr bool
	}{
		{
			args: args{
				r: strings.NewReader("100-200"),
			},
			want: []int64{100, 200},
		},
		{
			args: args{
				r: strings.NewReader("abcd-200"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				r: strings.NewReader("100-tool"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			args: args{
				r: badreadseeker.New(strings.NewReader("100-900"), io.ErrShortBuffer, badreadseeker.Read),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseRange(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRange() got = %v, want %v", got, tt.want)
			}
		})
	}
}
