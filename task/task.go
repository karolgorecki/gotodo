package task

import "github.com/satori/go.uuid"

// Task struct
type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

// taskStore is an interface which allows to store tasks
type taskStore interface {
	All() []*Task
	Get(id string) (foundTask *Task)
	Create(tsk *Task) (createdTask *Task)
	Update(id string, tsk *Task) (updatedTask *Task)
	DeleteAll()
	Delete(id string)
}

// db is an taskStore implementation
var db taskStore

// RegisterDB allows to use different task store implementation
func RegisterDB(database taskStore) {
	db = database
}

// NewID generates an unique id used for task ID
func NewID() (id string) {
	uuid := uuid.NewV4()
	return uuid.String()
}

// All gets all tasks from store
func All() (tasks []*Task) {
	tasks = db.All()
	return tasks
}

// Create makes new task
func Create(tsk *Task) (createdTask *Task) {
	tsk.ID = NewID()
	db.Create(tsk)
	return tsk
}

// Get returns the task for given ID
func Get(id string) (foundTask *Task) {
	foundTask = db.Get(id)
	return foundTask
}

// Update updated given task
func Update(id string, tsk *Task) (updatedTask *Task) {
	updatedTask = db.Update(id, tsk)
	return updatedTask
}

// DeleteAll removes all tasks from db
func DeleteAll() {
	db.DeleteAll()
}

// Delete removes one task for given ID
func Delete(id string) {
	db.Delete(id)
}
