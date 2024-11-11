package analrepo

import (
	"context"
	"testing"

	"github.com/godepo/groat"
	"github.com/godepo/groat/integration"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Deps struct {
	Faker    faker.Faker
	MockDB   *MockDB
	MockRows *MockRows
	DB       *pgxpool.Pool `groat:"pgxpool"`
}
type State struct {
	Faker  faker.Faker
	ctx    context.Context
	MockDB *MockDB
}

var suite *integration.Container[Deps, State, *Repository]

/*
func TestRepository_CountDepositsByUser(t *testing.T) {
	// StartRepoCountV1Empty OMIT
	t.Run("should be able return empty count for unknown user state",
		func(t *testing.T) {
			tc := suite.Case(t)
			cnt, err := tc.SUT.
				CountDepositsByUser(context.Background(), uuid.New())
			require.NoError(t, err)
			assert.Zero(t, cnt)
		})
	// EndRepoCountV1Empty OMIT
	t.Run("should be able return error", func(t *testing.T) {
		tc := suite.Case(t)
		userId := uuid.New()
		tc.When(InjectMocks(tc.SUT))
		tc.Deps.MockDB.EXPECT().Query(mock.Anything, "", userId).Return()
	})
}
*/

func TestRepository_CountDepositsByUserV2(t *testing.T) {
	t.Run("should be able return empty count for unknown user state",
		func(t *testing.T) {
			tc := suite.Case(t)
			cnt, err := tc.SUT.
				CountDepositsByUserV2(context.Background(), uuid.New())
			require.NoError(t, err)
			assert.Zero(t, cnt)
		})
}

func InjectMocks(sut *Repository) groat.When[Deps, State] {
	return func(t *testing.T, deps Deps, state State) State {
		sut.db = deps.MockDB
		return state
	}
}
