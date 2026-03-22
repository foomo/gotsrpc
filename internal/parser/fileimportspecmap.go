package parser

type fileImportSpecMap map[string]importSpec

func (m fileImportSpecMap) getPackagePath(packageName string) string {
	is, ok := m[packageName]
	if ok {
		packageName = is.path
	}

	return packageName
}
