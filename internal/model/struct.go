package model

// Struct describes a named Go type that is rendered as a TypeScript type:
// a struct, or a named map, slice or array.
type Struct struct {
	// IsError reports whether the type implements error.
	IsError bool
	Package string
	Name    string
	Fields  []*Field
	// UnionFields are the members of a union struct.
	UnionFields []*Field
	// InlineFields are embedded fields whose fields are inlined.
	InlineFields []*Field
	Map          *Map
	Array        *Array
	// TypeParams are the names of the generic type parameters.
	TypeParams []string
	// Ref points to the type this one is declared as, e.g. `type A = pkg.B`
	Ref *StructType
}

// FullName returns the package-qualified name, for example "service.User".
func (s *Struct) FullName() string {
	fullName := s.Package + "." + s.Name
	if len(fullName) == 0 {
		fullName = s.Name
	}

	return fullName
}
