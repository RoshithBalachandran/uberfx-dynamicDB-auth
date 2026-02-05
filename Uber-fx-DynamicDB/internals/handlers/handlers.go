package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roshith/dynamicDB/internals/service"
)

type AuthHandlers struct {
	authservice *service.AuthService
}

func NewHandlers(auth *service.AuthService) *AuthHandlers {
	return &AuthHandlers{authservice: auth}
}

func (h *AuthHandlers) Register(c *fiber.Ctx) error {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		UserType string `json:"user_type"`
	}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid structured data"})
	}
	user, err := h.authservice.Register(body.Name, body.Email, body.Password, body.UserType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"sucess": "Registerd sucessfull", "user": user})
}

func (h *AuthHandlers) Login(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		UserType string `json:"user_type"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid strutured data"})
	}
	token, err := h.authservice.Login(body.Email, body.Password, body.UserType)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"sucess": "Login sucessfull", "Token": token})
}

func (h *AuthHandlers) Profile(c *fiber.Ctx) error {

	idFloat := c.Locals("user_id").(float64)
	userType := c.Locals("user_type").(string)

	id := uint(idFloat)

	user, err := h.authservice.GetProfile(id, userType)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(fiber.Map{
		"profile": user,
		"db_used": userType,
	})
}
