package main

import (
	"database-seeder/seeder"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/bonvih/backend-app/accounts"
	"github.com/bonvih/backend-app/profiles"
)

func main() {
	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=postgres password=1981 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open db: %s", err)
	}

	if err = db.AutoMigrate(&accounts.Account{}); err != nil {
		log.Fatalf("Failed to migrate db: %s", err)
	}
	if err = db.AutoMigrate(&profiles.Profile{}); err != nil {
		log.Fatalf("Failed to migrate db: %s", err)
	}

	seeder.Execute(db, []string{})
	os.Exit(0)
}