package intcodemachine

import (
	"bytes"
	"fmt"
	"github.com/somebadcode/adventofcode2019/pkg/badreadseeker"
	"io"
	"math"
	"strings"
	"testing"
)

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
				r: 'x',
			},
			want: false,
		},
		{
			args: args{
				r: 0x2029,
			},
			want: true,
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

func Test_scanInstructions(t *testing.T) {
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
				data: []byte{0x77},
			},
			wantAdvance: 0,
			wantToken:   []byte{},
			wantErr:     false,
		},
		{
			args: args{
				data: []byte("5,4,3,2,1,0"),
			},
			wantAdvance: 2,
			wantToken:   []byte{'5'},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdvance, gotToken, err := scanInstructions(tt.args.data, tt.args.atEOF)
			if (err != nil) != tt.wantErr {
				t.Errorf("scanInstructions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotAdvance != tt.wantAdvance {
				t.Errorf("scanInstructions() gotAdvance = %v, want %v", gotAdvance, tt.wantAdvance)
			}
			if !bytes.Equal(gotToken, tt.wantToken) {
				t.Errorf("scanInstructions() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func TestMachine_LoadProgram(t *testing.T) {
	type fields struct {
		memory []int
		tape   io.ReadSeeker
		ip     int
		err    error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			fields: fields{
				tape: strings.NewReader("1,1,2,3,99"),
			},
			wantErr: false,
		},
		{
			fields: fields{
				tape: strings.NewReader(fmt.Sprintf("1,1%d,2,3,99", math.MaxInt64)),
			},
			wantErr: true,
		},
		{
			fields: fields{
				tape: strings.NewReader("once upon a time\n"),
			},
			wantErr: true,
		},
		{
			fields: fields{
				tape: badreadseeker.New(strings.NewReader("0,1,2,3,99"), io.ErrShortBuffer, badreadseeker.Read),
			},
			wantErr: true,
		},
		{
			fields: fields{
				tape: bytes.NewReader([]byte{}),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Machine{
				memory: tt.fields.memory,
				tape:   tt.fields.tape,
				ip:     tt.fields.ip,
				err:    tt.fields.err,
			}
			if err := m.LoadProgram(); (err != nil) != tt.wantErr {
				t.Errorf("LoadProgram() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
