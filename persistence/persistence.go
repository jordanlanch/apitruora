package persistence

import (
	"fmt"
	"os"

	"../dbmodels"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //for use postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //for use sqlite at tests
)

//ConnectToDB func that create a DB connection
func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(os.Getenv("DB"), os.Getenv("URL_DB"))
	return db, err
}

//CleanDB data base
func CleanDB() {
	db, err := ConnectToDB()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	db.Exec("drop table responses;")
	db.Exec("drop table servers;")
	db.Exec("drop table items;")
	defer db.Close()
}

func SetupDB() *gorm.DB {
	db, err := ConnectToDB()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	// Migrate the schema
	db.AutoMigrate(&dbmodels.Response{}, &dbmodels.Servers{},&dbmodels.Items{})

	return db
}

func CreateItems(db  *gorm.DB,Items *dbmodels.Items) (*dbmodels.Items, error){
	if err := db.Create(&Items).Error; err != nil {
		return nil, err
	} 
	return Items, nil
}
