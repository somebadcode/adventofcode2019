package intcodemachine

import (
	"github.com/somebadcode/adventofcode2019/pkg/badreadseeker"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestMachine_Output(t *testing.T) {
	type fields struct {
		memory []int
		tape   io.ReadSeeker
		ip     int
		err    error
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				memory: []int{1, 2, 3},
			},
			want: 1,
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
			if got := m.Output(); got != tt.want {
				t.Errorf("Output() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestMachine_Reset(t *testing.T) {
	type fields struct {
		memory []int
		tape   io.ReadSeeker
		ip     int
		err    error
	}
	tests := []struct {
		name       string
		fields     fields
		seek       int
		wantErr    bool
		wantIp     int
		wantOutput int
	}{
		{
			fields: fields{
				memory: []int{1,2,3,4,5,6,7,8,9,0},
				tape:   strings.NewReader("1,2,3,4,5,6,7,8,9,0"),
				ip:     5,
			},
			seek:       5,
			wantErr:    false,
			wantIp:     0,
			wantOutput: 1,
		},
		{
			fields: fields{
				memory: []int{1,2,3,0,1,0,0,0,99,99},
				tape:   badreadseeker.New(strings.NewReader("1,2,3,0,1,0,0,0,99,99"), io.ErrShortBuffer, badreadseeker.Read),
				ip:     5,
			},
			seek:    0,
			wantErr: true,
			wantIp:  0,
		},
		{
			fields: fields{
				memory: []int{1,2,3,4,5,6,7,8,9,0},
				tape:   badreadseeker.New(strings.NewReader("1,2,3,4,5,6,7,8,9,0"), io.ErrShortBuffer, badreadseeker.Seek),
				ip:     5,
			},
			seek:    0,
			wantErr: true,
			wantIp:  0,
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
			if err := m.Reset(); (err != nil) != tt.wantErr {
				t.Errorf("Reset() error = %v, wantErr %v", err, tt.wantErr)
				t.FailNow()
			}
			if got := m.ip; got != tt.wantIp {
				t.Errorf("Reset() ip = %v, wantIp %v", got, tt.wantIp)
				t.FailNow()
			}
			if got := m.Output(); got != tt.wantOutput {
				t.Errorf("Reset() + Output() = %v, wantOutput %v", got, tt.wantOutput)
			}
		})
	}
}
*/
func TestMachine_Run(t *testing.T) {
	type fields struct {
		memory []int
		tape   io.ReadSeeker
		ip     int
		err    error
	}
	tests := []struct {
		name       string
		fields     fields
		wantErr    bool
		wantOutput int
	}{
		{
			fields: fields{
				memory: []int{1, 1, 0, 0, 99},
			},
			wantErr:    false,
			wantOutput: 2,
		},
		{
			fields: fields{
				memory: []int{2, 1, 0, 0, 99},
			},
			wantErr:    false,
			wantOutput: 2,
		},
		{
			fields: fields{
				memory: []int{5, 1, 0, 0, 99},
			},
			wantErr:    true,
			wantOutput: 5,
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
			if err := m.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := m.Output(); got != tt.wantOutput {
				t.Errorf("Run() + Output() = %v, wantOutput %v", got, tt.wantOutput)
			}
		})
	}
}

func TestMachine_SetInput(t *testing.T) {
	type fields struct {
		memory []int
		tape   io.ReadSeeker
		ip     int
		err    error
	}
	type args struct {
		d0 int
		d1 int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantMemory []int
	}{
		{
			fields: fields{
				memory: []int{2, 0, 0, 0, 99},
			},
			args: args{
				d0: 4,
				d1: 6,
			},
			wantMemory: []int{2, 4, 6, 0, 99},
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
			m.SetInput(tt.args.d0, tt.args.d1)
			if !reflect.DeepEqual(tt.wantMemory, m.memory) {
				t.Errorf("SetInput() memory = %v, wantmemory %v", tt.fields.memory, tt.wantMemory)
			}
		})
	}
}

func TestMachine_add(t *testing.T) {
	type fields struct {
		memory []int
		tape   io.ReadSeeker
		ip     int
		err    error
	}
	tests := []struct {
		name       string
		fields     fields
		wantMemory []int
		wantErr    bool
	}{
		{
			fields: fields{
				memory: []int{1, 2, 3, 0, 99},
				ip:     0,
			},
			wantMemory: []int{3, 2, 3, 0, 99},
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
			m.add()
			if !reflect.DeepEqual(tt.wantMemory, m.memory) {
				t.Errorf("add() memory = %v, wantMemory %v", m.memory, tt.wantMemory)
			}
			if err := m.err; (err != nil) != tt.wantErr {
				t.Errorf("add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMachine_mul(t *testing.T) {
	type fields struct {
		memory []int
		tape   io.ReadSeeker
		ip     int
		err    error
	}
	tests := []struct {
		name       string
		fields     fields
		wantMemory []int
		wantErr    bool
	}{
		{
			fields: fields{
				memory: []int{2, 2, 3, 0, 99},
				ip:     0,
			},
			wantMemory: []int{0, 2, 3, 0, 99},
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
			m.mul()
			if !reflect.DeepEqual(tt.wantMemory, m.memory) {
				t.Errorf("mul() memory = %v, wantMemory %v", m.memory, tt.wantMemory)
			}
			if err := m.err; (err != nil) != tt.wantErr {
				t.Errorf("mul() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		tape io.ReadSeeker
	}
	tests := []struct {
		name    string
		args    args
		want    *Machine
		wantErr bool
	}{
		{
			args: args{
				tape: strings.NewReader("1,2,3"),
			},
			want: &Machine{
				memory: []int{1, 2, 3},
				ip:     0,
				err:    nil,
			},
		},
		{
			args: args{
				tape: badreadseeker.New(strings.NewReader("1,2,3"), io.ErrShortBuffer, badreadseeker.Read),
			},
			want: &Machine{
				memory: nil,
				ip:     0,
				err:    nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.tape)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(got.memory, tt.want.memory) {
					t.Errorf("New() memory = %v, wantErr %v", got.memory, tt.want.memory)
				}
			}
		})
	}
}

func TestMachine_Reset1(t *testing.T) {
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
				memory: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				tape:   strings.NewReader("1,2,3,4,5,6,7,8,9,0"),
				ip:     5,
			},
			wantErr: false,
		},
		{
			fields: fields{
				memory: []int{1, 2, 3, 0, 1, 0, 0, 0, 99, 99},
				tape:   badreadseeker.New(strings.NewReader("1,2,3,0,1,0,0,0,99,99"), io.ErrShortBuffer, badreadseeker.Read),
				ip:     5,
			},
			wantErr: true,
		},
		{
			fields: fields{
				memory: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				tape:   badreadseeker.New(strings.NewReader("1,2,3,4,5,6,7,8,9,0"), io.ErrShortBuffer, badreadseeker.Seek),
				ip:     5,
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
			if err := m.Reset(); (err != nil) != tt.wantErr {
				t.Errorf("Reset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
