package persistence

import (
	"fmt"

	"../dbmodels"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //for use postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //for use sqlite at tests
)

const addr = "postgresql://root@localhost:26257/apitruora?sslmode=disable"


func SetupDB() *gorm.DB {
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Migrate the schema
	db.AutoMigrate(&dbmodels.Response{}, &dbmodels.Servers{})

	return db
}

func CreateResponse(db  *gorm.DB,response *dbmodels.Response) (*dbmodels.Response, error){
	if err := db.Create(&response).Error; err != nil {
		return nil, err
	} 
		return response, nil
	
}
