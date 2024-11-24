package do

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////////////////////////////////////////////
// 							Explicit aliases
/////////////////////////////////////////////////////////////////////////////

func TestAs(t *testing.T) {
	is := assert.New(t)

	i := New()
	Provide(
		i,
		func(i Injector) (*lazyTestHeathcheckerOK, error) { return &lazyTestHeathcheckerOK{}, nil },
	)

	is.Nil(As[*lazyTestHeathcheckerOK, Healthchecker](i))
	is.EqualError(
		As[*lazyTestShutdownerOK, Healthchecker](i),
		"DI: `*github.com/nanostack-dev/do.lazyTestShutdownerOK` is not `github.com/nanostack-dev/do.Healthchecker`",
	)
	is.EqualError(
		As[*lazyTestHeathcheckerKO, Healthchecker](i),
		"DI: service `*github.com/nanostack-dev/do.lazyTestHeathcheckerKO` has not been declared",
	)
	is.EqualError(
		As[*lazyTestShutdownerOK, *lazyTestShutdownerOK](i),
		"DI: service `*github.com/nanostack-dev/do.lazyTestShutdownerOK` has not been declared",
	)
}

func TestMustAs(t *testing.T) {
	// @TODO
}

func TestAsNamed(t *testing.T) {
	is := assert.New(t)

	i := New()
	Provide(
		i,
		func(i Injector) (*lazyTestHeathcheckerOK, error) { return &lazyTestHeathcheckerOK{}, nil },
	)

	is.Nil(
		AsNamed[*lazyTestHeathcheckerOK, Healthchecker](
			i, "*github.com/nanostack-dev/do.lazyTestHeathcheckerOK",
			"github.com/nanostack-dev/do.Healthchecker",
		),
	)
	is.EqualError(
		AsNamed[*lazyTestShutdownerOK, Healthchecker](
			i, "*github.com/nanostack-dev/do.lazyTestShutdownerOK",
			"github.com/nanostack-dev/do.Healthchecker",
		),
		"DI: `*github.com/nanostack-dev/do.lazyTestShutdownerOK` is not `github.com/nanostack-dev/do.Healthchecker`",
	)
	is.EqualError(
		AsNamed[*lazyTestHeathcheckerKO, Healthchecker](
			i, "*github.com/nanostack-dev/do.lazyTestHeathcheckerKO",
			"github.com/nanostack-dev/do.Healthchecker",
		), "DI: service `*github.com/nanostack-dev/do.lazyTestHeathcheckerKO` has not been declared",
	)
	is.EqualError(
		AsNamed[*lazyTestShutdownerOK, *lazyTestShutdownerOK](
			i, "*github.com/nanostack-dev/do.lazyTestShutdownerOK",
			"*github.com/nanostack-dev/do.lazyTestShutdownerOK",
		), "DI: service `*github.com/nanostack-dev/do.lazyTestShutdownerOK` has not been declared",
	)
}

func TestMustAsNamed(t *testing.T) {
	// @TODO
}

/////////////////////////////////////////////////////////////////////////////
// 							Implicit aliases
/////////////////////////////////////////////////////////////////////////////

func TestInvokeAs(t *testing.T) {
	is := assert.New(t)

	i := New()
	Provide(
		i, func(i Injector) (*lazyTestHeathcheckerOK, error) {
			return &lazyTestHeathcheckerOK{foobar: "hello world"}, nil
		},
	)

	// found
	svc0, err := InvokeAs[*lazyTestHeathcheckerOK](i)
	is.EqualValues(&lazyTestHeathcheckerOK{foobar: "hello world"}, svc0)
	is.Nil(err)

	// found via interface
	svc1, err := InvokeAs[Healthchecker](i)
	is.EqualValues(&lazyTestHeathcheckerOK{foobar: "hello world"}, svc1)
	is.Nil(err)

	// not found
	svc2, err := InvokeAs[Shutdowner](i)
	is.Empty(svc2)
	is.EqualError(
		err,
		"DI: could not find service satisfying interface `github.com/nanostack-dev/do.Shutdowner`, available services: `*github.com/nanostack-dev/do.lazyTestHeathcheckerOK`",
	)
}

func TestMustInvokeAs(t *testing.T) {
	// @TODO
}
