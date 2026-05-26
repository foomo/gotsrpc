package server

import (
	"context"

	"github.com/foomo/gotsrpc/v2/tests/common"
	"github.com/foomo/gotsrpc/v2/tests/nested-generics/private"
)

type Middle[T any] interface {
	private.Base[T]
	GetWrapped(ctx context.Context) common.Response[T]
}

type Keyed[K comparable, V any] interface {
	GetByKey(ctx context.Context, key K) V
}

type Service interface {
	Middle[common.Item]
	Keyed[string, int]
	GetName(ctx context.Context) string
}
