package server

import (
	"context"
)

type Handler struct{}

func (h *Handler) GetValue(_ context.Context) Item {
	return Item{ID: "1", Name: "test"}
}

func (h *Handler) GetWrapped(_ context.Context) Response[Item] {
	return Response[Item]{
		Data: Item{ID: "1", Name: "test"},
	}
}

func (h *Handler) GetByKey(_ context.Context, key string) int {
	return len(key)
}

func (h *Handler) GetName(_ context.Context) string {
	return "service"
}
