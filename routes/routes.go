package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/stevenwr92/absensi/handler"
	"github.com/stevenwr92/absensi/utils"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/auth")

	auth.Get("/", handler.Accessible)
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)

	att := api.Group("/att")

	api.Use(utils.JWTMiddleware())

	auth.Get("/restricted", handler.Restricted)
	att.Get("/", handler.GetAttendance)
	att.Post("clock-in", handler.ClockIn)
	att.Post("clock-out", handler.ClockOut)
}
