package models

// usecase
type CarUsecase interface {
	AddCar(*Cars) (*Cars, error)
	GetCarsByClass(string) ([]Cars, error)
	GetCarsByOptions(...string) ([]CarBrandModel, error)
	GetAllCarsWithTotalPrice() ([]CarWithTotalPrice, error)
}

type SaleUsecase interface {
	AddSale(*Sales) (*Sales, error)
	GetSalesByEmployee(int, int, int) ([]SalespersonResult, error)
	GetMonthlySalesSummary() ([]MonthlySalesSummary, error)
}

// repository
type SaleRepository interface {
	AddSale(*Sales) (*Sales, error)
	GetSalesByEmployee(int, int, int) ([]SalespersonResult, error)
	GetMonthlySalesSummary() ([]MonthlySalesSummary, error)
}

type CarRepository interface {
	AddCar(*Cars) (*Cars, error)
	GetCarsByClass(string) ([]Cars, error)
	GetCarsByOptions(...string) ([]CarBrandModel, error)
	GetAllCarsWithTotalPrice() ([]CarWithTotalPrice, error)
}

type OptionRepository interface {
	Add(*Options) (*Options, error)
}

type SalespersonRepository interface {
	Add(*Salesperson) (*Salesperson, error)
}
