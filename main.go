package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/stevenwr92/absensi/routes"
	"github.com/stevenwr92/absensi/utils"
)

func main() {

	app := fiber.New()
	utils.ConnectDatabase()
	routes.Routes(app)

	err := app.Listen(":3002")
	if err != nil {
		fmt.Printf("Error :%v", err)
	}
}
