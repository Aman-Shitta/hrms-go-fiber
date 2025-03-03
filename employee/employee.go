package employee

import (
	"fmt"
	"net/http"

	"github.com/Aman-Shitta/hrms-go-fiber/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

func GetEmployees(c *fiber.Ctx) error {
	var employees []Employee
	query := bson.D{}
	cursor, err := database.MGinstance.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	cursor.All(c.Context(), &employees)
	return c.JSON(employees)
}

func CreateEmployee(c *fiber.Ctx) error {
	var employee Employee

	err := c.BodyParser(&employee)
	if err != nil {
		c.Status(http.StatusInternalServerError).SendString("data invalid")
		return err
	}

	inserted, err := database.MGinstance.Db.Collection("employees").InsertOne(c.Context(), &employee)
	if err != nil {
		c.Status(http.StatusInternalServerError).SendString("db insertion error")
		return err
	}

	return c.JSON(map[string]string{"success": fmt.Sprintf("Data inserted %v", inserted.InsertedID)})
}

func GetEmployeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee Employee

	empID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}

	query := bson.D{{Key: "_id", Value: empID}}

	res := database.MGinstance.Db.Collection("employees").FindOne(c.Context(), query)

	if err := res.Decode(&employee); err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}

	return c.JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	empId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("invalid id format: expected hex")

	}
	var employee Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	query := bson.D{{Key: "_id", Value: empId}}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "salary", Value: employee.Salary},
				{Key: "age", Value: employee.Age},
			},
		},
	}

	res := database.MGinstance.Db.Collection(
		"employees",
	).FindOneAndUpdate(
		c.Context(),
		query,
		update,
		options.FindOneAndUpdate().SetReturnDocument(
			options.After,
		),
	)

	if res.Err() != nil {
		return c.Status(http.StatusInternalServerError).SendString(res.Err().Error())
	}

	err = res.Decode(&employee)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(employee)

}

func DeletEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")

	empId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Status(http.StatusBadRequest).SendString("invalid id : " + err.Error())
	}
	query := bson.D{{Key: "_id", Value: empId}}

	if err := database.MGinstance.Db.Collection("employees").FindOneAndDelete(c.Context(), query).Err(); err != nil {
		c.Status(http.StatusBadRequest).SendString("invalid id probaly : " + err.Error())
	}

	return c.SendStatus(http.StatusOK)
}
