package model

type Scalar struct {
	Name    string
	Package string
	Type    ScalarType
}

func (st *Scalar) FullName() string {
	fullName := st.Package + "." + st.Name
	if len(fullName) == 0 {
		fullName = st.Name
	}

	return fullName
}
