package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )

var Db *gorm.DB
func InitDb() *gorm.DB { // OOP constructor
	Db = connectDB()
	fmt.Println("Tersambung ke database")
	return Db
}

func connectDB() (*gorm.DB) {
	dsn := "host=localhost user=yoga password=1234 dbname=DBgorm port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err !=nil {
		fmt.Println("Error...")
		return nil
	}
	return db
}