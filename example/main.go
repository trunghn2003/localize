package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/trunghn2003/localize"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID   uint                      `json:"id" gorm:"primaryKey"`
	Name localize.TranslatableField `json:"name" gorm:"type:jsonb"`
}

var db *gorm.DB

func main() {
	var err error
	dsn := "host=localhost user=postgres password=123 dbname=localize_example port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Product{})

	app := fiber.New()

	app.Post("/products", createProduct)
	app.Get("/products/:id", getProduct)

	log.Fatal(app.Listen(":3000"))
}

func createProduct(c *fiber.Ctx) error {
	var input struct {
		Name map[string]string `json:"name"`
	}

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	product := Product{ Name: input.Name }

	if err := db.Create(&product).Error; err != nil {
		return err
	}

	return c.JSON(product)
}

func getProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product

	if err := db.First(&product, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	// get the locale from the request, default to "en"
	locale := c.Get("Accept-Language", "en")

	return c.JSON(fiber.Map{
		"id":   product.ID,
		"name": product.Name.Get(locale, "en"),
	})
}
