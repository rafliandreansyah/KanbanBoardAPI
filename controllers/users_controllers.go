package controllers

import (
	"KanbanBoardAPI/database"
	"KanbanBoardAPI/helpers"
	"KanbanBoardAPI/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error{
	var err error

	db := database.GetDB()
	user := new(models.User)

	if helpers.ContentType(c) != fiber.MIMEApplicationJSON && helpers.ContentType(c) != fiber.MIMEApplicationForm {
		_ = c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "bad request",
		})
	}

	err = c.BodyParser(&user)
	if err != nil {
		_ = c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "objek tidak sesuai",
		})
	}


	err = db.Debug().Create(&user).Error
	if err != nil {
		_ = c.SendStatus(fiber.StatusPaymentRequired)

		return c.JSON(fiber.Map{
			"message": "registrasi gagal",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "registrasi berhasil",
		"data" : fiber.Map{
			"email": user.Email,
			"full_name": user.FullName,
		},
	})

}