package db

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks") // создаём новый бакет
var db *bolt.DB                  // создаём переменную типа нашей базы данных

type Task struct { // создаём структуру нашей БД, состоящей из ключа и его значения
	Key   int
	Value string
}

func Init(dbPath string) error { // инициализируем БД и сразу проверяем на ошибку
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket) // создаётся новый бакет, если он не существует
		return err
	})
}

func CreateTask(task string) (int, error) { // создаём задачу
	var id int
	err := db.Update(func(tx *bolt.Tx) error { // проводим транзакцию записи
		b := tx.Bucket(taskBucket)  // извлекаем бакет по имени "taskBucket"
		id64, _ := b.NextSequence() // получаем следущий уникальный идентификатор
		id = int(id64)
		key := itob(id) // Присваиваем ключ из нашего полученого id
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor() // cursor - позволяет пройтись по ключам и их значениях
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket) // Достаём бакеты
		return b.Delete(itob(key)) // удаляем ключ
	})
}

func itob(v int) []byte { // формируем ключ для задачи
	b := make([]byte, 8)                     // создаём массив из 8 байт
	binary.BigEndian.PutUint64(b, uint64(v)) // создание двоичного файла.тип кодировки. засовываем нашу переменную в массив байтов преобразуя её в uint64
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
