package usecase

import (
	"errors"

	"testBackend/logs"
	"testBackend/modules/entities/models"
	"testBackend/pkg/errs"
)

type carsService struct {
	carsRepo models.CarRepository
}

func NewcarsService(
	carsRepo models.CarRepository,
) models.CarUsecase {

	return &carsService{
		carsRepo,
	}
}

func (s *carsService) AddCar(req *models.Cars) (car *models.Cars, err error) {
	car, err = s.carsRepo.AddCar(req)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return car, nil
}

func (s *carsService) GetCarsByClass(req string) (cars []models.Cars, err error) {
	cars, err = s.carsRepo.GetCarsByClass(req)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("couldn't get cars data")
	}
	return cars, nil
}

func (s *carsService) GetCarsByOptions(req ...string) (cars []models.CarBrandModel, err error) {
	cars, err = s.carsRepo.GetCarsByOptions(req...)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("couldn't get cars data")
	}
	return cars, nil
}

func (s *carsService) GetAllCarsWithTotalPrice() (cars []models.CarWithTotalPrice, err error) {
	cars, err = s.carsRepo.GetAllCarsWithTotalPrice()
	if err != nil {
		logs.Error(err)
		return nil, errors.New("couldn't get cars data")
	}
	return cars, nil
}


