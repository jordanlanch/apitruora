package utils

import (
	"errors"
	"os"
)

var (
	errEnvRabbit    = errors.New("not found RABBIT_PATH environment variable")
	errEnvDBType    = errors.New("not found DATABASE_TYPE environment variable")
	errEnvDBUrl     = errors.New("not found TRACKING_DATABASE_URL environment variable")
	errEnvSecret    = errors.New("not found SECRET_KEY environment variable")
	errEnvDBUrlTest = errors.New("not found TRACKING_DATABASE_URL_TEST environment variable")
)

func ValidateEnvVars() error {

	if len(os.Getenv("RABBIT_PATH")) == 0 {
		return errEnvRabbit
	}
	if len(os.Getenv("DATABASE_TYPE")) == 0 {
		return errEnvDBType
	}
	if len(os.Getenv("TRACKING_DATABASE_URL")) == 0 {
		return errEnvDBUrl
	}
	if len(os.Getenv("SECRET_KEY")) == 0 {
		return errEnvSecret
	}
	return nil
}

func ValidateTestEnvVars() error {
	if len(os.Getenv("TRACKING_DATABASE_URL_TEST")) == 0 {
		return errEnvDBUrlTest
	}
	return nil
}
