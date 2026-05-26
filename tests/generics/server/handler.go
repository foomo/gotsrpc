package server

import (
	"context"

	"github.com/foomo/gotsrpc/v2/tests/generics/private"
)

type Handler struct{}

func (h *Handler) GetStringResponse(_ context.Context) Response[string] {
	return Response[string]{Data: "hello"}
}

func (h *Handler) GetItemResponse(_ context.Context) Response[Item] {
	return Response[Item]{Data: Item{ID: "1", Name: "test"}}
}

func (h *Handler) SetItemResponse(_ context.Context, req Response[Item]) bool {
	return req.Data.ID != ""
}

func (h *Handler) GetPair(_ context.Context) Pair[string, int] {
	return Pair[string, int]{First: "hello", Second: 42}
}

func (h *Handler) GetPagedItems(_ context.Context, page int) PagedResponse[Item] {
	return PagedResponse[Item]{Items: []Item{{ID: "1", Name: "item1"}}, Total: 1}
}

func (h *Handler) GetNestedGeneric(_ context.Context) PagedResponse[Pair[string, Item]] {
	return PagedResponse[Pair[string, Item]]{
		Items: []Pair[string, Item]{{First: "key", Second: Item{ID: "1", Name: "nested"}}},
		Total: 1,
	}
}

func (h *Handler) GetResult(_ context.Context) Result[Item] {
	item := Item{ID: "1", Name: "result"}
	return Result[Item]{Value: &item}
}

func (h *Handler) GetContainer(_ context.Context) Container[string, Item] {
	return Container[string, Item]{
		Data:    map[string]Item{"key": {ID: "1", Name: "contained"}},
		Default: Item{ID: "0", Name: "default"},
	}
}

func (h *Handler) SetEnvelope(_ context.Context, env *private.Envelope[Item]) string {
	return env.ID + ":" + env.Payload.Name
}

func (h *Handler) GetEnvelope(_ context.Context, id string) *private.Envelope[Item] {
	return &private.Envelope[Item]{
		ID:      id,
		Payload: Item{ID: "1", Name: "boxed"},
	}
}

func (h *Handler) RoundtripForeignEnvelope(_ context.Context, env *private.Envelope[private.Tag]) *private.Envelope[private.Tag] {
	return &private.Envelope[private.Tag]{
		ID: "echo:" + env.ID,
		Payload: private.Tag{
			Name:  env.Payload.Name,
			Value: env.Payload.Value,
		},
	}
}
