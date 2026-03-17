package server

import (
	"context"
)

type Service interface {
	GetStringResponse(ctx context.Context) Response[string]
	GetItemResponse(ctx context.Context) Response[Item]
	SetItemResponse(ctx context.Context, req Response[Item]) bool
	GetPair(ctx context.Context) Pair[string, int]
	GetPagedItems(ctx context.Context, page int) PagedResponse[Item]
	GetNestedGeneric(ctx context.Context) PagedResponse[Pair[string, Item]]
	GetResult(ctx context.Context) Result[Item]
	GetContainer(ctx context.Context) Container[string, Item]
}
