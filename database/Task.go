package database

import "time"


type Task struct {
    ID          int64
    Title       string
    Description string
    Status      string
    CreatedAt   time.Time
}