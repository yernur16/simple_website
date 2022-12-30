package router

import (
	"fmt"
	"net/http"
)

type TasksHandler struct{}

func (th *TasksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path != "/" {
			th.listTaskHandler(w, r)
			return
		}
		th.readTaskHandler(w, r)
	case http.MethodPost:
		th.createTaskHandler(w, r)
	case http.MethodPut:
		th.updateTaskHandler(w, r)
	case http.MethodDelete:
		th.deleteTaskHandler(w, r)
	default:
		msg := http.StatusText(http.StatusNotFound)
		ErrJSON(w, &httpErr{err: fmt.Errorf(msg), msg: msg}, http.StatusNotFound)
	}
}

func (th *TasksHandler) listTaskHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: List Tasks
}
func (th *TasksHandler) readTaskHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Read Task by ID
	taskId, err := extractTaskId(r)
	if err != nil {
		ErrJSON(w, &httpErr{err: err, msg: "not found"}, http.StatusNotFound)
		return
	}
	task, err := someDB.GetTaskById(taskId)
}
func (th *TasksHandler) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Delete Task by ID
}
func (th *TasksHandler) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Create Task by ID
}
func (th *TasksHandler) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Update Task by ID
}
