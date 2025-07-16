package models

// imports
import (
	"time"
)

// task item
type Task struct {
	ID            int         `json:"id"`                     				    // unique identifier of task
	Title         string      `json:"title"`                  				    // title of task
	Description   string      `json:"description"`            				    // description of task
	DueDate       time.Time   `json:"due_date" time_format:"2006-01-02T15:04:05Z"`  	    // due date of task (ISO 8601 format)
	Status        string      `json:"status" binding:"oneof=pending in_progress completed"`     // status of task
}
