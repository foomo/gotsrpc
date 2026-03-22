package model

type ServiceMethods []*Method

func (sm ServiceMethods) Len() int           { return len(sm) }
func (sm ServiceMethods) Swap(i, j int)      { sm[i], sm[j] = sm[j], sm[i] }
func (sm ServiceMethods) Less(i, j int) bool { return sm[i].Name < sm[j].Name }
