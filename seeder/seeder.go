package seeder

import (
	"fmt"
	"log"
	"reflect"

	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func Execute(db *gorm.DB, seedMethodNames []string) {
	s := Seeder{
		db: db,
	}

	seederType := reflect.TypeOf(s)

	if (len(seedMethodNames) == 0 ) {
		fmt.Println("Running all seed methods...")

		for i := 0; i < seederType.NumMethod(); i++ {
			m := seederType.Method(i)
			seed(s, m.Name)
		}
	}


	for _, seedMethodName := range seedMethodNames {
		seed(s, seedMethodName);
	}
} 

func seed(s Seeder, seedMedthodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMedthodName)

	if !m.IsValid() {
		log.Fatalf("Method %s does not exist on type Seeder\n", seedMedthodName)
	}

	fmt.Println("Seeding ", seedMedthodName, "...")
	m.Call(nil)
	fmt.Println("Seed ", seedMedthodName, "succeed")
}