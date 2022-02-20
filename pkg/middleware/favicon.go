package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

func Favicon() fiber.Handler {
	return favicon.New()
}
