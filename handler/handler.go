// handler package is deal with request
package handler

import "os"

var (
	idRsaPath       string
	idRsaPublicPath string
)

const (
	IdRsaPathDefault       = "./certs/id_rsa"
	IdRsaPathPublicDefault = "./certs/id_rsa.pub.pkcs8"
)

func init() {
	idRsaPath = getEnvString("ID_RSA_PATH", IdRsaPathDefault)
	idRsaPublicPath = getEnvString("ID_RSA_PUBLIC_PATH", IdRsaPathPublicDefault)
}

func getEnvString(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		v = defaultValue
	}

	return v
}
