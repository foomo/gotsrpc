package server

import (
	"context"
)

type Handler struct{}

func (h *Handler) StatusValue(_ context.Context, v Status) Status       { return v }
func (h *Handler) CategoryValue(_ context.Context, v Category) Category { return v }
func (h *Handler) PriorityValue(_ context.Context, v Priority) Priority { return v }
func (h *Handler) RatingValue(_ context.Context, v Rating) Rating       { return v }

func (h *Handler) TagsValue(_ context.Context, v Tags) Tags             { return v }
func (h *Handler) EntriesValue(_ context.Context, v Entries) Entries    { return v }
func (h *Handler) RegistryValue(_ context.Context, v Registry) Registry { return v }
func (h *Handler) IndexValue(_ context.Context, v Index) Index          { return v }
func (h *Handler) LabelMapValue(_ context.Context, v LabelMap) LabelMap { return v }

func (h *Handler) EntryValue(_ context.Context, v Entry) Entry                { return v }
func (h *Handler) DetailValue(_ context.Context, v Detail) Detail             { return v }
func (h *Handler) DataRecordValue(_ context.Context, v DataRecord) DataRecord { return v }

func (h *Handler) MapOfEntries(_ context.Context, v map[Category][]Entry) map[Category][]Entry {
	return v
}

func (h *Handler) DataRecordNil(_ context.Context, v DataRecord) DataRecord { return v }
