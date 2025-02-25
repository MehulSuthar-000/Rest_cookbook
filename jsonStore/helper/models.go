package helper

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "gopher"
	password = "gopher"
	dbname   = "mydb"
)

type Shipment struct {
	gorm.Model
	Packages []Package
	Data     string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Package struct {
	gorm.Model
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

// GORM creates tables with plural names
// Use this to suppress it

func (Shipment) TableName() string {
	return "Shipment"
}

func (Package) TableName() string {
	return "Package"
}

func InitDB() (db *gorm.DB, err error) {
	var connectionString = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)

	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	/*
		// The below AutoMigrate is equivalent to this
		if !db.HasTable("Shipment") {
			db.CreateTable(&Shipment{})
		}

		if !db.HasTable("Package") {
			db.CreateTable(&Package{})
		}
	*/
	db.AutoMigrate(&Shipment{}, &Package{})
	return
}
