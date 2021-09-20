package seeder

import (
	"log"

	"github.com/bonvih/backend-app/accounts"
)

func (s Seeder) SeedAccounts() {
	if err := s.db.Create(&accounts.Account{
		Firstname: "john",
		Email: "john@gmail.com",
		Password: "1234",
		ProfileID: "0001",
	}).Error; err != nil {
		log.Fatalf("Could not create an account: %s", err)
	}
}