package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/vippsas/summerstudents-backend/internal/config"
	"github.com/vippsas/summerstudents-backend/internal/database"
	"github.com/vippsas/summerstudents-backend/internal/database/rizz"
	"os"
	"strings"
	"time"
)

const (
	contextTimeout = 10 * time.Second
)

type Api struct {
	Logger     *logrus.Entry
	Config     config.Config
	Repository rizz.Repository
}

func isLocal() bool {
	fmt.Println(os.Getenv("ENV"))
	return strings.ToLower(os.Getenv("ENV")) == "local"
}

func New() Api {

	conCfg := config.MSSQLConnectionConfig{
		SQLDialTimeout:               2,
		SQLConnectionTimeout:         2,
		SQLMaxOpenConnections:        50,
		SQLMaxIdleConnections:        15,
		SQLMaxConnectionLifetimeSecs: 300,
	}

	cfg := config.MSSQLDBConfig{
		SQLHost:               "sql-rafqkaoj4kdlc.database.windows.net",
		SQLDatabase:           "rizz-and-comp-db",
		SQLPassword:           "SuperSecret1337",
		SQLUsername:           "sa",
		LocalSQLHost:          "localhost",
		LocalSQLPort:          "1436",
		LocalSQLDbName:        "summerstudents-db",
		MSSQLConnectionConfig: conCfg,
	}

	conf := config.Config{
		DBConfig: cfg,
		IsLocal:  isLocal,
	}

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})

	logEntry := log.WithFields(logrus.Fields{
		"app": "summerstudents-backend",
		"env": "test",
	})

	db, err := database.NewDb(&conf)
	repo := rizz.NewRepository(db)

	if err != nil {
		log.WithError(err)
	}
	return Api{
		Logger:     logEntry,
		Config:     conf,
		Repository: *repo,
	}
}
