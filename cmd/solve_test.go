package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := os.Chdir("./..")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func Test_solve(t *testing.T) {
	type args struct {
		path   string
		config *viper.Viper
		logger *log.Logger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				path:   "./testdata",
				config: mustGetConfig(),
				logger: log.New(os.Stdout, "", 0),
			},
			wantErr: false,
		},
		{
			args: args{
				path:   "./wrong",
				config: mustGetConfig(),
				logger: log.New(os.Stdout, "", 0),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := solve(tt.args.path, tt.args.config, tt.args.logger); (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
