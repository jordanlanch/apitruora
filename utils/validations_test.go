package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEnvVars(t *testing.T) {
	envVarsSetted := ValidateEnvVars()
	assert.Equal(t, envVarsSetted, nil, "No errors in enviroment variables")

	URL_APISERVER := os.Getenv("URL_APISERVER")
	os.Setenv("URL_APISERVER", "")
	assert.Equal(t, errEnvURLAPI, ValidateEnvVars())
	os.Setenv("URL_APISERVER", URL_APISERVER)

	RABBITURL_DB_PATH := os.Getenv("URL_DB")
	os.Setenv("URL_DB", "")
	assert.Equal(t, errEnvURLDB, ValidateEnvVars())
	os.Setenv("URL_DB", URL_DB)

	DB := os.Getenv("DB")
	os.Setenv("DB", "")
	assert.Equal(t, errEnvDB, ValidateEnvVars())
	os.Setenv("DB", DB)

	URL_PAGE := os.Getenv("URL_PAGE")
	os.Setenv("URL_PAGE", "")
	assert.Equal(t, errEnvURLPage, ValidateEnvVars())
	os.Setenv("URL_PAGE", URL_PAGE)
}
