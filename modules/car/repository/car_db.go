package repository

import (
	"log"
	"strings"

	"gorm.io/gorm"

	"testBackend/logs"
	"testBackend/modules/entities/models"
)

type carRepositoryDB struct {
	db *gorm.DB
}

func NewoptionRepositoryDB(db *gorm.DB) models.CarRepository {
	if err := db.AutoMigrate(&models.Cars{}); err != nil {
		logs.Fatal(err.Error())
	}
	// Create a carRepositoryDB instance
    carRepo := carRepositoryDB{db: db}
    
    // Call CreateView to ensure the view exists
    if err := carRepo.CreateView(); err != nil {
        logs.Debug(err.Error())
    }
    if err := carRepo.CreateExpensiveCarView(); err != nil {
        logs.Debug(err.Error())
    }
	if err := carRepo.CreateLuxuriousCarView(); err != nil {
        logs.Debug(err.Error())
    }
	return &carRepositoryDB{db: db}
}

func (r *carRepositoryDB) CreateView() error {
	
	
	var viewExists bool
    err := r.db.Raw("SELECT EXISTS(SELECT 1 FROM information_schema.views WHERE table_name = 'economiccar')").Scan(&viewExists).Error
    if err != nil {
        logs.Warn("Failed to check for economiccar view:")
		return err
    }
    if !viewExists {
        // Create the EconomicCar view if it doesn't exist
        err = r.db.Exec("CREATE VIEW economiccar AS SELECT * FROM cars WHERE price <= 1000000 WITH CASCADED CHECK OPTION").Error
        if err != nil {
			return err
        }
        log.Println("EconomicCar view created successfully.")
    } else {
        log.Println("EconomicCar view already exists.")
    }
	return nil
}

func (r *carRepositoryDB) CreateExpensiveCarView() error {
    var viewExists bool
    err := r.db.Raw("SELECT EXISTS(SELECT 1 FROM information_schema.views WHERE table_name = 'expensivecar')").Scan(&viewExists).Error
    if err != nil {
        log.Fatal("Failed to check for expensivecar view:", err)
        return err
    }
    if !viewExists {
        // Create the ExpensiveCar view if it doesn't exist
        err = r.db.Exec("CREATE VIEW expensivecar AS SELECT * FROM cars WHERE price > 1000000 WITH CASCADED CHECK OPTION").Error
        if err != nil {
            log.Fatal("Failed to create ExpensiveCar view:", err)
            return err
        }
        log.Println("ExpensiveCar view created successfully.")
    } else {
        log.Println("ExpensiveCar view already exists.")
    }
    return nil
}

func (r *carRepositoryDB) CreateLuxuriousCarView() error {
    var viewExists bool
    err := r.db.Raw("SELECT EXISTS(SELECT 1 FROM information_schema.views WHERE table_name = 'luxuriouscar')").Scan(&viewExists).Error
    if err != nil {
        log.Fatal("Failed to check for luxuriouscar view:", err)
        return err
    }
    if !viewExists {
        // Create the LuxuriousCar view if it doesn't exist
        err = r.db.Exec("CREATE VIEW luxuriouscar AS SELECT * FROM ExpensiveCar WHERE price > 3000000 WITH CASCADED CHECK OPTION").Error
        if err != nil {
            log.Fatal("Failed to create luxuriouscar view:", err)
            return err
        }
        log.Println("luxuriouscar view created successfully.")
    } else {
        log.Println("luxuriouscar view already exists.")
    }
    return nil
}



func (r *carRepositoryDB) AddCar(car *models.Cars) (*models.Cars, error) {
	if err := r.db.Create(car).Error; err != nil {
		return nil, err
	}
	return car, nil
}
func (r *carRepositoryDB) GetCarsByClass(class string) ([]models.Cars, error) {
	var cars []models.Cars
	if err := r.db.Table(class).Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}
func (r *carRepositoryDB) GetCarsByOptions(optionNames ...string) ([]models.CarBrandModel, error) {
	var cars []models.CarBrandModel
	// สร้าง query string สำหรับการสร้าง IN clause
	var placeholders []string
	var values []interface{}
	for _, optionName := range optionNames {
		placeholders = append(placeholders, "?")
		values = append(values, optionName)
	}
	// สร้าง query string ที่รวม IN clause
	query := "SELECT brand, model FROM cars WHERE serial_no IN (SELECT serial_no FROM options WHERE option_name IN (" + strings.Join(placeholders, ",") + "))"
	// ใช้ Raw SQL Query ในการดึงข้อมูลรถยนต์
	if err := r.db.Raw(query, values...).Scan(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *carRepositoryDB) GetAllCarsWithTotalPrice() ([]models.CarWithTotalPrice, error) {
	var cars []models.CarWithTotalPrice

	// เตรียมคำสั่ง SQL
	query := `
		SELECT
		    c.serial_no,
		    c.brand,
		    c.model,
		    c.price AS car_price,
		    COALESCE(SUM(o.price), 0) AS total_option_price,
		    c.price + COALESCE(SUM(o.price), 0) AS total_price
		FROM
		    cars c
		LEFT JOIN OPTIONS o ON c.serial_no = o.serial_no
		GROUP BY
		    c.serial_no,
		    c.brand,
		    c.model,
		    c.price;
	`

	// คิวรีฐานข้อมูลด้วยคำสั่ง SQL
	if err := r.db.Raw(query).Scan(&cars).Error; err != nil {
		return nil, err
	}

	return cars, nil
}
