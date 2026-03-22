package model

type StructType struct {
	Name    string
	Package string
}

func (st *StructType) FullName() string {
	fullName := st.Package + "." + st.Name
	if len(fullName) == 0 {
		fullName = st.Name
	}
	return fullName
}
