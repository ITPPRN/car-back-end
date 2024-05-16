package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"testBackend/logs"
	"testBackend/modules/entities/models"
)

type carHandler struct {
	carSrv models.CarUsecase
}

func NewcarController(router fiber.Router, carSrv models.CarUsecase) {
	controllers := &carHandler{
		carSrv: carSrv,
	}
	router.Post("/", controllers.addCar)
	router.Get("cars/class/:class", controllers.GetCarsByClass)
	router.Get("cars/option", controllers.GetCarsByOptions)
	router.Get("cars/total", controllers.GetAllCarsWithTotalPrice)

}

func (h *carHandler) addCar(c *fiber.Ctx) error {

	var newCar models.Cars
	if err := c.BodyParser(&newCar); err != nil {
		logs.Info("Invalid request: " + err.Error())
		return badReqErrResponse(c, "Invalid request: "+err.Error())
	}

	m, err := h.carSrv.AddCar(&newCar)
	if err != nil {
		return responseWithError(c, err)
	}

	return responseSuccess(c, m)
}

func (h *carHandler) GetCarsByClass(c *fiber.Ctx) error {

	class := c.Params("class")
	if class == "" {
		return badReqErrResponse(c, "Require parameters")
	}

	res, err := h.carSrv.GetCarsByClass(class)
	if err != nil {
		return responseWithError(c, err)
	}
	return responseSuccess(c, res)
}

func (h *carHandler) GetCarsByOptions(c *fiber.Ctx) error {

	// ดึงค่า roles จากพารามิเตอร์ของ request
	roles := strings.Split(c.Query("option"), ",")

	// เรียกใช้งาน profileService.GetByRole ส่งค่า roles ที่ได้รับมา
	res, err := h.carSrv.GetCarsByOptions(roles...)
	if err != nil {
		// หากเกิดข้อผิดพลาดในการดึงข้อมูลโปรไฟล์ด้วย roles ที่ระบุให้คืนค่า Not Found
		return responseWithError(c, err)
	}

	// หากไม่เกิดข้อผิดพลาดให้คืนค่าข้อมูลโปรไฟล์ที่ได้รับมาในรูปแบบ JSON
	return responseSuccess(c, res)
}

func (h *carHandler) GetAllCarsWithTotalPrice(c *fiber.Ctx) error {

	// เรียกใช้งาน profileService.GetByRole ส่งค่า roles ที่ได้รับมา
	res, err := h.carSrv.GetAllCarsWithTotalPrice()
	if err != nil {
		// หากเกิดข้อผิดพลาดในการดึงข้อมูลโปรไฟล์ด้วย roles ที่ระบุให้คืนค่า Not Found
		return responseWithError(c, err)
	}

	// หากไม่เกิดข้อผิดพลาดให้คืนค่าข้อมูลโปรไฟล์ที่ได้รับมาในรูปแบบ JSON
	return responseSuccess(c, res)
}
