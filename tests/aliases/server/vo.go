package server

// Status is a string enum representing entity status.
type Status string

const (
	StatusActive  Status = "active"
	StatusPending Status = "pending"
	StatusClosed  Status = "closed"
)

// Category is a string enum for categorization.
type Category string

const (
	CategoryA Category = "a"
	CategoryB Category = "b"
)

// Priority is a numeric enum for priority levels.
type Priority int

const (
	PriorityLow    Priority = 1
	PriorityMedium Priority = 2
	PriorityHigh   Priority = 3
)

// Rating is a float64 alias for ratings.
type Rating float64

// Tags is a typed alias for a string slice.
type Tags []string

// Entries is a typed alias for a slice of Entry pointers.
type Entries []*Entry

// Registry is a typed alias for a map of entries by string key.
type Registry map[string]Entry

// Index is a typed alias for a map of entry slices by category.
type Index map[Category][]Entry

// LabelMap is a typed alias for a string-to-string map.
type LabelMap map[string]string

// Entry is a struct with enum fields.
type Entry struct {
	ID       string   `json:"id"`
	Status   Status   `json:"status"`
	Priority Priority `json:"priority"`
	Rating   Rating   `json:"rating"`
	Tags     Tags     `json:"tags,omitempty"`
}

// Detail is a struct with mixed field types.
type Detail struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Entry       Entry    `json:"entry"`
	Labels      LabelMap `json:"labels,omitempty"`
}

// DataRecord is a deeply nested struct with optional fields.
type DataRecord struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	Status     Status     `json:"status"`
	Amount     *Amount    `json:"amount,omitempty"`
	Items      Entries    `json:"items,omitempty"`
	Metadata   *Metadata  `json:"metadata,omitempty"`
	Categories []Category `json:"categories,omitempty"`
}

// Amount represents a monetary value with currency.
type Amount struct {
	Value    int64  `json:"value"`
	Currency string `json:"currency"`
}

// Metadata holds record metadata.
type Metadata struct {
	CreatedBy string  `json:"createdBy"`
	Note      *string `json:"note,omitempty"`
}
