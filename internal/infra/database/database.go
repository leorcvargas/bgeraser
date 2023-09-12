// Package database is responsible for all
// database related code
package database

import (
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leorcvargas/bgeraser/ent"
	"github.com/leorcvargas/bgeraser/internal/infra/config"
)

func NewEntClient(cfg *config.Config) *ent.Client {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	driver := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver))

	return client
}
