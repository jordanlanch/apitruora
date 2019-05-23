package server

import (
	"testing"
	"os"

	"../persistence"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	os.Setenv("URL_DB", os.Getenv("URL_DB_TEST"))
	os.Setenv("DB", os.Getenv("DB_TEST"))
	persistence.CleanDB()
	db := persistence.SetupDB()
	defer db.Close()

	_, err := GetDataAPIServer(db,os.Getenv("URL_PAGE"))
	assert.Nil(t, err)
}

func Test_getLogoAndTitle(t *testing.T) {
	_,_, err := getLogoAndTitle("truora.com")
	assert.Nil(t, err)
}
