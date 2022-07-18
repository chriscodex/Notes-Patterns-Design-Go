package main

import (
	"fmt"
	"sync"
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
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDataBaseInstance()
		}()
	}
	wg.Wait()
}
