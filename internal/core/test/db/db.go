package db

import (
	"context"
	"fmt"
	"github.com/icy37785/go-frame/pkg/storage/sql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Store manages the set of APIs for user access.
type Store struct {
	log *zap.SugaredLogger
	//tr           database.Transactor
	orm          *gorm.DB
	db           sqlx.ExtContext
	isWithinTran bool
}

// NewStore constructs a data for api access.
func NewStore(log *zap.SugaredLogger, orm *gorm.DB, db *sqlx.DB) Store {
	return Store{
		log: log,
		//tr:  db,
		orm: orm,
		db:  db,
	}
}

// Query retrieves a list of existing searches from the database.
func (s Store) Query(ctx context.Context, pageNumber int, rowsPerPage int) ([]Test, error) {
	data := struct {
		Offset      int `db:"offset"`
		RowsPerPage int `db:"rows_per_page"`
	}{
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
	}
	const q = `
	SELECT
		id,title
	FROM
		test
	ORDER BY
		id DESC
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var tests []Test
	if err := sql.NamedQuerySlice(ctx, s.log, s.db, q, data, &tests); err != nil {
		return nil, fmt.Errorf("selecting tests: %w", err)
	}

	return tests, nil
}
