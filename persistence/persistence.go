package persistence

import (
	"fmt"
	"os"

	"../dbmodels"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //for use postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //for use sqlite at tests
)


func SetupDB() *gorm.DB {
	db, err := gorm.Open(os.Getenv("DB"), os.Getenv("URL_DB"))
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
