package config

import "github.com/go-playground/validator/v10"

// TLSRequire validates the secret key and certificate path when TLS is enabled.
func TLSRequire(ls validator.FieldLevel) bool {
	serverCnf := ls.Parent().Interface().(ServerConfig)
	if serverCnf.EnableTLS {
		return serverCnf.SecretKeyPath != "" && serverCnf.CertificatePath != ""
	}
	return true
}
