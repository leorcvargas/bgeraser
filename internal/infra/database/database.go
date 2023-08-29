// Package database is responsible for all
// database related code
package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/leorcvargas/bgeraser/internal/infra/config"
	"github.com/leorcvargas/bgeraser/internal/infra/database/imagesdb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func NewPostgresDatabase(config *config.Config) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.Database.Host,
			config.Database.User,
			config.Database.Password,
			config.Database.Name,
			config.Database.Port,
		)

		gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln("Unable to connect to database:", err)
		}

		// TODO: use versioned migrations
		if err = gormDB.AutoMigrate(&imagesdb.Model{}); err != nil {
			log.Fatalln("Unable to migrate database:", err)
		}

		db = gormDB
	})

	return db
}
