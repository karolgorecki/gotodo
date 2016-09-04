package boltstore

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/karolgorecki/todo/task"
)

const (
	dbName     = "tasks.db"
	bucketName = "Tasks"
)

// BoltStore is a task store implementaiton based on Bolt
type BoltStore struct {
	bolt *bolt.DB
}

func init() {
	var err error

	// Create new bold db
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create new bucket bucketName
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Tell task package to use Bolt store
	task.RegisterDB(&BoltStore{db})
}

// All returns all tasks in bolt db
func (bs *BoltStore) All() (tasks []*task.Task) {
	var tsks []*task.Task

	err := bs.bolt.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tsk := task.Task{}
			err := json.Unmarshal(v, &tsk)
			if err != nil {
				log.Fatal(err)
			}

			tsks = append(tsks, &tsk)

		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return tsks
}

// Create saves the task to bolt db
func (bs *BoltStore) Create(tsk *task.Task) (createdTask *task.Task) {
	var err error

	encoded, err := json.Marshal(tsk)

	if err != nil {
		log.Fatal(err)
	}

	err = bs.bolt.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(tsk.ID), []byte(encoded))
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
	return tsk
}

// Get returns task for given ID
func (bs *BoltStore) Get(id string) (foundTask *task.Task) {
	var err error

	err = bs.bolt.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(id))

		if v == nil {
			foundTask = nil
			return nil
		}
		err := json.Unmarshal(v, &foundTask)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return foundTask
}

// Update updates given task in bolt db
func (bs *BoltStore) Update(id string, tsk *task.Task) (updateTask *task.Task) {
	foundTask := bs.Get(id)

	if foundTask == nil {
		return nil
	}

	tsk.ID = id

	// If user forgets name
	if tsk.Name == "" {
		tsk.Name = foundTask.Name
	}

	encoded, err := json.Marshal(tsk)
	if err != nil {
		log.Fatal(err)
	}

	err = bs.bolt.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(tsk.ID), []byte(encoded))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	return tsk
}

// DeleteAll removes all tasks from bolt db
func (bs *BoltStore) DeleteAll() {
	err := bs.bolt.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucketName))
		if err != nil {
			log.Fatal(err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}

// Delete removes one task from bolt db
func (bs *BoltStore) Delete(id string) {
	err := bs.bolt.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Delete([]byte(id))

		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}
