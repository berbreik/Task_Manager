package handler

import "PyPs/models"
// ServiceInterface is a interface that contains all the methods that the service needs to implement
type ServiceInterface interface {
	PostAdminService(new *models.Projects) error
	PostTaskService(new *models.Task) error
	UpdateTaskService(new *models.Task) error
	GetAllService() ([]*models.Task, error)
	GetTaskByIdService(id int64) (*models.Task, error)
	UpdateStatusDevService(new *models.Task) (*models.Task, error)
}