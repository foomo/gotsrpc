package server

import (
	"context"
)

type Middle interface {
	Base
	GetMiddleName(ctx context.Context) string
}
