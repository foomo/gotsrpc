package model

// ServiceMethods is a list of methods. It implements [sort.Interface] and
// sorts by name.
type ServiceMethods []*Method

// Len implements [sort.Interface].
func (sm ServiceMethods) Len() int { return len(sm) }

// Swap implements [sort.Interface].
func (sm ServiceMethods) Swap(i, j int) { sm[i], sm[j] = sm[j], sm[i] }

// Less implements [sort.Interface].
func (sm ServiceMethods) Less(i, j int) bool { return sm[i].Name < sm[j].Name }
