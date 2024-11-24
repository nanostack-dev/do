package tests

import (
	"testing"

	"github.com/nanostack-dev/do"
	"github.com/nanostack-dev/do/tests/fixtures"

	"github.com/stretchr/testify/assert"
)

func TestParallelShutdown(t *testing.T) {
	is := assert.New(t)

	root, driver, passenger := fixtures.GetPackage()
	is.NotPanics(
		func() {
			_ = do.MustInvoke[*fixtures.Driver](driver)
			_ = do.MustInvokeNamed[*fixtures.Passenger](passenger, "passenger-1")
			_ = do.MustInvokeNamed[*fixtures.Passenger](passenger, "passenger-2")
			_ = do.MustInvokeNamed[*fixtures.Passenger](passenger, "passenger-3")
			root.Shutdown()
		},
	)
}
