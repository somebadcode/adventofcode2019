package vector

import (
	"fmt"
	"reflect"
	"testing"
	"unicode/utf8"
)

func Test_scanVectors(t *testing.T) {
	type args struct {
		data  []byte
		atEOF bool
	}
	tests := []struct {
		name        string
		args        args
		wantAdvance int
		wantToken   []byte
		wantErr     bool
	}{
		{
			args: args{
				data: []byte("R10,L40"),
			},
			wantAdvance: 4,
			wantToken:   []byte("R10"),
			wantErr:     false,
		},
		{
			args: args{
				data:  []byte("U50"),
				atEOF: true,
			},
			wantAdvance: 3,
			wantToken:   []byte("U50"),
			wantErr:     false,
		},
		{
			args: args{
				data:  []byte("U1"),
				atEOF: false,
			},
			wantAdvance: 0,
			wantToken:   nil,
			wantErr:     false,
		},
		{
			args: args{
				data:  []byte(fmt.Sprintf("%c", 0xFF0F000D)),
				atEOF: false,
			},
			wantAdvance: 0,
			wantToken:   []byte(string(utf8.RuneError)),
			wantErr:     true,
		},
		{
			args: args{
				data:  []byte(fmt.Sprintf("U2%c,L3", 0xD800)),
				atEOF: false,
			},
			wantAdvance: 0,
			wantToken:   []byte(string(utf8.RuneError)),
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdvance, gotToken, err := scanVectors(tt.args.data, tt.args.atEOF)
			if (err != nil) != tt.wantErr {
				fmt.Println(tt.wantToken)
				t.Errorf("scanVectors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAdvance != tt.wantAdvance {
				t.Errorf("scanVectors() gotAdvance = %v, want %v", gotAdvance, tt.wantAdvance)
			}
			if !reflect.DeepEqual(gotToken, tt.wantToken) {
				t.Errorf("scanVectors() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func Test_isLineEnding(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				r: '\n',
			},
			want: true,
		},
		{
			args: args{
				r: 'n',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLineEnding(tt.args.r); got != tt.want {
				t.Errorf("isLineEnding() = %v, want %v", got, tt.want)
			}
		})
	}
}
