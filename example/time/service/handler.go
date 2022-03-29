package service

import (
	"fmt"
	"time"
)

type Handler struct{}

func (h *Handler) Time(v time.Time) time.Time {
	fmt.Println(v.String())
	return v
}

func (h *Handler) TimeStruct(v TimeStruct) TimeStruct {
	fmt.Println(v.Time.String())
	return v
}
