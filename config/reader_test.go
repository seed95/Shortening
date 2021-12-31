package config

import (
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		path string
		cfg  *Config
	}

	tests := []struct {
		name     string
		args     args
		wantsErr bool
	}{
		{
			name: "parse config.yaml",
			args: args{
				path: "testdata/config.yaml",
				cfg:  &Config{},
			},
			wantsErr: false,
		},
		{
			name: "parse config.yml",
			args: args{
				path: "testdata/config.yml",
				cfg:  &Config{},
			},
			wantsErr: false,
		},
		{
			name: "parse config.json",
			args: args{
				path: "testdata/config.json",
				cfg:  &Config{},
			},
			wantsErr: true,
		},
		{
			name: "empty config file",
			args: args{
				path: "",
				cfg:  &Config{},
			},
			wantsErr: true,
		},
		{
			name: "parse a file that does not exist",
			args: args{
				path: "config.yml",
				cfg:  &Config{},
			},
			wantsErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Parse(tt.args.path, tt.args.cfg); (err != nil) != tt.wantsErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantsErr)
			}
		})
	}
}

func TestParseYaml(t *testing.T) {
	type args struct {
		path string
		cfg  *Config
	}

	tests := []struct {
		name     string
		args     args
		wantsErr bool
	}{
		{
			name: "parse config.yaml",
			args: args{
				path: "testdata/config.yaml",
				cfg:  &Config{},
			},
			wantsErr: false,
		},
		{
			name: "parse config.yml",
			args: args{
				path: "testdata/config.yml",
				cfg:  &Config{},
			},
			wantsErr: false,
		},
		{
			name: "parse config.json",
			args: args{
				path: "testdata/config.json",
				cfg:  &Config{},
			},
			wantsErr: true,
		},
		{
			name: "empty config file",
			args: args{
				path: "",
				cfg:  &Config{},
			},
			wantsErr: true,
		},
		{
			name: "parse a file that does not exist",
			args: args{
				path: "config.yml",
				cfg:  &Config{},
			},
			wantsErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseYaml(tt.args.path, tt.args.cfg); (err != nil) != tt.wantsErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantsErr)
			}
		})
	}
}
