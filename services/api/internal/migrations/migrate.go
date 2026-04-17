package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/mmtaee/ocserv-dashboard/common/pkg/database"
	"github.com/mmtaee/ocserv-dashboard/common/pkg/logger"
)

var Migrations = []*gormigrate.Migration{
	Migration001,
	Migration002,
}

func Migrate() {
	logger.Info("Starting database migrations...")
	db := database.GetPostgres()

	m := gormigrate.New(db, gormigrate.DefaultOptions, Migrations)
	if err := m.Migrate(); err != nil {
		logger.Fatal("Failed to run migrations: %v", err)
	}

	logger.Info("Database migrations complete")
}
