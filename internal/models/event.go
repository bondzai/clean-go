package models

import "time"

type Event struct {
	ID        string    `json:"id"`
	Message   *string   `json:"message,omitempty"`
	TimeStamp time.Time `json:"time_stamp"`
}
