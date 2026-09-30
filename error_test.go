package gotsrpc_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/foomo/gotsrpc/v3"
	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// customError is a value-receiver error used to exercise reflection paths.
type customError struct {
	Code int
	Msg  string
}

func (c customError) Error() string { return c.Msg }

// ptrError is a pointer-receiver error used to exercise the typed-nil guard.
type ptrError struct{ Msg string }

func (p *ptrError) Error() string { return p.Msg }

func TestNewError(t *testing.T) {
	t.Parallel()

	t.Run("typed nil returns nil", func(t *testing.T) {
		t.Parallel()

		var (
			typed *ptrError
			err   error = typed
		)
		assert.Nil(t, gotsrpc.NewError(err))
	})

	t.Run("already an *Error is returned as-is", func(t *testing.T) {
		t.Parallel()

		orig := gotsrpc.NewError(errors.New("boom"))
		require.NotNil(t, orig)
		assert.Same(t, orig, gotsrpc.NewError(orig))
	})

	t.Run("plain error populates fields", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(customError{Code: 42, Msg: "kaboom"})
		require.NotNil(t, e)
		assert.Equal(t, "kaboom", e.Msg)
		assert.Equal(t, "gotsrpc_test.customError", e.Type)
		assert.Contains(t, e.Pkg, "foomo/gotsrpc")
		assert.Equal(t, customError{Code: 42, Msg: "kaboom"}, e.Data)
	})

	t.Run("withStack wrapper is skipped", func(t *testing.T) {
		t.Parallel()

		inner := customError{Code: 1, Msg: "inner"}
		e := gotsrpc.NewError(pkgerrors.WithStack(inner))
		require.NotNil(t, e)
		assert.Equal(t, "inner", e.Msg)
		assert.Equal(t, "gotsrpc_test.customError", e.Type)
	})

	t.Run("single wrapped error sets cause and trims msg", func(t *testing.T) {
		t.Parallel()

		inner := errors.New("root cause")
		e := gotsrpc.NewError(fmt.Errorf("wrapper: %w", inner))
		require.NotNil(t, e)
		require.NotNil(t, e.ErrCause)
		assert.Equal(t, "wrapper", e.Msg)
		assert.Equal(t, "root cause", e.ErrCause.Msg)
	})

	t.Run("joined errors populate ErrCauses", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(errors.Join(errors.New("a"), errors.New("b")))
		require.NotNil(t, e)
		require.Len(t, e.ErrCauses, 2)
		assert.Equal(t, "a", e.ErrCauses[0].Msg)
		assert.Equal(t, "b", e.ErrCauses[1].Msg)
	})
}

func TestError_As(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver returns false", func(t *testing.T) {
		t.Parallel()

		var (
			e      *gotsrpc.Error
			target customError
		)
		assert.False(t, e.As(&target))
	})

	t.Run("nil target returns false", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(customError{Msg: "x"})
		require.NotNil(t, e)
		assert.False(t, e.As(nil))
	})

	t.Run("type mismatch returns false", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(customError{Msg: "x"})
		require.NotNil(t, e)

		var target ptrError
		assert.False(t, e.As(&target))
	})

	t.Run("matching type decodes data", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(customError{Code: 7, Msg: "boom"})
		require.NotNil(t, e)

		var target customError
		require.True(t, e.As(&target))
		assert.Equal(t, 7, target.Code)
		assert.Equal(t, "boom", target.Msg)
	})
}

func TestError_Is(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver returns false", func(t *testing.T) {
		t.Parallel()

		var e *gotsrpc.Error
		assert.False(t, e.Is(customError{Msg: "x"}))
	})

	t.Run("nil arg returns false", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(customError{Msg: "x"})
		require.NotNil(t, e)
		assert.False(t, e.Is(nil))
	})

	t.Run("matching msg/type/pkg returns true", func(t *testing.T) {
		t.Parallel()

		target := customError{Code: 1, Msg: "same"}
		e := gotsrpc.NewError(target)
		require.NotNil(t, e)
		assert.True(t, e.Is(target))
		assert.ErrorIs(t, e, target)
	})

	t.Run("differing message returns false", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(customError{Msg: "one"})
		require.NotNil(t, e)
		assert.False(t, e.Is(customError{Msg: "two"}))
	})
}

func TestError_Cause(t *testing.T) {
	t.Parallel()

	t.Run("returns cause when present", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(fmt.Errorf("outer: %w", errors.New("inner")))
		require.NotNil(t, e)
		assert.Equal(t, e.ErrCause, e.Cause())
	})

	t.Run("returns nil when no cause", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(errors.New("solo"))
		require.NotNil(t, e)
		assert.NoError(t, e.Cause())
		assert.NoError(t, pkgerrors.Cause(e))
	})
}

func TestError_Unwrap(t *testing.T) {
	t.Parallel()

	t.Run("nil receiver returns nil", func(t *testing.T) {
		t.Parallel()

		var e *gotsrpc.Error
		assert.Nil(t, e.Unwrap())
	})

	t.Run("no causes returns nil", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(errors.New("solo"))
		require.NotNil(t, e)
		assert.Nil(t, e.Unwrap())
	})

	t.Run("single cause", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(fmt.Errorf("outer: %w", errors.New("inner")))
		require.NotNil(t, e)
		assert.Len(t, e.Unwrap(), 1)
	})

	t.Run("joined causes", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(errors.Join(errors.New("a"), errors.New("b")))
		require.NotNil(t, e)
		assert.Len(t, e.Unwrap(), 2)
	})
}

func TestError_Error(t *testing.T) {
	t.Parallel()

	t.Run("no cause", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(errors.New("solo"))
		require.NotNil(t, e)
		assert.Equal(t, "solo", e.Error())
	})

	t.Run("with cause chain", func(t *testing.T) {
		t.Parallel()

		e := gotsrpc.NewError(fmt.Errorf("outer: %w", errors.New("inner")))
		require.NotNil(t, e)
		assert.Equal(t, "outer: inner", e.Error())
	})
}

func TestError_Format(t *testing.T) {
	t.Parallel()

	e := gotsrpc.NewError(customError{Code: 3, Msg: "boom"})
	require.NotNil(t, e)

	assert.Equal(t, "boom", fmt.Sprintf("%s", e))
	// Format writes the raw message for both %s and %q (no quoting).
	assert.Equal(t, "boom", fmt.Sprintf("%q", e))
	assert.Equal(t, "boom", fmt.Sprintf("%v", e))

	verbose := fmt.Sprintf("%+v", e)
	assert.Contains(t, verbose, "gotsrpc_test.customError")
	assert.Contains(t, verbose, "Data:")
	assert.Contains(t, verbose, "boom")
}
