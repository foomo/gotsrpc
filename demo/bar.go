package demo

type Bar interface {
	Hello(number int64) int
	Repeat(one, two string) (three, four bool)
}
