// config package provides features to control the application configuration
package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/jesseduffield/yaml"
)

// Config holds the application configuration.
type Config struct {
	// Server is the server configuration.
	Server ServerConfig `yaml:"server" json:"server"`
}

// ServerConfig holds the server configuration.
type ServerConfig struct {
	// Port is the port number to listen on.
	Port int `yaml:"port" json:"port"`

	// EnableTLS indicates whether to enable TLS.
	EnableTLS bool `yaml:"enable_tls" json:"enable_tls"`

	// NOTE: SecretKeyPath and CertificatePath are the required options when EnableTLS is true.
	// SecretKeyPath is the path to the secret key file.
	SecretKeyPath string `yaml:"secret" json:"secret" validate:"tls_require"`
	// CertificatePath is the path to the certificate file.
	CertificatePath string `yaml:"certificate" json:"certificate" validate:"tls_require"`
}

type loader func(path string) (*Config, error)

var (
	ErrInvalidFileType = errors.New("given file type is not supported")
)

func defaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:            8080,
		EnableTLS:       false,
		SecretKeyPath:   "",
		CertificatePath: "",
	}
}

func defaultConfig() *Config {
	return &Config{
		Server: *defaultServerConfig(),
	}
}

// LoadFromFile loads the configuration from a specified file.
// This function will return an error if the file does not exist or is not a valid configuration file.
// This function suggests that the file type from extension.
func LoadFromFile(path string) (*Config, error) {
	var loader loader

	// suggest file type
	ext := filepath.Ext(path)
	switch ext {
	case ".yaml", ".yml":
		loader = loadYAML
	case ".json":
		loader = loadJSON
	default:
		return nil, ErrInvalidFileType
	}

	config, err := loader(path)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// loadJSON loads the configuration from a specified JSON file.
func loadJSON(path string) (*Config, error) {
	b, err := loadBinary(path)
	if err != nil {
		return nil, err
	}

	config := *defaultConfig()
	if err := json.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// loadYAML loads the configuration from a specified YAML file.
func loadYAML(path string) (*Config, error) {
	b, err := loadBinary(path)
	if err != nil {
		return nil, err
	}

	config := *defaultConfig()
	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// loadBinary loads the binary data from a specified file.
func loadBinary(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// ValidateConfig validates the configuration.
func ValidateConfig(cnf *Config) error {
	if err := ValidateServerConfig(&cnf.Server); err != nil {
		return err
	}
	return nil
}

// ValidateServerConfig validates the server configuration.
func ValidateServerConfig(cnf *ServerConfig) error {
	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterValidation("tls_require", TLSRequire)

	if err := v.Struct(cnf); err != nil {
		return err
	}
	return nil
}
