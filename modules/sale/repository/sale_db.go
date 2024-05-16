package repository

import (
	"gorm.io/gorm"

	"testBackend/logs"
	"testBackend/modules/entities/models"
)

type saleRepositoryDB struct {
	db *gorm.DB
}

func NewsaleRepositoryDB(db *gorm.DB) models.SaleRepository {
	if err := db.AutoMigrate(&models.Sales{}); err != nil {
		logs.Fatal(err.Error())
	}
	return &saleRepositoryDB{db: db}
}

func (r *saleRepositoryDB) AddSale(sale *models.Sales) (*models.Sales, error) {
	if err := r.db.Create(sale).Error; err != nil {
		return nil, err
	}
	return sale, nil
}

func (r *saleRepositoryDB) GetSalesByEmployee(month int, year int, minSales int) ([]models.SalespersonResult, error) {
	var salespeople []models.SalespersonResult

	// คิวรีฐานข้อมูลด้วยคำสั่ง SQL
	if err := r.db.Table("sales").
		Select("salespersons.name, COUNT(*) as num_cars_sold").
		Joins("JOIN salespersons ON sales.sale_person_id = salespersons.sale_person_id").
		Where("EXTRACT(YEAR FROM to_timestamp(sales.year::text, 'YYYY')) = ? AND EXTRACT(MONTH FROM to_timestamp(sales.month::text, 'MM')) = ?", year, month).
		Group("salespersons.name").
		Having("COUNT(*) >= ?", minSales).
		Scan(&salespeople).Error; err != nil {
		return nil, err
	}

	return salespeople, nil
}

func (r *saleRepositoryDB) GetMonthlySalesSummary() ([]models.MonthlySalesSummary, error) {
	var monthlySales []models.MonthlySalesSummary

	// เตรียมคำสั่ง SQL
	query := `
		SELECT month, year, COUNT(serial_no) AS num_cars_sold, SUM(price) AS total_sales
		FROM sales
		GROUP BY month, year
		ORDER BY num_cars_sold DESC
	`

	// คิวรีฐานข้อมูลด้วยคำสั่ง SQL
	if err := r.db.Raw(query).Scan(&monthlySales).Error; err != nil {
		return nil, err
	}

	return monthlySales, nil
}
