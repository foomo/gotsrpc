package server

import (
	"context"
)

type Service interface {
	// Scalar aliases
	StatusValue(ctx context.Context, v Status) Status
	CategoryValue(ctx context.Context, v Category) Category
	PriorityValue(ctx context.Context, v Priority) Priority
	RatingValue(ctx context.Context, v Rating) Rating

	// Typed collections
	TagsValue(ctx context.Context, v Tags) Tags
	EntriesValue(ctx context.Context, v Entries) Entries
	RegistryValue(ctx context.Context, v Registry) Registry
	IndexValue(ctx context.Context, v Index) Index
	LabelMapValue(ctx context.Context, v LabelMap) LabelMap

	// Structs with enums
	EntryValue(ctx context.Context, v Entry) Entry
	DetailValue(ctx context.Context, v Detail) Detail
	DataRecordValue(ctx context.Context, v DataRecord) DataRecord

	// Complex nesting
	MapOfEntries(ctx context.Context, v map[Category][]Entry) map[Category][]Entry

	// Nil optionals
	DataRecordNil(ctx context.Context, v DataRecord) DataRecord
}
