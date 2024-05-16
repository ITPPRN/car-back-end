package repository

import (
	"gorm.io/gorm"

	"testBackend/logs"
	"testBackend/modules/entities/models"
)

type salespersonRepositoryDB struct {
	db *gorm.DB
}

func NewsalespersonRepositoryDB(db *gorm.DB) models.SalespersonRepository {
	if err := db.AutoMigrate(&models.Salesperson{}); err != nil {
		logs.Fatal(err.Error())
	}
	return &salespersonRepositoryDB{db: db}
}


func (r *salespersonRepositoryDB) Add(person *models.Salesperson) (*models.Salesperson, error) {
	if err := r.db.Create(person).Error; err != nil {
		return nil, err
	}
	return person, nil
}