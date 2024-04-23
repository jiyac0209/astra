package domain

import "time"

type Astra struct {
	UUID      string    `json:"uuid" omitempty`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp" omitempty`
}
