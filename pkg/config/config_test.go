package config

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	t.Parallel()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "should success to load data from json file",
			args: args{
				path: filepath.Join(dir, "testdata", "success.json"),
			},
			want: &Config{
				Server: ServerConfig{
					Port:      8080,
					EnableTLS: false,
				},
			},
			wantErr: false,
		},
		{
			name: "should success to load data from yaml file",
			args: args{
				path: filepath.Join(dir, "testdata", "success.yaml"),
			},
			want: &Config{
				Server: ServerConfig{
					Port:      8080,
					EnableTLS: false,
				},
			},
			wantErr: false,
		},
		{
			name: "should success to load data from yml file",
			args: args{
				path: filepath.Join(dir, "testdata", "success.yml"),
			},
			want: &Config{
				Server: ServerConfig{
					Port:      8080,
					EnableTLS: false,
				},
			},
			wantErr: false,
		},
		{
			name: "should fail to load data because unsupported file type is specified",
			args: args{
				path: filepath.Join(dir, "testdata", "notfound.txt"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(sub *testing.T) {
			sub.Parallel()

			actual, err := LoadFromFile(tt.args.path)
			if tt.wantErr && err == nil {
				sub.Errorf("LoadFromFile should return error but got nil.")
				return
			}
			if !tt.wantErr && err != nil {
				sub.Errorf("LoadFromFile should not return error but got %v.", err)
				return
			}
			if !reflect.DeepEqual(actual, tt.want) {
				sub.Errorf("LoadFromFile should return %v but %v", tt.want, actual)
			}
		})
	}
}

func Test_loadJSON(t *testing.T) {
	t.Parallel()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "should success to load data from specified path",
			args: args{
				path: filepath.Join(dir, "testdata", "success.json"),
			},
			want: &Config{
				Server: ServerConfig{
					Port:      8080,
					EnableTLS: false,
				},
			},
			wantErr: false,
		},
		{
			name: "should fail to load data from specified path",
			args: args{
				path: filepath.Join(dir, "testdata", "notfound.json"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should fail to convert struct from specified file",
			args: args{
				path: filepath.Join(dir, "testdata", "invalid.json"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(sub *testing.T) {
			sub.Parallel()

			actual, err := loadJSON(tt.args.path)
			if tt.wantErr && err == nil {
				sub.Errorf("loadJSON should return error but got nil.")
				return
			}
			if !tt.wantErr && err != nil {
				sub.Errorf("loadJSON should not return error but got %v.", err)
				return
			}
			if !reflect.DeepEqual(actual, tt.want) {
				sub.Errorf("loadJSON should return %v but %v", tt.want, actual)
			}
		})
	}
}

func Test_loadYAML(t *testing.T) {
	t.Parallel()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "should success to load data from specified path",
			args: args{
				path: filepath.Join(dir, "testdata", "success.yaml"),
			},
			want: &Config{
				Server: ServerConfig{
					Port:      8080,
					EnableTLS: false,
				},
			},
			wantErr: false,
		},
		{
			name: "should fail to load data from specified path",
			args: args{
				path: filepath.Join(dir, "testdata", "notfound.yaml"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should fail to convert struct from specified file",
			args: args{
				path: filepath.Join(dir, "testdata", "invalid.yaml"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(sub *testing.T) {
			sub.Parallel()

			actual, err := loadYAML(tt.args.path)
			if tt.wantErr && err == nil {
				sub.Errorf("loadYAML should return error but got nil.")
				return
			}
			if !tt.wantErr && err != nil {
				sub.Errorf("loadYAML should not return error but got %v.", err)
				return
			}
			if !reflect.DeepEqual(actual, tt.want) {
				sub.Errorf("loadYAML should return %v but %v", tt.want, actual)
			}
		})
	}
}

func Test_loadBinary(t *testing.T) {
	t.Parallel()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "should success to load data from specified path",
			args: args{
				path: filepath.Join(dir, "testdata", "testdata.txt"),
			},
			want:    []byte("test\n"),
			wantErr: false,
		},
		{
			name: "should fail to load data from specified path",
			args: args{
				path: filepath.Join(dir, "testdata", "notfound.txt"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(sub *testing.T) {
			sub.Parallel()

			actual, err := loadBinary(tt.args.path)
			if tt.wantErr && err == nil {
				sub.Errorf("loadBinary should return error but got nil.")
				return
			}
			if !tt.wantErr && err != nil {
				sub.Errorf("loadBinary should not return error but got %v.", err)
				return
			}
			if !reflect.DeepEqual(actual, tt.want) {
				sub.Errorf("loadBinary should return %v but %v", tt.want, actual)
			}
		})
	}
}
