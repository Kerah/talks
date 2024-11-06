package analrepo

import (
	"context"
	"os"
	"testing"

	"github.com/godepo/groat"
	"github.com/godepo/groat/integration"
	"github.com/godepo/pgrx"
	"github.com/jaswdr/faker/v2"
)

func mainProvider(t *testing.T) *groat.Case[Deps, State, *Repository] {
	tcs := groat.New[Deps, State, *Repository](t, func(t *testing.T, deps Deps) *Repository {
		return New(deps.DB)
	})
	tcs.Before(func(t *testing.T, deps Deps) Deps {
		deps.Faker = faker.New()
		deps.MockDB = NewMockDB(t)
		deps.MockRows = NewMockRows(t)
		return deps
	})
	tcs.Given(func(t *testing.T, state State) State {
		state.Faker = tcs.Deps.Faker
		state.ctx = context.Background()
		tcs.State.MockDB = tcs.Deps.MockDB
		return state
	})
	return tcs
}

func TestMain(m *testing.M) {
	suite = integration.New[Deps, State, *Repository](m, mainProvider,
		pgrx.New[Deps](
			pgrx.WithContainerImage("docker.io/postgres:16"),
			pgrx.WithMigrationsPath("./sql"),
		),
	)
	os.Exit(suite.Go())
}
