package rizz

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vippsas/summerstudents-backend/internal/config"
	"github.com/vippsas/summerstudents-backend/internal/database"
)

func SetupTestRepo() (*Repository, func() error, error) {
	cfg := config.MSSQLDBConfig{
		SQLHost:        "sql-rafqkaoj4kdlc.database.windows.net",
		SQLDatabase:    "rizz-and-comp-database",
		SQLPassword:    "SuperSecret1337",
		SQLUsername:    "sa",
		LocalSQLHost:   "localhost",
		LocalSQLPort:   "1436",
		LocalSQLDbName: "summerstudents-db",
	}

	localDb, err := database.LocalDb(&cfg)

	if err != nil {
		return nil, nil, err
	}

	db := sqlx.NewDb(localDb, "sqlserver")
	repo := NewRepository(db)

	cleanup := func() error {
		tables := []string{"RIZZ.presetAmounts", "RIZZ.recurringAgreements", "RIZZ.Merchants"}

		// Execute DELETE statements for each table.
		for _, table := range tables {
			stmt := fmt.Sprintf("DELETE FROM %s", table)
			_, err := db.Exec(stmt)
			if err != nil {
				return err
			}
		}

		return nil
	}

	_ = cleanup()

	return repo, cleanup, nil
}
