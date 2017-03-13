package models

import (
	"fmt"

	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ngenerio/instantly/pkg/config"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func Setup() error {
	log.Info(fmt.Sprintf("Coonfiguring db: %s %s %s", config.Settings.DBName, config.Settings.DBPath, config.Settings.Env))
	gooseDriver := goose.DBDriver{
		Name:    config.Settings.DBName,
		OpenStr: config.Settings.DBPath,
	}

	gooseDriver.Dialect = &goose.PostgresDialect{}
	gooseDriver.Import = "github.com/lib/pq"

	migrateConf := &goose.DBConf{
		MigrationsDir: config.Settings.MigrationsDir,
		Env:           config.Settings.Env,
		Driver:        gooseDriver,
	}

	latest, err := goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)
	if err != nil {
		log.Error(err)
		return err
	}

	db, err = gorm.Open(config.Settings.DBName, config.Settings.DBPath)

	if err != nil {
		log.Error(err)
		return err
	}

	err = goose.RunMigrationsOnDb(migrateConf, migrateConf.MigrationsDir, latest, db.DB())
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
