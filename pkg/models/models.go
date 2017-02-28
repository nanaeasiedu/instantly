package models

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ngenerio/instantly/pkg/config"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func Setup() error {
	gooseDriver := goose.DBDriver{
		Name:    config.Settings.DBName,
		OpenStr: config.Settings.DBPath,
	}

	if config.Settings.DBName == "mysql" {
		gooseDriver.Dialect = &goose.MySqlDialect{}
		gooseDriver.Import = "github.com/go-sql-driver/mysql"
	} else {
		gooseDriver.Dialect = &goose.PostgresDialect{}
		gooseDriver.Import = "github.com/lib/pq"
	}

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
