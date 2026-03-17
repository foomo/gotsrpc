package server

import (
	"context"
)

type Middle interface {
	Base
	GetMiddleName(ctx context.Context) string
}

type Extra interface {
	GetAge(ctx context.Context) int
}

type ExtendedService interface {
	Middle
	Extra
	GetLastName(ctx context.Context) string
	GetPerson(ctx context.Context) Person
}
