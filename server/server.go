package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/karolgorecki/gotodo/task"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandlers creates routing for all users requests
func RegisterHandlers() (router http.Handler) {
	rt := httprouter.New()
	rt.GET("/", showHomepage)
	rt.GET("/tasks", getTasks)
	rt.GET("/tasks/:id", getTask)
	rt.POST("/tasks", createTask)
	rt.PUT("/tasks/:id", updateTask)
	rt.DELETE("/tasks", deleteTasks)
	rt.DELETE("/tasks/:id", deleteTask)
	rt.ServeFiles("/dist/*filepath", http.Dir("./dist/"))
	fmt.Println("Running on http://localhost:8000")
	return rt
}

// showHomepage handles the GET for index route
func showHomepage(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var err error

	tpl, err := template.ParseFiles("./dist/index.html")
	if err != nil {
		log.Fatal(err)
	}

	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(http.StatusOK)
	// tpl.Execute(rw, nil)
	if err = tpl.Lookup("index.html").Execute(rw, nil); err != nil {
		log.Fatal(err)
	}
}

// getTasks handles the GET request for tasks
func getTasks(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var err error
	tsks := task.All()
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusOK)

	// If there are no tasks return empty json
	if tsks == nil {
		if _, err = rw.Write([]byte("[]")); err != nil {
			log.Fatal(err)
		}
		return
	}

	// Send to user all tasks
	if err = json.NewEncoder(rw).Encode(tsks); err != nil {
		log.Fatal(err)
	}
}

// getTask handles the GET request for specific task
func getTask(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var err error
	tsk := task.Get(p.ByName("id"))

	// Handle if the task doesn't exists
	if tsk == nil {
		rw.WriteHeader(http.StatusNotFound)
		if _, err = rw.Write([]byte("Task with given ID doesn't exists")); err != nil {
			log.Fatal(err)
		}
		return
	}

	// Send to user the found task
	rw.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(rw).Encode(tsk); err != nil {
		log.Fatal(err)
	}
}

// createTask get's data from user and returns created task
func createTask(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var err error
	tsk := task.Task{}

	// Get task data from user
	err = json.NewDecoder(req.Body).Decode(&tsk)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("There was a problem with creating new task:", err)
		return

	}

	// Create new task
	task.Create(&tsk)

	// Return to user created task
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(rw).Encode(tsk); err != nil {
		log.Fatal(err)
	}
}

// updateTask update given task
func updateTask(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var err error
	tsk := task.Task{}

	// Get task data from user
	err = json.NewDecoder(req.Body).Decode(&tsk)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("There was a problem with updating task:", err)
		return
	}

	// Update task
	t := task.Update(p.ByName("id"), &tsk)

	// If task doesn't exists
	if t == nil {
		rw.WriteHeader(http.StatusBadRequest)
		if _, err = rw.Write([]byte("Task was not updated")); err != nil {
			log.Fatal(err)
		}
		return
	}

	// Return to user created task
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(rw).Encode(tsk); err != nil {
		log.Fatal(err)
	}
}

// deleteTasks removes all tasks
func deleteTasks(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	task.DeleteAll()
	rw.WriteHeader(http.StatusNoContent)
}

// deleteTask removes task for given ID
func deleteTask(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	task.Delete(p.ByName("id"))
	rw.WriteHeader(http.StatusOK)
}
