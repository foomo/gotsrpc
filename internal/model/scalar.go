package model

// Scalar is a named Go type whose underlying type is a scalar, such as
// `type Status string`.
type Scalar struct {
	Name    string
	Package string
	Type    ScalarType
}

// FullName returns the package-qualified name, for example "time.Time".
func (st *Scalar) FullName() string {
	fullName := st.Package + "." + st.Name
	if len(fullName) == 0 {
		fullName = st.Name
	}

	return fullName
}
