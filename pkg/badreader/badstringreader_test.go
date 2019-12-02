package badreader

import (
	"io"
	"reflect"
	"testing"
)

func TestBadStringReader_Read(t *testing.T) {
	type fields struct {
		buffer []byte
		error  error
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{
			fields: fields{
				buffer: []byte("garbage"),
				error:  io.ErrNoProgress,
			},
			args: args{
				p: make([]byte, 4),
			},
			wantN:   4,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BadStringReader{
				buffer: tt.fields.buffer,
				error:  tt.fields.error,
			}
			gotN, err := r.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Read() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestNewBadStringReader(t *testing.T) {
	type args struct {
		s string
		e error
	}
	tests := []struct {
		name string
		args args
		want *BadStringReader
	}{
		{
			args: args{
				s: "garbage",
				e: io.ErrUnexpectedEOF,
			},
			want: &BadStringReader{
				buffer: []byte("garbage"),
				error:  io.ErrUnexpectedEOF,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBadStringReader(tt.args.s, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBadStringReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
