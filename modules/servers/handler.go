package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"testBackend/logs"
	_carCon "testBackend/modules/car/controller"
	_carRe "testBackend/modules/car/repository"
	_carUsecase "testBackend/modules/car/usecase"
	_optionRe "testBackend/modules/option/repository"
	"testBackend/modules/sale/controller"
	"testBackend/modules/sale/repository"
	"testBackend/modules/sale/usecase"
	_salesPersonRe "testBackend/modules/sale_person/repository"
)

func (s *server) Handlers() error {

	// Group a version
	v1 := s.App.Group("/v1")
	v1.Use(cors.New(cors.Config{
		AllowOrigins: s.Cfg.App.AllowOrigins,
	}))
	v1.Use(logs.LogHttp)

	carGroup := v1.Group("/car")
	saleGroup := v1.Group("/sale")
	//repo
	carRepo := _carRe.NewoptionRepositoryDB(s.Db)
	salesPersonRepo := _salesPersonRe.NewsalespersonRepositoryDB(s.Db)
	optionRepo := _optionRe.NewoptionRepositoryDB(s.Db)
	saleRepo := repository.NewsaleRepositoryDB(s.Db)

	_ = salesPersonRepo
	_= optionRepo

	//usecase
	carUse := _carUsecase.NewcarsService(carRepo)
	saleUse := usecase.NewsalesService(saleRepo)

	_carCon.NewcarController(carGroup, carUse)
	controller.NewsaleController(saleGroup,saleUse)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil

}
