package models

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"github.com/ngenerio/instantly/pkg/config"
)

var db *gorm.DB

func Setup() error {
	migrateConf := &goose.DBConf{
		MigrationsDir: config.Settings.MigrationsDir,
		Env:           config.Settings.Env,
		Driver: goose.DBDriver{
			Name:    "postgres",
			OpenStr: config.Settings.PostgresURL,
			Dialect: &goose.PostgresDialect{},
			Import:  "github.com/lib/pq",
		},
	}

	latest, err := goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)
	if err != nil {
		log.Error(err)
		return err
	}

	db, err = gorm.Open("postgres", config.Settings.PostgresURL)
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
