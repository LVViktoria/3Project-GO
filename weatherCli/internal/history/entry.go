package history

import (
	"time"
)

type Entry struct {
	City        string    `json:"city"`
	Temperature string    `json:"temperature"`
	Time        time.Time `json:"time"`
}
