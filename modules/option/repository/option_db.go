package repository

import (
	"gorm.io/gorm"

	"testBackend/logs"
	"testBackend/modules/entities/models"
)

type optionRepositoryDB struct {
	db *gorm.DB
}

func NewoptionRepositoryDB(db *gorm.DB) models.OptionRepository {
	if err := db.AutoMigrate(&models.Options{}); err != nil {
		logs.Fatal(err.Error())
	}
	return &optionRepositoryDB{db: db}
}

func (r *optionRepositoryDB) Add(option *models.Options) (*models.Options, error) {
	if err := r.db.Create(option).Error; err != nil {
		return nil, err
	}
	return option, nil
}