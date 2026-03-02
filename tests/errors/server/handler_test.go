package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/errors/server"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	h := server.New()
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", nil)

	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		err := h.Error(w, r)
		assert.Error(t, err)
	})

	t.Run("Scalar", func(t *testing.T) {
		t.Parallel()
		ret := h.Scalar(w, r)
		assert.NotNil(t, ret)
	})

	t.Run("MultiScalar", func(t *testing.T) {
		t.Parallel()
		ret := h.MultiScalar(w, r)
		assert.NotNil(t, ret)
	})

	t.Run("Struct", func(t *testing.T) {
		t.Parallel()
		ret := h.Struct(w, r)
		assert.NotNil(t, ret)
	})

	t.Run("StructError", func(t *testing.T) {
		t.Parallel()
		err := h.StructError(w, r)
		assert.Error(t, err)
	})

	t.Run("TypedError", func(t *testing.T) {
		t.Parallel()
		err := h.TypedError(w, r)
		assert.Error(t, err)
	})

	t.Run("ScalarError", func(t *testing.T) {
		t.Parallel()
		err := h.ScalarError(w, r)
		assert.Error(t, err)
	})

	t.Run("CustomError", func(t *testing.T) {
		t.Parallel()
		err := h.CustomError(w, r)
		assert.Error(t, err)
	})

	t.Run("WrappedError", func(t *testing.T) {
		t.Parallel()
		err := h.WrappedError(w, r)
		assert.Error(t, err)
	})

	t.Run("TypedWrappedError", func(t *testing.T) {
		t.Parallel()
		err := h.TypedWrappedError(w, r)
		assert.Error(t, err)
	})

	t.Run("TypedScalarError", func(t *testing.T) {
		t.Parallel()
		err := h.TypedScalarError(w, r)
		assert.Error(t, err)
	})

	t.Run("TypedCustomError", func(t *testing.T) {
		t.Parallel()
		err := h.TypedCustomError(w, r)
		assert.Error(t, err)
	})
}
