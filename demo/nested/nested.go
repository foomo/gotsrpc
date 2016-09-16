package nested

type Nested struct {
	Name              string
	SuperNestedString struct {
		Ha int64
	}
	SuperNestedPtr *struct {
		Bla string
	}
}

type Amount int
type True bool

const ItIsTrue True = true
