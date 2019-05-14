package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEnvVars(t *testing.T) {
	envVarsSetted := ValidateEnvVars()
	assert.Equal(t, envVarsSetted, nil, "No errors in enviroment variables")

	SECRET_KEY := os.Getenv("SECRET_KEY")
	os.Setenv("SECRET_KEY", "")
	assert.Equal(t, errEnvSecret, ValidateEnvVars())
	os.Setenv("SECRET_KEY", SECRET_KEY)

	RABBIT_PATH := os.Getenv("RABBIT_PATH")
	os.Setenv("RABBIT_PATH", "")
	assert.Equal(t, errEnvRabbit, ValidateEnvVars())
	os.Setenv("RABBIT_PATH", RABBIT_PATH)

	DATABASE_TYPE := os.Getenv("DATABASE_TYPE")
	os.Setenv("DATABASE_TYPE", "")
	assert.Equal(t, errEnvDBType, ValidateEnvVars())
	os.Setenv("DATABASE_TYPE", DATABASE_TYPE)

	TRACKING_DATABASE_URL := os.Getenv("TRACKING_DATABASE_URL")
	os.Setenv("TRACKING_DATABASE_URL", "")
	assert.Equal(t, errEnvDBUrl, ValidateEnvVars())
	os.Setenv("TRACKING_DATABASE_URL", TRACKING_DATABASE_URL)
}
