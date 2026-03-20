package server

import (
	"context"

	"github.com/foomo/gotsrpc/v2/tests/common"
)

type Service interface {
	GetStringResponse(ctx context.Context) common.Response[string]
	GetItemResponse(ctx context.Context) common.Response[common.Item]
	SetItemResponse(ctx context.Context, req common.Response[common.Item]) bool
	GetPair(ctx context.Context) Pair[string, int]
	GetPagedItems(ctx context.Context, page int) PagedResponse[common.Item]
	GetNestedGeneric(ctx context.Context) PagedResponse[Pair[string, common.Item]]
	GetResult(ctx context.Context) Result[common.Item]
	GetContainer(ctx context.Context) Container[string, common.Item]
}
