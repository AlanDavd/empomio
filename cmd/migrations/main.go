// package migrations
package main

import (
	"log"
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/alandavd/empomio/cmd/migrations/config"
)

func main() {
	config := config.LoadConfigFlags()

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		config.PostgreSQLUser(),
		config.PostgreSQLPassword(),
		config.PostgreSQLName(),
		config.PostgreSQLHost(),
		config.PostgreSQLPort(),
		config.PostgreSQLSSLMode())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	switch config.Action() {
	case "migrate":
		if err := m.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	case "rollback":
		if err := m.RollbackLast(); err != nil {
			log.Fatalf("Rollback failed: %v", err)
		}
	}

	log.Println("Operations performed successfully")
}

var migrations = []*gormigrate.Migration{
	// {
	// 	ID: "",
	// 	Migrate: func(tx *gorm.DB) error {
	// 		return nil
	// 	},
	// 	Rollback: func(tx *gorm.DB) error {
	// 		return nil
	// 	},
	// },
}