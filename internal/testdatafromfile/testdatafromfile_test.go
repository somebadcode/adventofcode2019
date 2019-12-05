package testdatafromfile

import (
	"bytes"
	"testing"
)

func TestDataFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		panic bool
	}{
		{
			name: "data from file",
			args: args{
				filename: "testdata.txt",
			},
			want: []byte("DO NOT TOUCH THIS FILE!"),
		},
		{
			name: "panic",
			args: args{
				filename: "panic.fail",
			},
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if err := recover(); err == nil {
						t.Errorf("From() did not panic, expected a panic")
					}
				}()
			}
			got := From(tt.args.filename)
			buffer := make([]byte, len(tt.want))
			n, err := got.Read(buffer)
			if err != nil || n != len(tt.want) {
				t.Errorf("From() + Read() = %v, want %v", buffer, tt.want)
			} else if i := bytes.Compare(buffer, tt.want); i != 0 {
				t.Errorf("From() + Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
