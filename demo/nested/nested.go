package nested

type CustomTypeNested string

const (
	CustomTypeNestedOne   CustomTypeNested = "one"
	CustomTypeNestedTwo   CustomTypeNested = "two"
	CustomTypeNestedThree CustomTypeNested = "three"
)

type Nested struct {
	Name              string
	Any               Any
	AnyMap            map[string]Any
	AnyList           []Any
	SuperNestedString struct {
		Ha int64
	}
	SuperNestedPtr *struct {
		Bla string
	}
}

type Any interface{}
type Amount int
type True bool

const ItIsTrue True = true

type JustAnotherStingType string
