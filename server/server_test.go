package server

import (
	"testing"

	"../persistence"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	db := persistence.SetupDB_Test()
	defer db.Close()

	_, err := GetDataAPIServer(db,"truora.com")
	assert.Nil(t, err)
}

func Test_getLogo(t *testing.T) {
	_, err := getLogo()
	assert.Nil(t, err)
}
