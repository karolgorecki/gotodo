package task_test

import (
	"testing"

	"github.com/karolgorecki/todo/task"
	"github.com/karolgorecki/todo/testdb"
)

func TestRegisterDB(t *testing.T) {
	// nothing to test here
	task.RegisterDB(&testdb.TestDB{})
}

func TestNewID(t *testing.T) {
	id1 := task.NewID()
	id2 := task.NewID()
	if id1 == id2 {
		t.Error("The generated ID's are the same")
	}
}

func TestGetTasks(t *testing.T) {
	tsks := task.All()
	if len(tsks) != 3 {
		t.Errorf("Number of returned tasks should be 3. Got %d", len(tsks))
	}

	switch tsks[0].ID {
	case "1", "2", "3":
	default:
		t.Errorf("The ID of task should be 1,2 or 3. Got %q", tsks[0].ID)
	}
}

func TestCreateTask(t *testing.T) {
	tsk := task.Task{
		Name: "New task",
	}
	task.Create(&tsk)
	if tsk.ID == "" {
		t.Errorf("Task should have an unique ID. Got %q", tsk.ID)
	}
}

func TestGetTask(t *testing.T) {
	tsk := task.Get("1")
	if tsk.Name != "Buy milk" {
		t.Error("Get returned wrong task")
	}

	tsk = task.Get("2")
	if tsk.Name != "Write todo app" {
		t.Error("Get returned wrong task")
	}

	tsk = task.Get("1000")
	if tsk != nil {
		t.Error("Returned not existing task")
	}
}

func TestUpdateTask(t *testing.T) {
	tsk := task.Task{
		Name: "Task updated",
		Done: true,
	}
	updated := task.Update("1", &tsk)
	if updated.Done != true || updated.Name != "Task updated" {
		t.Error("The task was not updated correctly")
	}

	tsk = task.Task{
		Name: "I'm not exitsing",
		Done: true,
	}
	updated = task.Update("1000", &tsk)
	if updated != nil {
		t.Error("Updated non existing task")
	}
}

func TestDeleteTask(t *testing.T) {
	task.Delete("2")
	if len(testdb.Tsks) != 2 {
		t.Error("Task was no removed")
	}
	for _, tsk := range testdb.Tsks {
		if tsk.ID == "2" {
			t.Error("Removed wrong task")
		}
	}
}

func TestDeleteAllTasks(t *testing.T) {
	task.DeleteAll()
	if len(testdb.Tsks) != 0 {
		t.Error("Tasks were not deleted")
	}
}
