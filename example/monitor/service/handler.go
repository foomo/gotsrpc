package service

import (
	"fmt"
)

type Handler struct{}

func (h *Handler) Hello(v string) string {
	fmt.Println(v)
	return v
}
