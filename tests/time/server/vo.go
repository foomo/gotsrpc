package server

import (
	"time"
)

type TimeStruct struct {
	Time        time.Time  `json:"time"`
	TimePtr     *time.Time `json:"timePtr"`
	TimePtrOmit *time.Time `json:"timePtrOmit,omitempty"`
}
