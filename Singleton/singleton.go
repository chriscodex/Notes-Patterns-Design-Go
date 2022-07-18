package main

import "fmt"

type database struct {
}

func (database) createSingleConnection() {

}

var db *database

func getDataBaseInstance() *database {
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &database{}
	}
}
