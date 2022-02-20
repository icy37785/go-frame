package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	maxAge = 12
)

func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		//AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, HEAD, PUT, PATCH, POST, DELETE",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           maxAge * 3600,
	})
}
