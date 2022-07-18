package main

import (
	"fmt"
	"time"
)

type database struct {
}

func (database) createSingleConnection() {
	fmt.Println("Creating singleton for Database")
	time.Sleep(2 * time.Second)
	fmt.Println("Creating done")
}

var db *database

func getDataBaseInstance() *database {
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &database{}
		db.createSingleConnection()
	} else {
		fmt.Println("DB already created")
	}

}
