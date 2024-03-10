package main

import (
	"list/cmd"
	"list/db"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir() // Хранение бд в домашнем каталоге пользователя
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
