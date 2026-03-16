package server

import "context"

type Handler struct{}

func (h *Handler) VariantA(_ context.Context, i1 Base) (r1 Base)                       { return i1 }
func (h *Handler) VariantB(_ context.Context, i1 BCustomType) (r1 BCustomType)         { return i1 }
func (h *Handler) VariantC(_ context.Context, i1 BCustomTypes) (r1 BCustomTypes)       { return i1 }
func (h *Handler) VariantD(_ context.Context, i1 BCustomTypesMap) (r1 BCustomTypesMap) { return i1 }

func (h *Handler) VariantE(_ context.Context, i1 *Base) (r1 *Base) { return i1 }

func (h *Handler) VariantF(_ context.Context, i1 []*Base) (r1 []*Base) { return i1 }

func (h *Handler) VariantG(_ context.Context, i1 map[string]*Base) (r1 map[string]*Base) { return i1 }

func (h *Handler) VariantH(_ context.Context, i1 Base, i2 *Base, i3 []*Base, i4 map[string]Base) (r1 Base, r2 *Base, r3 []*Base, r4 map[string]Base) {
	return i1, i2, i3, i4
}
