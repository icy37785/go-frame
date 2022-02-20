package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

func Logging(log *zap.SugaredLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now().UTC()
		log.Infow("request started", "traceid", c.GetRespHeader(fiber.HeaderXRequestID), "method", c.Method(), "path", c.Path(),
			"ip", c.IP())

		// Continue.
		err := c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)
		log.Infow("request completed", "traceid", c.GetRespHeader(fiber.HeaderXRequestID), "method", c.Method(), "path", c.Path(),
			"ip", c.IP(), "statuscode", c.Response().StatusCode(), "since", latency)
		//return c.Next()
		return err
	}
}
