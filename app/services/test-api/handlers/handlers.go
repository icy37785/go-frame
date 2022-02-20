package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icy37785/go-frame/app/services/test-api/handlers/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Config struct {
	Log *zap.SugaredLogger
	Orm *gorm.DB
	DB  *sqlx.DB
}

func NewHTTPServer(cfg Config) *fiber.App {
	return v1.Routes(v1.Config{
		Log: cfg.Log,
		Orm: cfg.Orm,
		DB:  cfg.DB,
	})
}
