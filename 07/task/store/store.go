package store

import (
	"github.com/boltdb/bolt"
	"log"
	"encoding/json"
	"encoding/binary"
	"gophercises/07/task/config"
)

type Store struct {
	DB     *bolt.DB
	Bucket *bolt.Bucket
}

func OpenDB() *Store {
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	s := &Store{DB: db}
	if err = s.createBucket(); err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *Store) createBucket() error {
	s.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(config.Bucket))
		if err != nil {
			return err
		}
		s.Bucket = b

		return nil
	})

	return nil
}

func InsertTask(db *bolt.DB, args []string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.Bucket))

		id, _ := b.NextSequence()

		body, err := json.Marshal(args)
		if err != nil {
			return err
		}

		if err = b.Put(itob(id), body); err != nil {
			return err
		}

		return nil
	})
}

func (s *Store) Load() ([][]byte, error) {
	values := [][]byte{}
	err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(config.Bucket))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			values = append(values, v)
		}

		return nil
	})

	return values, err
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
