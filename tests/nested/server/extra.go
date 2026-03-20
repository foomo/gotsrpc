package server

import (
	"context"
)

type Extra interface {
	GetAge(ctx context.Context) int
}
