package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		auth := c.Get("Authorization")
		if auth == "" {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": "missing token"})
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": "invalid token"})
		}

		claims := token.Claims.(jwt.MapClaims)

		// store in request context
		c.Locals("user_id", claims["id"])
		c.Locals("user_email", claims["email"])
		c.Locals("user_type", claims["user_type"])

		return c.Next()
	}
}
