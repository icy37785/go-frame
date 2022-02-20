package test

import (
	"context"
	"fmt"
	"github.com/icy37785/go-frame/internal/core/test/db"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Core manages the set of APIs for user access.
type Core struct {
	store db.Store
}

func NewCore(log *zap.SugaredLogger, ormBD *gorm.DB, sqlDB *sqlx.DB) Core {
	return Core{
		store: db.NewStore(log, ormBD, sqlDB),
	}
}

// Query retrieves a list of existing searches from the database.
func (c Core) Query(ctx context.Context, pageNumber int, rowsPerPage int) ([]Test, error) {
	dbSearches, err := c.store.Query(ctx, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return toTestSlice(dbSearches), nil
}
