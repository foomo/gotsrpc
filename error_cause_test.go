package gotsrpc //nolint:testpackage

import (
	"errors"
	"fmt"
	"testing"
	"time"

	pkgerrors "github.com/pkg/errors"
)

func requireTerminates(t *testing.T, name string, fn func()) {
	t.Helper()

	done := make(chan struct{})

	go func() {
		defer close(done)

		fn()
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatalf("%s did not terminate within 5s: Cause() is returning the receiver", name)
	}
}

func TestErrorCauseTerminates(t *testing.T) {
	t.Parallel()

	err := NewError(errors.New("boom"))
	if err.ErrCause != nil {
		t.Fatalf("precondition failed: expected no wrapped cause, got %v", err.ErrCause)
	}

	requireTerminates(t, "errors.Cause", func() { _ = pkgerrors.Cause(err) })
}

func TestErrorCauseDoesNotReturnReceiver(t *testing.T) {
	t.Parallel()

	err := NewError(errors.New("boom"))

	if got := err.Cause(); errors.Is(got, error(err)) {
		t.Fatal("Cause() returned the receiver; pkg/errors.Cause will never terminate")
	}
}

func TestErrorCauseUntypedNil(t *testing.T) {
	t.Parallel()

	err := &Error{Msg: "boom"}

	if got := err.Cause(); got != nil {
		t.Fatalf("Cause() = %#v, want untyped nil", got)
	}
}

func TestErrorCauseReturnsWrappedCause(t *testing.T) {
	t.Parallel()

	err := NewError(fmt.Errorf("outer: %w", errors.New("inner")))

	got := err.Cause()
	if got == nil {
		t.Fatal("Cause() = nil, want the wrapped cause")
	}

	if got.Error() != "inner" {
		t.Fatalf("Cause().Error() = %q, want %q", got.Error(), "inner")
	}
}
