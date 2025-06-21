package models

import "time"

// Task represents a SAST scanning task.
type Task struct {
    ID        int64     `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
}

// Report represents the scan result associated with a task.
type Report struct {
    ID      int64  `json:"id"`
    TaskID  int64  `json:"task_id"`
    Content string `json:"content"`
}
