package model

// ServiceList is a list of services. It implements [sort.Interface] and
// sorts by name.
type ServiceList []*Service

// Len implements [sort.Interface].
func (sl ServiceList) Len() int { return len(sl) }

// Swap implements [sort.Interface].
func (sl ServiceList) Swap(i, j int) { sl[i], sl[j] = sl[j], sl[i] }

// Less implements [sort.Interface].
func (sl ServiceList) Less(i, j int) bool { return sl[i].Name < sl[j].Name }
