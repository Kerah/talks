package analrepo

import (
	"context"
	"fmt"

	"talks/ostrovok_ru/intro_pdd/pkg/resql"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type DB interface {
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
}

type Repository struct {
	db DB
}

func New(db DB) *Repository {
	return &Repository{
		db: db,
	}
}

/*
// StartRepoCountV1 OMIT
func (repo *Repository) CountDepositsByUser(
	ctx context.Context,
	id uuid.UUID,
) (deposit int, err error) {
	query, args, err := SQL().
		Select("count(id) as total").
		From("intro.accounts").
		Where(sq.Eq{"user_id": id}).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building select query: %w", err)
	}

	deposit, err = pgxutil.SelectRow(ctx, repo.db, query, args, pgx.RowTo[int])
	if err != nil {
		return 0, fmt.Errorf("error executing select query: %w", err)
	}

	return deposit, nil
}

// EndRepoCountV1 OMIT
*/
// StartRepoCountV2 OMIT
func (repo *Repository) CountDepositsByUserV2(
	ctx context.Context,
	id uuid.UUID,
) (deposit int, err error) {
	query := resql.SQL().
		Select("count(id) as total").
		From("intro.accounts").
		Where(sq.Eq{"user_id": id})
	deposit, err = resql.SelectRow[int](ctx, repo.db, query, pgx.RowTo[int])
	if err != nil {
		return 0, fmt.Errorf("error executing select query: %w", err)
	}
	return deposit, nil
}

// EndRepoCountV2 OMIT
