package server

import (
	"context"
)

type ExtendedService interface {
	Middle
	Extra
	GetLastName(ctx context.Context) string
	GetPerson(ctx context.Context) Person
}
