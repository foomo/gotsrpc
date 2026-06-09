package server

type Pair[A, B any] struct {
	First  A `json:"first"`
	Second B `json:"second"`
}

type PagedResponse[T any] struct {
	Items []T `json:"items"`
	Total int `json:"total"`
}

type Result[T any] struct {
	Value *T     `json:"value,omitempty"`
	Error string `json:"error,omitempty"`
}

type Container[K comparable, V any] struct {
	Data    map[K]V `json:"data"`
	Default V       `json:"default"`
}
