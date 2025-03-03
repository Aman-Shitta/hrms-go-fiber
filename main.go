package main

import (
	"log"

	"github.com/Aman-Shitta/hrms-go-fiber/database"
	"github.com/Aman-Shitta/hrms-go-fiber/employee"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// DB setup
	initDB()
	// fiber app
	app := fiber.New()

	app.Route("/api/v1", regsiterRoutes, "employees")
	app.Listen(":3000")
}

func initDB() {
	// db setup
	err := database.NewMongoInstance()

	if err != nil {
		log.Fatal(err)
	}

}

func regsiterRoutes(router fiber.Router) {
	router.Get("/employee", employee.GetEmployees)
	router.Post("/employees", employee.CreateEmployee)
	router.Get("/employee/:id", employee.GetEmployeByID)
	router.Put("/employee/:id", employee.UpdateEmployee)
	router.Delete("/employee/:id", employee.DeletEmployeeById)
}
