package model_test

import (
	"sort"
	"testing"

	"github.com/foomo/gotsrpc/v3/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestFullName(t *testing.T) {
	t.Parallel()

	t.Run("scalar", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "time.Time", (&model.Scalar{Package: "time", Name: "Time"}).FullName())
		assert.Equal(t, ".Time", (&model.Scalar{Name: "Time"}).FullName())
	})

	t.Run("struct type", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "service.User", (&model.StructType{Package: "service", Name: "User"}).FullName())
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "service.User", (&model.Struct{Package: "service", Name: "User"}).FullName())
	})
}

func TestServiceList_Sort(t *testing.T) {
	t.Parallel()

	list := model.ServiceList{
		{Name: "Charlie"},
		{Name: "Alpha"},
		{Name: "Bravo"},
	}
	assert.Equal(t, 3, list.Len())

	sort.Sort(list)
	assert.Equal(t, "Alpha", list[0].Name)
	assert.Equal(t, "Bravo", list[1].Name)
	assert.Equal(t, "Charlie", list[2].Name)
}

func TestServiceMethods_Sort(t *testing.T) {
	t.Parallel()

	methods := model.ServiceMethods{
		{Name: "Zebra"},
		{Name: "Apple"},
		{Name: "Mango"},
	}
	assert.Equal(t, 3, methods.Len())

	sort.Sort(methods)
	assert.Equal(t, "Apple", methods[0].Name)
	assert.Equal(t, "Mango", methods[1].Name)
	assert.Equal(t, "Zebra", methods[2].Name)
}
