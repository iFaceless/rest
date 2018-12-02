RESTful API
============
               _
 _ __ ___  ___| |_
| '__/ _ \/ __| __|
| | |  __/\__ \ |_
|_|  \___||___/\__|

[rest](https://github.com/ifaceless/rest) is a simple wrapper for [go-chi/chi](https://github.com/go-chi/chi), which aims at writing RESTful API with struct-based handler (not pure function-based handler) in an elegant way.

Following example demonstrates how to register routes for a struct-based handler previously：

```golang
r := chi.NewRouter()

hd = NewTasksHandler()

r.Get("/tasks/{task_id:(\\d+)}", hd.Get)
r.Post("/tasks/{task_id:(\\d+)}", hd.Post)
r.Delete("/tasks/{task_id:(\\d+)", hd.Delete}
```

However, with this simple wrapper, you can register you routes much simpler:

```golang
r := rest.NewRouter()
r.MountHandler("/tasks/{task_id:(\\d+)}", &TasksHandler{})
```

# Install

```
go get -u https://github.com/ifaceless/rest
```

# Quick Start

Full demonstration can be found [here](./example/main.go)~

```golang
func main() {
	r := rest.NewRouter()
	r.MountHandler("/tasks", &TasksHandler{})
	r.MountHandler("/tasks/{task_id:(\\d+)}", &TaskHandler{})
	rest.Run(r, 8000)
}

type TasksHandler struct {
	rest.BaseHandler
}

// Get matches the HTTP GET method, you can add your custom
// logic here to override the same method defined in the `BaseHandler`
func (hd *TasksHandler) Get() {
	hd.RenderJSON(map[string]string{"message": "Get Tasks"})
}

// Post matches the HTTP Post method, you can add your custom
// logic here to override the same method defined in the `BaseHandler`
func (hd *TasksHandler) Post() {
	hd.RenderJSON(map[string]string{"message": "Create Task"})
}

type TaskHandler struct {
	rest.BaseHandler
}

func (hd *TaskHandler) Get() {
	hd.RenderJSON(map[string]string{"message": "GET Task"})
}

func (hd *TaskHandler) Delete() {
	hd.RenderJSON(map[string]string{"message": "Delete Task"})
}
```

# Core APIs

- `Prepare()`
- `Finish()`
- `Get()`
- `Post()`
- `Patch()`
- `Put()`
- `Delete()`
- `RenderJSON()`
- `RenderError()`
- `Offset()`
- `Limit()`
- `URLParam()`
- `QueryArgument()`
- `IntQueryArgument()`
- `Int64QueryArgument()`
- `R`: `*http.Request`
- `W`: `http.ResponseWriter`

# Help & Dev & Bug Report

Please feel free to report any bugs or ask for new features~