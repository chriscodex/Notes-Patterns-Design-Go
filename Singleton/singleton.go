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
	fmt.Println("Database created")
}

var db *database
var mutex sync.Mutex

func getDataBaseInstance() *database {
	mutex.Lock()
	defer mutex.Unlock()
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
