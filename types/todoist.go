package types

import "time"

type Project struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Order        int    `json:"order"`  // read-only
	Indent       int    `json:"indent"` // read-only
	CommentCount int    `json:"comment_count"`
}

type Task struct {
	ID           int    `json:"id"`
	ProjectID    int    `json:"project_id"` // read-only
	Content      string `json:"content"`
	Completed    bool   `json:"completed"`
	LabelIDs     []int  `json:"label_ids"`
	Order        int    `json:"order"`  // read-only
	Indent       int    `json:"indent"` // read-only
	Priority     int    `json:"priority"`
	Due          Due    `json:"due"`
	URL          string `json:"url"`
	CommentCount int    `json:"comment_count"`
}

type TaskApply struct {
	Content         string    `json:"content"`
	ProjectID       int       `json:"project_id,omitempty"`
	LabelIDs        []int     `json:"label_ids,omitempty"`
	Order           int       `json:"order,omitempty"` // read-only
	Priority        int       `json:"priority,omitempty"`
	DueHumanDefined string    `json:"due_string,omitempty"`
	DueDate         string    `json:"due_date,omitempty"`
	DueDatetime     time.Time `json:"due_datetime,omitempty"`
}

type Due struct {
	HumanDefinedDate string    `json:"string"`
	Date             string    `json:"date"`
	DateTime         time.Time `json:"datetime,omitempty"`
}

type Label struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}
