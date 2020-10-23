package main

import (
	"log"
	"path/filepath"

	"./cmd"
	"./db"
	"github.com/mitchellh/go-homedir"
)

const dbName string = "tasks.db"

func main() {
	home, _ := homedir.Dir()
	path := filepath.Join(home, dbName)
	err := db.Init(path)
	if err != nil {
		log.Fatal(err)
	}

	must(cmd.RootCmd.Execute())
}

//must is an error wrapper
func must(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
