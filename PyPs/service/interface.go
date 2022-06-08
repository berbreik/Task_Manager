package service

import "PyPs/models"
// DataStoreInterface is a interface that contains all the methods that the store needs to implement
type DatastoreInterface interface {
	PostAdmin(new *models.Projects) error
	PostTask(new *models.Task) error
	UpdateTask(new *models.Task) error
	GetAll() ([]*models.Task, error)
	GetTaskById(id int64) (*models.Task,error)
	UpdateStatusDev(new *models.Task) (*models.Task, error)
}