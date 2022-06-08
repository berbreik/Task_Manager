package store

import (
	"PyPs/models"
	"database/sql"
	"errors"
)
// datastore is a struct that contains a sql.DB
type datastore struct {
	db *sql.DB
}
// NewDataStore is a function that returns a new datastore
func NewDataStore(db *sql.DB) *datastore {
	return &datastore{db: db}
}
// PostAdmin is a function that inserts a new project into the database
func (s *datastore) PostAdmin(new *models.Projects) error {
	_, err := s.db.Exec("INSERT INTO projects (name) VALUES (?)", new.Name)
	if err != nil {
		return errors.New("error initializing new project")
	}
	return nil
}
// PostTask is a function that inserts a new task into the database
func (s *datastore) PostTask(new *models.Task) error {
	_, err := s.db.Exec("INSERT INTO tasks (description,assignee,status,project_id) VALUES (?,?,?,?)", new.Description, new.Assignee, new.Status, new.ProjectId)
	if err != nil {
		return errors.New("error initializing new task")
	}
	return nil
}
// UpdateTask is a function that updates a task in the database
func (s *datastore) UpdateTask(new *models.Task) error {
	feilds, values := FormQuery(new)

	query := "UPDATE tasks SET " + feilds + " WHERE id=?"
	_, err := s.db.Exec(query, values, new.Id)
	if err != nil {
		return errors.New("error updating task")
	}
	return nil
}
// GetAll is a function that returns all tasks from the database
func (s *datastore) GetAll() ([]*models.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, errors.New("error getting all tasks")
	}
	defer rows.Close()
	var tasks []*models.Task
	for rows.Next() {
		var task *models.Task
		err := rows.Scan(&task.Id, &task.Description, &task.Assignee, &task.Status, &task.ProjectId)
		if err != nil {
			return nil, errors.New("error getting all tasks")
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
// GetTaskById is a function that returns a task from the database
func (s *datastore) GetTaskById(id int64) (*models.Task,error) {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE id=?", id)
	if err != nil {
		return nil,errors.New("error getting task by id")
	}
	defer rows.Close()
	var tasks models.Task
	for rows.Next() {
		err := rows.Scan(&tasks.Id, &tasks.Description, &tasks.Assignee, &tasks.Status, &tasks.ProjectId)
		if err != nil {
			return nil, errors.New("error getting task by id")
		}
	}
	return &tasks,nil
}
// UpdateStatusDev is a function that updates the status of a task by developer
func (s *datastore) UpdateStatusDev(new *models.Task) (*models.Task, error) {
	feilds, values := FormQuery(new)

	query := "UPDATE tasks SET " + feilds + " WHERE id=?"
	_, err := s.db.Exec(query, values, new.Id)
	if err != nil {
		return nil, errors.New("error updating task")
	}
	return new, nil
}
