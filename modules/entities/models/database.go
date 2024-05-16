package models

type Cars struct {
	SerialNo     int `gorm:"primaryKey"`
	Brand        string
	Model        string
	Manufacturer string
	Price        float64
}

type Options struct {
	SerialNo   int `gorm:"foreignKey:SerialNo"`
	OptionName string
	Price      float64
}

type Salesperson struct {
	SalePersonID int    `gorm:"primaryKey" json:"sale_person_id"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
}

// TableName ระบุชื่อตารางสำหรับโมเดล Salesperson
func (Salesperson) TableName() string {
    return "salespersons"
}

type Sales struct {
	SerialNo     int `gorm:"foreignKey:SerialNo"`
	SalePersonID int `gorm:"foreignKey:SalePersonID"`
	Price        float64
	Day          int
	Month        int
	Year         int
}
