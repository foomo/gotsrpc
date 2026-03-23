package common

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Response[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error,omitempty"`
}

type Simple struct {
	Bool    bool    `json:"bool"`
	Int     int     `json:"int"`
	Int64   int64   `json:"int64"`
	Float64 float64 `json:"float64"`
	String  string  `json:"string"`
}

type Nested struct {
	Name  string `json:"name"`
	Child Simple `json:"child"`
}

type Other struct {
	Label string `json:"label"`
}
