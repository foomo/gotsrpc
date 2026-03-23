package model

type Service struct {
	Name        string
	Methods     ServiceMethods
	Endpoint    string
	IsInterface bool
}
