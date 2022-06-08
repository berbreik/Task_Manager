// package models models of the application
package models

// Projects describes the structure of the projects table
type Projects struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Tasks Task   `json:"tasks"`
}

// task describes the structure of the tasks table and its relationship with the projects table
type Task struct {
	Id          int64  `json:"id"`
	ProjectId   int64  `json:"project_id"`
	Description string `json:"description"`
	Assignee    string `json:"assignee"`
	Status      string `json:"status"`
}

// User describes the structure of the users table
type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	UserType role   `json:"role"`
}
// role describes the structure of the user_type table
type role string

const (
	ADMIN     role = "admin"
	DEVELOPER role = "developer"
	MANAGER   role = "manager"
)
