package main

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"github.com/iFaceless/rest"
	"encoding/json"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Checked     bool      `json:"checked"`
	isDeleted   bool
}

var (
	tasks = []*Task{
		{1, "Fix bugs", time.Now(), false, false},
		{2, "Add unit tests", time.Now(), false, false},
		{3, "Release stable version", time.Now(), false, false},
	}
)

func main() {
	r := rest.NewRouter()
	r.MountHandler("/tasks", &TasksHandler{})
	r.MountHandler("/tasks/{task_id:(\\d+)}", &TaskHandler{})
	rest.Run(r, 8000)
}

type TasksHandler struct {
	rest.BaseHandler
}

func (hd *TasksHandler) Get() {
	visibleTasks := make([]*Task, 0)
	for _, t := range tasks {
		if t.isDeleted == false {
			visibleTasks = append(visibleTasks, t)
		}
	}
	hd.RenderJSON(visibleTasks)
}

func (hd *TasksHandler) Post() {
	var task Task
	json.NewDecoder(hd.R.Body).Decode(&task)
	task.ID = len(tasks) + 1
	tasks = append(tasks, &task)
	hd.RenderJSON(map[string]bool{"success": true})
}

type TaskHandler struct {
	rest.BaseHandler
}

func (hd *TaskHandler) Get() {
	taskID := cast.ToInt(chi.URLParam(hd.R, "task_id"))
	task := lookupTask(taskID)
	if task == nil || task.isDeleted {
		hd.RenderError(rest.HTTPResourceNotFoundError)
		return
	}
	hd.RenderJSON(task)
}

func (hd *TaskHandler) Delete() {
	taskID := cast.ToInt(chi.URLParam(hd.R, "task_id"))
	task := lookupTask(taskID)
	if task == nil || task.isDeleted {
		hd.RenderError(rest.HTTPResourceNotFoundError)
		return
	}
	task.isDeleted = true
	hd.RenderJSON(map[string]bool{"success": true})
}

func lookupTask(id int) *Task {
	for _, t := range tasks {
		if t.ID == id {
			return t
		}
	}
	return nil
}
