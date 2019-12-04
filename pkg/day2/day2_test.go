package day2

import (
	"io"
	"strconv"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	type args struct {
		r io.ReadSeeker
	}
	tests := []struct {
		name string
		args args
		want string
		pos  int
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
			want: "invalid opcode: 0",
		},
		{
			args: args{
				r: strings.NewReader("1,12,2,0,1,0,0,0,1,0,0,0,1,0,0,0"),
			},
			want: "unexpected end of program",
		},
		{
			args: args{
				r: strings.NewReader("1,12,abc"),
			},
			want: "strconv.Atoi: parsing \"abc\": invalid syntax",
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOne(tt.args.r); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	type args struct {
		r io.ReadSeeker
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				r: strings.NewReader("1,2,3,0,99"),
			},
			want: strconv.Itoa(100*12 + 2),
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.r); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
