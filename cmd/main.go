package main

import (
	"log"
	"time"

	"romeo-lima/config"
	"romeo-lima/internal/customers/delivery"
	"romeo-lima/internal/customers/repository"
	"romeo-lima/internal/customers/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()
	app := fiber.New()

	app.Use(limiter.New(
		limiter.Config{
			Max:        1,
			Expiration: 30 * time.Second, // reset after this duration
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.IP()
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(
					fiber.Map{
						"error": "Too many request, please try again later.",
					})
			},
		},
	))

	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	customerRepo := repository.NewPostgresCustomerRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	delivery.NewCustomerHandler(app, customerUsecase)

	log.Fatal(app.Listen(":3000"))
}
