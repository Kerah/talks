package resql

import (
	"context"
	"errors"
	"math/rand"
	"sync/atomic"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSelectRow(t *testing.T) {
	t.Run("should be able to fail", func(t *testing.T) {
		// StartSqlInvalid OMIT
		t.Run("when sql is invalid", func(t *testing.T) {
			dbMock := NewMockDB(t)

			res, err := SelectRow[int](
				context.Background(),
				dbMock,
				SQL().Select(),
				pgx.RowTo[int],
			)
			require.ErrorIs(t, err, ErrInvalidSql)
			assert.Zero(t, res)
		})
		// EndSqlInvalid OMIT
		// StartSuccess OMIT
		t.Run("successful", func(t *testing.T) {
			exp := rand.Int()

			cntNext := &atomic.Bool{}
			cntScan := &atomic.Bool{}

			dbMock := NewMockDB(t)

			rows := NewMockRows(t)
			// EndSuccess OMIT

			// StartSuccessExpects OMIT
			dbMock.EXPECT().Query(mock.Anything, "SELECT count").Return(rows, nil)

			rows.EXPECT().Next().RunAndReturn(func() bool {
				return cntNext.CompareAndSwap(false, true)
			})

			rows.EXPECT().Scan(mock.Anything).Run(func(dest ...interface{}) {
				require.True(t, cntScan.CompareAndSwap(false, true))
				assert.Len(t, dest, 1)
				ptr, ok := dest[0].(*int)
				require.True(t, ok)
				*ptr = exp
			}).Return(nil)

			rows.EXPECT().Close()

			rows.EXPECT().Err().Return(nil)

			rows.EXPECT().CommandTag().Return(pgconn.NewCommandTag("SELECT"))
			// EndSuccessExpects OMIT

			// StartSuccessThen OMIT
			res, err := SelectRow[int](
				context.Background(),
				dbMock,
				SQL().Select("count"),
				pgx.RowTo[int],
			)
			require.NoError(t, err)
			require.Equal(t, exp, res)
			// EndSuccessThen OMIT
		})
		/*
			// StartDBFailV1 OMIT
			t.Run("when fail db", func(t *testing.T) {
				dbMock := NewMockDB(t)
				expErr := errors.New(uuid.NewString())
				dbMock.EXPECT().
					Query(mock.Anything, "SELECT count").
					Return(nil, expErr)

				res, err := SelectRow[int](context.Background(),
					dbMock,
					SQL().Select("count"),
					pgx.RowTo[int],
				)
				require.ErrorIs(t, err, expErr)
				assert.Zero(t, res)
			})
			// EndDBFailV2 OMIT
		*/

		t.Run("when fail db", func(t *testing.T) {
			dbMock := NewMockDB(t)
			expErr := errors.New(uuid.NewString())
			rows := NewMockRows(t)
			dbMock.EXPECT().Query(mock.Anything, "SELECT count").Return(rows, nil)
			rows.EXPECT().Next().Return(false)
			rows.EXPECT().Close()
			rows.EXPECT().Err().Return(expErr)

			_, err := SelectRow[int](
				context.Background(),
				dbMock,
				SQL().Select("count"),
				pgx.RowTo[int],
			)
			require.ErrorIs(t, err, expErr)
		})
	})
}
