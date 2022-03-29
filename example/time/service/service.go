package service

import (
	"time"
)

type Service interface {
	Time(v time.Time) time.Time
	TimeStruct(v TimeStruct) TimeStruct
}
