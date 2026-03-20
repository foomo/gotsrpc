package server

import (
	"context"

	"github.com/foomo/gotsrpc/v2/tests/common"
)

type Handler struct{}

func (h *Handler) GetStringResponse(_ context.Context) common.Response[string] {
	return common.Response[string]{Data: "hello"}
}

func (h *Handler) GetItemResponse(_ context.Context) common.Response[common.Item] {
	return common.Response[common.Item]{Data: common.Item{ID: "1", Name: "test"}}
}

func (h *Handler) SetItemResponse(_ context.Context, req common.Response[common.Item]) bool {
	return req.Data.ID != ""
}

func (h *Handler) GetPair(_ context.Context) Pair[string, int] {
	return Pair[string, int]{First: "hello", Second: 42}
}

func (h *Handler) GetPagedItems(_ context.Context, page int) PagedResponse[common.Item] {
	return PagedResponse[common.Item]{Items: []common.Item{{ID: "1", Name: "item1"}}, Total: 1}
}

func (h *Handler) GetNestedGeneric(_ context.Context) PagedResponse[Pair[string, common.Item]] {
	return PagedResponse[Pair[string, common.Item]]{
		Items: []Pair[string, common.Item]{{First: "key", Second: common.Item{ID: "1", Name: "nested"}}},
		Total: 1,
	}
}

func (h *Handler) GetResult(_ context.Context) Result[common.Item] {
	item := common.Item{ID: "1", Name: "result"}
	return Result[common.Item]{Value: &item}
}

func (h *Handler) GetContainer(_ context.Context) Container[string, common.Item] {
	return Container[string, common.Item]{
		Data:    map[string]common.Item{"key": {ID: "1", Name: "contained"}},
		Default: common.Item{ID: "0", Name: "default"},
	}
}
