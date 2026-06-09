package server

import (
	"context"
)

type Handler struct{}

func (h *Handler) GetFirstName(_ context.Context) string {
	return "John"
}

func (h *Handler) GetMiddleName(_ context.Context) string {
	return "Michael"
}

func (h *Handler) GetAge(_ context.Context) int {
	return 30
}

func (h *Handler) GetLastName(_ context.Context) string {
	return "Doe"
}

func (h *Handler) GetPerson(_ context.Context) Person {
	return Person{
		FirstName:  "John",
		MiddleName: "Michael",
		LastName:   "Doe",
		Age:        30,
	}
}
