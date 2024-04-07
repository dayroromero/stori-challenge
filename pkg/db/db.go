package db

import (
	"log"

	"github.com/dayroromero/stori-challenge/pkg/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`) //Implements uuid_generate_v4()
	db.AutoMigrate(
		&models.Account{},
		&models.Transaction{},
		&models.User{},
	)

	return Handler{db}
}

func (h *Handler) Close() error {
	db, err := h.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}