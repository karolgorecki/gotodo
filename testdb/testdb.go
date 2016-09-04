package testdb

import "github.com/karolgorecki/todo/task"

func init() {
	task.RegisterDB(&TestDB{})
}

// Tsks holds tasks for testing purpose
var Tsks []*task.Task

// TestDB is a store for testing
type TestDB struct{}

// All mocks the get all tasks
func (db *TestDB) All() (tasks []*task.Task) {
	Tsks = []*task.Task{
		{"1", "Buy milk", false},
		{"2", "Write todo app", true},
		{"3", "Go to gym", false},
	}
	return Tsks
}

// Create fakes store action & returns saved task
func (db *TestDB) Create(tsk *task.Task) (task *task.Task) {
	return tsk
}

// Get returns a task for given ID
func (db *TestDB) Get(id string) (foundTask *task.Task) {
	for _, tsk := range Tsks {
		if tsk.ID == id {
			return tsk
		}
	}
	return nil
}

// Update updates the given task
func (db *TestDB) Update(id string, tsk *task.Task) (updateTask *task.Task) {
	for _, t := range Tsks {
		if id == t.ID {
			t = tsk
			t.ID = id
			return t
		}
	}
	return nil
}

// DeleteAll removes all tasks
func (db *TestDB) DeleteAll() {
	Tsks = nil
}

// Delete removes all tasks
func (db *TestDB) Delete(id string) {
	for k, v := range Tsks {
		if v.ID == id {
			Tsks = append(Tsks[:k], Tsks[k+1:]...)
		}
	}
}
