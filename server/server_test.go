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
	db := persistence.SetupDB()
	defer db.Close()

	_, err := GetDataAPIServer(db,os.Getenv("URL_PAGE"))
	assert.Nil(t, err)
}

func Test_getLogo(t *testing.T) {
	_, err := getLogo()
	assert.Nil(t, err)
}
