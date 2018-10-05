package store

import (
	"github.com/boltdb/bolt"
	"log"
	"gophercises/07/task/config"
	"fmt"
)

type Store struct {
	DB *bolt.DB
}

func OpenDB() *Store {
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	s := &Store{DB: db}
	if err = s.createBuckets(); err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *Store) createBuckets() error {
	s.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(config.BucketTasks))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(config.BucketDone))
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}

func (s *Store) InsertTask(arg string) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.BucketTasks))

		id, _ := b.NextSequence()
		key := fmt.Sprintf("task%d", id)
		if err := b.Put([]byte(key), []byte(arg)); err != nil {
			return err
		}

		return nil
	})
}

type V struct {
	Key   string
	Value string
}

func (s *Store) LoadToDos() ([]V, error) {
	values := []V{}
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.BucketTasks))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			value := V{
				Key:string(k),
				Value:string(v),
			}
			values = append(values, value)
		}

		return nil
	})

	return values, err
}

func (s *Store) LoadCompletedTasks() ([]V, error) {
	values := []V{}
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.BucketDone))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			value := V{
				Key:string(k),
				Value:string(v),
			}
			values = append(values, value)
		}

		return nil
	})

	return values, err
}

func (s *Store) DeleteTask(key []byte) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(config.BucketTasks))
		return bucket.Delete(key)
	})
}

func (s *Store) GetTask(key []byte) ([]byte, error) {
	var value []byte
	err := s.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(config.BucketTasks))
		value = bucket.Get(key)
		return nil
	})

	return value, err
}

func (s *Store) MarkDone(key, value []byte) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		bDone := tx.Bucket([]byte(config.BucketDone))

		if err := bDone.Put(key, value); err != nil {
			return err
		}

		return nil
	})
}
