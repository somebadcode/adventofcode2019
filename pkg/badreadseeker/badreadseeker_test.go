package badreadseeker

import (
	"crypto/sha256"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestBadStringReader_Read(t *testing.T) {
	type fields struct {
		readSeeker io.ReadSeeker
		when       when
		error      error
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
				readSeeker: strings.NewReader("garbage"),
				when:       Read,
				error:      io.ErrNoProgress,
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
			r := &BadReadSeeker{
				readSeeker: tt.fields.readSeeker,
				when:       tt.fields.when,
				error:      tt.fields.error,
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

func TestNewBadReadSeeker(t *testing.T) {
	type args struct {
		readSeeker io.ReadSeeker
		e          error
	}
	tests := []struct {
		name string
		args args
		want *BadReadSeeker
	}{
		{
			args: args{
				readSeeker: strings.NewReader("garbage"),
				e:          io.ErrUnexpectedEOF,
			},
			want: &BadReadSeeker{
				readSeeker: strings.NewReader("garbage"),
				error:      io.ErrUnexpectedEOF,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.readSeeker, tt.args.e, 0)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			h1, h2 := sha256.New224(), sha256.New224()
			_, err := io.Copy(h1, tt.want.readSeeker)
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(h2, got)
			if err != nil {
				panic(err)
			}
		})
	}
}

func TestBadReadSeeker_Seek(t *testing.T) {
	type fields struct {
		readSeeker io.ReadSeeker
		when       when
		error      error
	}
	type args struct {
		offset int64
		whence int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{
				readSeeker: strings.NewReader("garbage"),
				when:       Seek,
				error:      io.ErrClosedPipe,
			},
			args: args{
				offset: 0,
				whence: io.SeekEnd,
			},
			wantErr: true,
		},
		{
			fields: fields{
				readSeeker: strings.NewReader("garbage"),
			},
			args: args{
				offset: 0,
				whence: io.SeekEnd,
			},
			wantErr: false,
		},
		{
			fields: fields{
				readSeeker: New(strings.NewReader("garbage"), io.ErrClosedPipe, Seek),
			},
			args: args{
				offset: 0,
				whence: io.SeekEnd,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BadReadSeeker{
				readSeeker: tt.fields.readSeeker,
				when:       tt.fields.when,
				error:      tt.fields.error,
			}
			_, err := r.Seek(tt.args.offset, tt.args.whence)
			if (err != nil) != tt.wantErr {
				t.Errorf("Seek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
