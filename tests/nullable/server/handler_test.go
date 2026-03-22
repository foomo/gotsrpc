package server_test

import (
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/nullable/server"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	h := &server.Handler{}

	t.Run("VariantA", func(t *testing.T) {
		t.Parallel()
		v := server.Base{B1: "hello", D1: server.ACustomTypeOne}
		assert.Equal(t, v, h.VariantA(t.Context(), v))
	})

	t.Run("VariantB", func(t *testing.T) {
		t.Parallel()
		v := server.BCustomType("test")
		assert.Equal(t, v, h.VariantB(t.Context(), v))
	})

	t.Run("VariantC", func(t *testing.T) {
		t.Parallel()
		v := server.BCustomTypes{"a", "b"}
		assert.Equal(t, v, h.VariantC(t.Context(), v))
	})

	t.Run("VariantD", func(t *testing.T) {
		t.Parallel()
		v := server.BCustomTypesMap{"x": "y"}
		assert.Equal(t, v, h.VariantD(t.Context(), v))
	})

	t.Run("VariantE", func(t *testing.T) {
		t.Parallel()
		v := &server.Base{B1: "ptr"}
		assert.Equal(t, v, h.VariantE(t.Context(), v))
	})

	t.Run("VariantE_nil", func(t *testing.T) {
		t.Parallel()
		assert.Nil(t, h.VariantE(t.Context(), nil))
	})

	t.Run("VariantF", func(t *testing.T) {
		t.Parallel()
		v := []*server.Base{{B1: "one"}, {B1: "two"}}
		assert.Equal(t, v, h.VariantF(t.Context(), v))
	})

	t.Run("VariantG", func(t *testing.T) {
		t.Parallel()
		v := map[string]*server.Base{"k": {B1: "val"}}
		assert.Equal(t, v, h.VariantG(t.Context(), v))
	})

	t.Run("VariantH", func(t *testing.T) {
		t.Parallel()
		i1 := server.Base{B1: "one"}
		i2 := &server.Base{B1: "two"}
		i3 := []*server.Base{{B1: "three"}}
		i4 := map[string]server.Base{"k": {B1: "four"}}
		r1, r2, r3, r4 := h.VariantH(t.Context(), i1, i2, i3, i4)
		assert.Equal(t, i1, r1)
		assert.Equal(t, i2, r2)
		assert.Equal(t, i3, r3)
		assert.Equal(t, i4, r4)
	})
}
