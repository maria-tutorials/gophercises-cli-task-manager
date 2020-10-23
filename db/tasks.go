package db

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var tasksBucket = []byte("tasks")
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
		return err
	})
}

// CreateTask adds a new task to the database
func CreateTask(task string) (int, error) {
	id := 0
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
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
func AllTasks() ([]Task, error) {
	t := []Task{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
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

// GetTask shows a single task TODO:
func GetTask() {

}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(64))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
