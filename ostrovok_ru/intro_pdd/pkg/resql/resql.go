package resql

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgxutil"
)

var ErrInvalidSql = errors.New("invalid sql")

// StartIF OMIT
type DB interface {
	Query(
		ctx context.Context,
		sql string,
		optionsAndArgs ...interface{},
	) (pgx.Rows, error)
}

// EndIF OMIT

// SQL returns initialized statement builder.
func SQL() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

// StartSelectRow OMIT
func SelectRow[T any](
	ctx context.Context,
	db DB,
	sql sq.SelectBuilder,
) (T, error) {
	query, args, err := sql.ToSql()
	if err != nil {
		var zero T
		return zero, fmt.Errorf("invalid builded sql: %w: %w", ErrInvalidSql, err)
	}
	res, err := pgxutil.SelectRow(ctx, db, query, args, pgx.RowTo[T])
	if err != nil {
		return res, fmt.Errorf("unexpected error: %w", err)
	}
	return res, nil
}

// EndSelectRow OMIT
