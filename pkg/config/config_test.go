package config

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func Test_defaultServerConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want *ServerConfig
	}{
		{
			name: "should success to create default server config",
			want: &ServerConfig{
				Port:      8080,
				EnableTLS: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(sub *testing.T) {
			sub.Parallel()

			if got := defaultServerConfig(); !reflect.DeepEqual(got, tt.want) {
				sub.Errorf("defaultServerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "should success to create default config",
			want: &Config{
				Server: *defaultServerConfig(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(sub *testing.T) {
			sub.Parallel()

			if got := defaultConfig(); !reflect.DeepEqual(got, tt.want) {
				sub.Errorf("defaultConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func Test_ValidateConfig(t *testing.T) {
	t.Parallel()

	type args struct {
		config *Config
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should success to validate config",
			args: args{
				config: &Config{
					Server: ServerConfig{
						Port:      8080,
						EnableTLS: false,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "should fail to validate config because server config is invalid",
			args: args{
				config: &Config{
					Server: ServerConfig{
						Port:      8080,
						EnableTLS: true,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(sub *testing.T) {
			sub.Parallel()

			if err := ValidateConfig(tt.args.config); (err != nil) != tt.wantErr {
				sub.Errorf("ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ValidateServerConfig(t *testing.T) {
	t.Parallel()

	type args struct {
		config *ServerConfig
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should success to validate config",
			args: args{
				config: &ServerConfig{
					Port:      8080,
					EnableTLS: false,
				},
			},
			wantErr: false,
		},
		{
			name: "should fail to validate config because secret key path is empty",
			args: args{
				config: &ServerConfig{
					Port:            8080,
					EnableTLS:       true,
					SecretKeyPath:   "",
					CertificatePath: "test",
				},
			},
			wantErr: true,
		},
		{
			name: "should fail to validate config because certificate path is empty",
			args: args{
				config: &ServerConfig{
					Port:            8080,
					EnableTLS:       true,
					SecretKeyPath:   "test",
					CertificatePath: "",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(sub *testing.T) {
			sub.Parallel()

			if err := ValidateServerConfig(tt.args.config); (err != nil) != tt.wantErr {
				sub.Errorf("ValidateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
