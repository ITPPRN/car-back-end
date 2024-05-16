package models

//req

type CarRequest struct {
	SerialNo     int
	Brand        string
	Model        string
	Manufacturer string
	Price        float64
}

type OptionsRequest struct {
	SerialNo   int
	OptionName string
	Price      float64
}

type SalespersonRequest struct {
	SalePersonID int
	Name         string
	Phone        string
}

type SalesRequest struct {
	SerialNo     int
	SalePersonID int
	Price        float64
	Day          int
	Month        int
	Year         int
}

// resonse
type CarResponse struct {
	SerialNo     int
	Brand        string
	Model        string
	Manufacturer string
	Price        float64
}

type CarByOptionResponse struct {
	Brand string
	Model string
}

type MonthlySalesSummary struct {
	Month       int     `json:"Month"`
	Year        int     `json:"Year"`
	NumCarsSold int     `json:"num_cars_sold"`
	TotalSales  float64 `json:"total_sales"`
}

type CarWithTotalPrice struct {
	SerialNo         int     `json:"serial_no"`
	Brand            string  `json:"brand"`
	Model            string  `json:"model"`
	CarPrice         float64 `json:"car_price"`
	TotalOptionPrice float64 `json:"total_option_price"`
	TotalPrice       float64 `json:"total_price"`
}

type CarBrandModel struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type SalespersonResult struct {
	Name        string `json:"name"`
	NumCarsSold int    `json:"num_cars_sold"`
}

type ResponseData struct {
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

type ResponseError struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
}
