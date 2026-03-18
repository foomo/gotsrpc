package server

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Response[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error,omitempty"`
}
