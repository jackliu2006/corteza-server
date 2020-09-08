package helpers

import (
	"context"
	"github.com/cortezaproject/corteza-server/app"
	"github.com/cortezaproject/corteza-server/pkg/cli"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/rand"
	"github.com/cortezaproject/corteza-server/store/sqlite"
	"time"
)

func NewIntegrationTestApp(ctx context.Context, initTestServices func(*app.CortezaApp) error) *app.CortezaApp {
	var (
		a   = app.New()
		err error
	)

	// Make sure SQLite is registered
	a.Store, err = sqlite.ConnectInMemory(ctx)
	cli.HandleError(err)

	// When running integration tests, we want to upgrade the db. Always.
	a.Opt.Upgrade.Always = true

	// Create a new JWT secret (to prevent any security weirdness)
	a.Opt.Auth.Secret = string(rand.Bytes(32))
	a.Opt.Auth.Expiry = time.Minute

	a.Log = logger.MakeDebugLogger()

	cli.HandleError(a.InitStore(ctx))
	cli.HandleError(initTestServices(a))
	cli.HandleError(a.InitServices(ctx))
	return a
}
