package routes

import (
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/stevenwr92/absensi/handler"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/auth")

	auth.Get("/", handler.Accessible)
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)

	att := api.Group("/att")

	api.Use(jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte("secret")},
		ErrorHandler: jwtError,
	}))

	auth.Get("/restricted", handler.Restricted)
	att.Get("/", handler.GetAttendance)
	att.Post("clock-in", handler.ClockIn)
	att.Post("clock-out", handler.ClockOut)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
