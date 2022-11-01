package migration

import (
	"fmt"
	"learning/golang-gormdatabase/database"
	"learning/golang-gormdatabase/models"
)

func RunMigration() {
	err := database.Db.AutoMigrate(
		&models.User{}, 
		&models.Locker{},
		&models.Post{},
		&models.Tag{},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Database Migrated")
}