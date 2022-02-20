package app

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AppsConfig struct {
	App *fiber.App
	Log *zap.SugaredLogger
	Cfg *Config
}
