package testgrp

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/icy37785/go-frame/internal/core/test"
	"github.com/icy37785/go-frame/pkg/app"
	"strconv"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Test test.Core
}

func (h *Handlers) Ping(ctx *fiber.Ctx) error {
	return app.Success(ctx, nil)
}

func (h *Handlers) Query(ctx *fiber.Ctx) error {
	page := ctx.Query("page", "1")

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		return app.Error(ctx, err)
	}
	rows := ctx.Query("rows", "10")
	rowsPerPage, err := strconv.Atoi(rows)
	if err != nil {
		return app.Error(ctx, err)
	}

	tests, err := h.Test.Query(context.Background(), pageNumber, rowsPerPage)
	if err != nil {
		return fmt.Errorf("unable to query for users: %w", err)
	}
	return app.Success(ctx, tests)
}
