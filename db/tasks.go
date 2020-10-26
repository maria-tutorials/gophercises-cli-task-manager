package db

import (
	"encoding/binary"
	"errors"
	"log"
	"time"

	"../consts"

	"github.com/boltdb/bolt"
)

var tasksBucket = consts.TASKS_BUCKET
var completedBucket = consts.COMPLETED_BUCKET
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

// Init connects to bold and creates a new bucket if needed
func Init(path string) error {
	var err error
	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(completedBucket)
		if err != nil {
			return err
		}
		return nil
	})
}

// CreateTask adds a new task to the database
func CreateTask(task string, bucket []byte) (int, error) {
	id := 0
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

// AllTasks displays all tasks
func AllTasks(bucket []byte) ([]Task, error) {
	t := []Task{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			t = append(t, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return t, nil
}

// GetTask shows a single task
func GetTask(key int, bucket []byte) (Task, error) {
	t := Task{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		v := b.Get(itob(key))

		if v == nil {
			return errors.New("Nothing there")
		}

		t = Task{
			Key:   key,
			Value: string(v),
		}
		return nil
	})

	if err != nil {
		return t, err
	}
	return t, nil
}

// CompleteTask "marks" tasks as done
func CompleteTask(key int) error {

	t, err := GetTask(key, tasksBucket)
	if err != nil {
		return err
	}
	_, err = CreateTask(t.Value, completedBucket)
	if err != nil {
		return err
	}

	err = DeleteTask(key, tasksBucket)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask deletes a task
func DeleteTask(key int, bucket []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(64))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
