package usecase

import (
	"errors"

	"testBackend/logs"
	"testBackend/modules/entities/models"
	"testBackend/pkg/errs"
)

type salesService struct {
	salesRepo models.SaleRepository
}

func NewsalesService(
	salesRepo models.SaleRepository,
) models.SaleUsecase {

	return &salesService{
		salesRepo,
	}
}

func (s *salesService) AddSale(req *models.Sales) (sale *models.Sales, err error) {
	sale, err = s.salesRepo.AddSale(req)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return sale, nil
}

func (s *salesService) GetSalesByEmployee(month int, year int, minSales int) (sales []models.SalespersonResult, err error) {
	sales, err = s.salesRepo.GetSalesByEmployee(month, year, minSales)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("couldn't get sales data")
	}
	return sales, nil
}

func (s *salesService) GetMonthlySalesSummary() (sales []models.MonthlySalesSummary, err error) {
	sales, err = s.salesRepo.GetMonthlySalesSummary()
	if err != nil {
		logs.Error(err)
		return nil, errors.New("couldn't get sales data")
	}
	return sales, nil
}
