package delivery

import (
	"net/http"
	"romeo-lima/internal/customers/usecase"
	"romeo-lima/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	CUsecase usecase.CustomerUsecase
}

func NewCustomerHandler(app *fiber.App, us usecase.CustomerUsecase) {
	handler := &CustomerHandler{
		CUsecase: us,
	}
	app.Get("/customers/:id", handler.GetCustomer)
	app.Post("/customers", handler.CreateCustomer)
}

func (ch *CustomerHandler) GetCustomer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	customer, err := ch.CUsecase.GetCustomerByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if customer == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}
	return c.JSON(customer)
}

func (ch *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}
	validate := validator.New()
	if err := validate.Struct(customer); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := ch.CUsecase.CreateCustomer(c.Context(), &customer)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(customer)
}
