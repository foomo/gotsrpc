package server

import (
	"context"
)

type Base interface {
	GetFirstName(ctx context.Context) string
}
