package model

// StructType references a named type by package and name.
type StructType struct {
	Name    string
	Package string
}

// FullName returns the package-qualified name, for example "service.User".
func (st *StructType) FullName() string {
	fullName := st.Package + "." + st.Name
	if len(fullName) == 0 {
		fullName = st.Name
	}

	return fullName
}
