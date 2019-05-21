package utils

import (
	"errors"
	"os"
)

var (
	errEnvURLAPI    = errors.New("not found URL_APISERVER environment variable")
	errEnvURLDB    = errors.New("not found URL_DB environment variable")
	errEnvDB     = errors.New("not found DB environment variable")
	errEnvURLPage    = errors.New("not found URL_PAGE environment variable")
	errEnvURLDBTest    = errors.New("not found URL_DB_TEST environment variable")
	errEnvDBTest    = errors.New("not found DB_TEST environment variable")
)

func ValidateEnvVars() error {

	if len(os.Getenv("URL_APISERVER")) == 0 {
		return errEnvURLAPI
	}
	if len(os.Getenv("URL_DB")) == 0 {
		return errEnvURLDB
	}
	if len(os.Getenv("DB")) == 0 {
		return errEnvDB
	}

	if len(os.Getenv("URL_PAGE")) == 0 {
		return errEnvURLPage
	}
	return nil
}

func ValidateTestEnvVars() error {
	if len(os.Getenv("URL_DB_TEST")) == 0 {
		return errEnvURLDBTest
	}
	if len(os.Getenv("DB_TEST")) == 0 {
		return errEnvDBTest
	}
	return nil
}
