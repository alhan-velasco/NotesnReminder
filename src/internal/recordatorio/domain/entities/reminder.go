package entities

import "time"

type Reminder struct {
    ID          int32
    Title       string
    Description string
    DateTime    time.Time
}
