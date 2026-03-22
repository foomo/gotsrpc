package model

type Struct struct {
	IsError      bool
	Package      string
	Name         string
	Fields       []*Field
	UnionFields  []*Field
	InlineFields []*Field
	Map          *Map
	Array        *Array
}

func (s *Struct) FullName() string {
	fullName := s.Package + "." + s.Name
	if len(fullName) == 0 {
		fullName = s.Name
	}
	return fullName
}
