package server

import (
	"context"

	"github.com/foomo/gotsrpc/v2/tests/common"
)

type Handler struct{}

func (h *Handler) GetValue(_ context.Context) common.Item {
	return common.Item{ID: "1", Name: "test"}
}

func (h *Handler) GetWrapped(_ context.Context) common.Response[common.Item] {
	return common.Response[common.Item]{
		Data: common.Item{ID: "1", Name: "test"},
	}
}

func (h *Handler) GetByKey(_ context.Context, key string) int {
	return len(key)
}

func (h *Handler) GetName(_ context.Context) string {
	return "service"
}
