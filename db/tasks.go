package db

import "github.com/boltdb/bolt"

var taskBucket = []byte("tasks") // создаём новый бакет
var db *bolt.DB                  // создаём переменную типа нашей базы данных

type Task struct { // создаём структуру нашей БД, состоящей из ключа и его значения
	Key   int
	Value string
}

func Init(dbPath string) error { // инициализируем БД
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
	})
	return 0, nil
}
