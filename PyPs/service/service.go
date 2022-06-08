package service

import (
	"PyPs/models"
	"errors"
)
// TaskService is a struct that contains the datastore interface
type TaskService struct {
	TaskStore DatastoreInterface
}
// NewService is a function that returns a new task service and initializes the datastore
func NewService(TaskStore DatastoreInterface) *TaskService {
	return &TaskService{TaskStore: TaskStore}
}
// PostAdminService is a function that validates the project and then posts it to the datastore
func (s *TaskService) PostAdminService(new *models.Projects) error {
	
	if new.Name == ""{
		return errors.New("invalid name")
	}

	err := s.TaskStore.PostAdmin(new)
	if err != nil {
		return errors.New("error initializing new project")
	}
	return nil
}
// PostTaskService is a function that validates the task and then posts it to the datastore
func (s *TaskService) PostTaskService(new *models.Task)error{

	if new.Description == ""{
		return errors.New("invalid description")
	}
	if new.Assignee == ""{
		return errors.New("invalid assignee")
	}
	if new.Status != "start" && new.Status != "end" && new.Status != "complete" && new.Status != "incomplete"{
		return errors.New("invalid status")
	}
	if new.ProjectId == 0{
		return errors.New("invalid project id")
	}

	err:= s.TaskStore.PostTask(new)
	if err != nil {
		return errors.New("error initializing new task")
	}
	return nil
}
// UpdateTaskService is a function that validates the task and then updates it to the datastore
func (s *TaskService) UpdateTaskService(new *models.Task) error {
	err := s.TaskStore.UpdateTask(new)
	if err != nil {
		return errors.New("error updating task")
	}
	return nil
}
// GetAllService is a function that returns all tasks from the datastore
func (s *TaskService) GetAllService() ([]*models.Task, error) {
	tasks, err := s.TaskStore.GetAll()
	if err != nil {
		return nil, errors.New("error getting all tasks")
	}
	return tasks, nil
}
// GetTaskByIdService is a function that returns a task from the datastore
func (s *TaskService) GetTaskByIdService(id int64) (*models.Task,error) {

	if id == 0{
		return nil,errors.New("invalid id")
	}
	task,err := s.TaskStore.GetTaskById(id)
	if err != nil {
		return nil,errors.New("error getting task by id")
	}
	return task,nil
}
// UpdateStatusService is a function that validates the status and then updates it to the datastore
func (s *TaskService) UpdateStatusDevService(new *models.Task) (*models.Task, error) {

	if new.Status != "start" && new.Status != "end" && new.Status != "complete" && new.Status != "incomplete" && new.Status != ""{
		return nil,errors.New("invalid status")
	}

	task,err := s.TaskStore.UpdateStatusDev(new)
	if err != nil {
		return nil,errors.New("error updating status")
	}
	return task,nil
}