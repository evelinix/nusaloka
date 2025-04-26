package config

import (
	"fmt"
	"time"

	"github.com/evelinix/nusaloka/internal/account/model"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PostgresDB *gorm.DB

func InitDatabase() {

	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		AccountConfig.DBHost, AccountConfig.DBUser, AccountConfig.DBPass, AccountConfig.DBName, AccountConfig.DBPort)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("[db] Failed to connect to Postgres")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("[db] Failed to get DB instance")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		log.Fatal().
			Err(err).
			Msg("[db] Failed to ping DB")
	}

	PostgresDB = db
	log.Info().Msg("[db] PostgreSQL connected successfully")
}

func AutoMigrateDatabase() {
	if err := PostgresDB.AutoMigrate(
		&model.User{},
		&model.Referal{},
		&model.Webauth{},
	); err != nil {
		log.Fatal().
			Err(err).
			Msg("[db] Failed to migrate tables")
	}
	log.Info().Msg("[db] Migrated tables successfully")
}
