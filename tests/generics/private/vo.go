package private

type Envelope[Payload any] struct {
	ID      string  `json:"id"`
	Payload Payload `json:"payload"`
}

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
