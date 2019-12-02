package day1

import (
	"github.com/somebadcode/adventofcode2019/pkg/badreader"
	"io"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				r: strings.NewReader("garbage"),
			},
			want: "strconv.ParseInt: parsing \"garbage\": invalid syntax",
		},
		{
			args: args{
				r: badreader.NewBadStringReader("666", io.ErrUnexpectedEOF),
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
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				r: strings.NewReader("garbage"),
			},
			want: "strconv.ParseInt: parsing \"garbage\": invalid syntax",
		},
		{
			args: args{
				r: badreader.NewBadStringReader("666", io.ErrUnexpectedEOF),
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
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.r); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
