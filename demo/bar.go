package demo

type (
	Inner struct {
		One string `json:"one"`
	}
	OuterNested struct {
		Inner Inner  `json:"inner"`
		Two   string `json:"two"`
	}
	OuterInline struct {
		Inner `json:",inline"`
		Two   string `json:"two"`
	}
)

type Bar interface {
	Hello(number int64) int
	Repeat(one, two string) (three, four bool)
	Inheritance(inner Inner, nested OuterNested, inline OuterInline) (Inner, OuterNested, OuterInline)
}
