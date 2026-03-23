package model

type ServiceList []*Service

func (sl ServiceList) Len() int           { return len(sl) }
func (sl ServiceList) Swap(i, j int)      { sl[i], sl[j] = sl[j], sl[i] }
func (sl ServiceList) Less(i, j int) bool { return sl[i].Name < sl[j].Name }
