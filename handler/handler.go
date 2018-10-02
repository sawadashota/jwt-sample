// handler package is deal with request
package handler

import "os"

var (
	idRsaPath string
)

const (
	IdRsaPathDefault = "./certs/id_rsa"
)

func init() {
	idRsaPath = getEnvString("ID_RSA_PATH", IdRsaPathDefault)
}

func getEnvString(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		v = defaultValue
	}

	return v
}
