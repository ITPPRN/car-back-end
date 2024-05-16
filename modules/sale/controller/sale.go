package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"testBackend/logs"
	"testBackend/modules/entities/models"
)

type saleHandler struct {
	saleSrv models.SaleUsecase
}

func NewsaleController(router fiber.Router, saleSrv models.SaleUsecase) {
	controllers := &saleHandler{
		saleSrv: saleSrv,
	}
	router.Post("/", controllers.AddSale)
	router.Get("sales/employee", controllers.GetSalesByEmployee)
	router.Get("sales/summary", controllers.GetMonthlySalesSummary)

}

func (h *saleHandler) AddSale(c *fiber.Ctx) error {

	var newSales models.Sales
	if err := c.BodyParser(&newSales); err != nil {
		logs.Info("Invalid request: " + err.Error())
		return badReqErrResponse(c, "Invalid request: "+err.Error())
	}

	m, err := h.saleSrv.AddSale(&newSales)
	if err != nil {
		return responseWithError(c, err)
	}

	return responseSuccess(c, m)
}

func (h *saleHandler) GetSalesByEmployee(c *fiber.Ctx) error {

	// รับพารามิเตอร์จาก query string
	month, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid month")
	}

	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid year")
	}

	minSales, err := strconv.Atoi(c.Query("minSales"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid minSales")
	}

	res, err := h.saleSrv.GetSalesByEmployee(month, year, minSales)
	if err != nil {
		return responseWithError(c, err)
	}
	return responseSuccess(c, res)
}

func (h *saleHandler) GetMonthlySalesSummary(c *fiber.Ctx) error {

	// เรียกใช้งาน profileService.GetByRole ส่งค่า roles ที่ได้รับมา
	res, err := h.saleSrv.GetMonthlySalesSummary()
	if err != nil {
		// หากเกิดข้อผิดพลาดในการดึงข้อมูลโปรไฟล์ด้วย roles ที่ระบุให้คืนค่า Not Found
		return responseWithError(c, err)
	}

	// หากไม่เกิดข้อผิดพลาดให้คืนค่าข้อมูลโปรไฟล์ที่ได้รับมาในรูปแบบ JSON
	return responseSuccess(c, res)
}
