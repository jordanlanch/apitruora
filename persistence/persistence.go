package persistence

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //for use postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //for use sqlite at tests
)

var (
	ErrDBConnection = errors.New("failed to connect database")
)

//ConnectToDB func that create a DB connection
func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(os.Getenv("DATABASE_TYPE"), os.Getenv("DATABASE_URL"))
	return db, err
}

//CleanDB data base
func CleanDB() {
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal(ErrDBConnection)
	}
	db.Exec("drop table server;")
	defer db.Close()
}

//MigrateDB database
func MigrateDB() (*gorm.DB, error) {
	db, err := ConnectToDB()

	if err != nil {
		utils.Error(err)
		panic(ErrDBConnection)
	}

	db.AutoMigrate(Server{})
	return db, err
}

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"-"`
}

// Server db struct
type Server struct {
	Model
	VIN        string `gorm:"type:varchar(30);unique_index" json:"vin"`
	Plate      string `gorm:"not null" json:"plate"`
	Status     string `gorm:"not null" json:"status"`
	WebhookURL string `gorm:"not null" json:"webhook_url"`
	Active     bool   `gorm:"not null;default:true" json:"active"`
}
