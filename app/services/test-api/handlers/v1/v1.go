package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icy37785/go-frame/app/services/test-api/handlers/v1/testgrp"
	"github.com/icy37785/go-frame/internal/core/test"
	"github.com/icy37785/go-frame/pkg/app"
	"github.com/icy37785/go-frame/pkg/middleware"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log *zap.SugaredLogger
	Orm *gorm.DB
	DB  *sqlx.DB
}

func Routes(c Config) *fiber.App {
	f := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	tgh := testgrp.Handlers{
		Test: test.NewCore(c.Log, c.Orm, c.DB),
	}

	// 使用中间件
	f.Use(middleware.Recover())
	f.Use(middleware.Favicon())
	f.Use(middleware.Cors())
	f.Use(middleware.RequestID())
	f.Use(middleware.Logging(c.Log))

	// load app.yaml router

	f.Get("/ping", tgh.Ping)
	f.Get("/test", tgh.Query)
	f.Get("/health", app.HealthCheck)

	// 404 Handler
	f.Use(func(c *fiber.Ctx) error {
		return app.RouteNotFound(c) // => 404 "Not Found"
	})

	return f
}
