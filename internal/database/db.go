package database

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/microsoft/go-mssqldb/azuread"
	"github.com/sirupsen/logrus"
	"github.com/vippsas/summerstudents-backend/internal/config"
	"net/url"
	"time"
)

var localEnv = false
var log *logrus.Entry

func NewDb(cfg *config.Config) (db *sqlx.DB, err error) {
	var sqldb *sql.DB

	if cfg.IsLocal() {
		fmt.Println("Connecting to local")
		sqldb, err = LocalDb(&cfg.DBConfig)
	} else {
		fmt.Println("Connecting to azure")
		sqldb, err = dbAuthAD(&cfg.DBConfig)
	}

	db = sqlx.NewDb(sqldb, "sqlserver")

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}

func LocalDb(cfg *config.MSSQLDBConfig) (*sql.DB, error) {
	//username := "sa"
	//password := "SuperSecret1337"
	//host := "localhost"
	//port := "1436"
	//dbName := "summerstudents-db"

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s",
		cfg.LocalSQLHost, cfg.SQLUsername, cfg.SQLPassword, cfg.LocalSQLPort, cfg.LocalSQLDbName)

	db, err := sql.Open("sqlserver", connectionString)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}

func dbAuthAD(cfg *config.MSSQLDBConfig) (db *sql.DB, err error) {
	dsn, err := azureSqlAdDsn(cfg, azureSqlDsnOptions{readOnly: false, encrypt: true})
	if err != nil {
		return nil, err
	}

	db, err = sql.Open(azuread.DriverName, dsn)
	if err != nil {
		return nil, err
	}

	// Same values as the old vennebetaling lib function
	db.SetMaxOpenConns(cfg.SQLMaxOpenConnections)
	db.SetMaxIdleConns(cfg.SQLMaxIdleConnections)
	db.SetConnMaxLifetime(300 * time.Second)

	return db, nil
}

type azureSqlDsnOptions struct {
	readOnly bool
	encrypt  bool
}

func azureSqlAdDsn(cfg *config.MSSQLDBConfig, options azureSqlDsnOptions) (string, error) {
	dsn, err := url.Parse(fmt.Sprint("sqlserver://", cfg.SQLHost))
	if err != nil {
		return "", fmt.Errorf("failed to parse base DSN: %v", err)
	}

	query := dsn.Query()

	// The driver recommends this too: https://github.com/microsoft/go-mssqldb#common-parameters
	if query.Get("connection timeout") != "" {
		return "", fmt.Errorf("don't use the 'connection timeout' query parameter, use context.Context instead")
	}

	if options.encrypt {
		query.Set("encrypt", "true")
		query.Set("TrustServerCertificate", "false")
		query.Set("hostNameInCertificate", "*.database.windows.net")
	}

	if options.readOnly {
		query.Set("ApplicationIntent", "ReadOnly")
	} else {
		query.Del("ApplicationIntent")
	}

	query.Set("database", cfg.SQLDatabase)
	query.Set("fedauth", azuread.ActiveDirectoryMSI)

	dsn.RawQuery = query.Encode()
	return dsn.String(), nil
}
