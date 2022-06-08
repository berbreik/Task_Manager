package handler

import (
	"PyPs/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TaskHandler is a struct that contains a TaskService
type TaskHandler struct {
	TaskService ServiceInterface
}

// NewHandler is a function that returns a new TaskHandler and initializes it with a TaskService
func NewHandler(TaskService ServiceInterface) *TaskHandler {
	return &TaskHandler{TaskService: TaskService}
}

// PostAdminHandler is a function that handles the POST request to /Admin/Projects
func (h *TaskHandler) PostAdminHandler(w http.ResponseWriter, r *http.Request) {
	var new models.Projects
	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.TaskService.PostAdminService(&new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// PostTaskHandler is a function that handles the POST request to /Manager/Task
func (h *TaskHandler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var new models.Task
	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.TaskService.PostTaskService(&new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateTaskHandler is a function that handles the PUT request to /Manager/Task/{id}
func (h *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var new models.Task
	err := json.NewDecoder(r.Body).Decode(&new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.TaskService.UpdateTaskService(&new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetAllHandler is a function that handles the GET request to /Manager/Task
func (h *TaskHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.TaskService.GetAllService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByIdHandler is a function that handles the GET request to /Manager/Task/{id} and returns a Task
func (h *TaskHandler) GetTaskByIdHandler(w http.ResponseWriter, r *http.Request) {
	id1, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := int64(id1)
	task, err := h.TaskService.GetTaskByIdService(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
