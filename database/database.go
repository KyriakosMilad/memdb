package database

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Database struct {
	items map[string]string
	mu    sync.RWMutex
}

const databaseDiskFile = "memdb.json"

func InitDatabase() *Database {
	fmt.Println("initialising memdb")
	f, err := os.Open(databaseDiskFile)
	if err != nil {
		return &Database{items: map[string]string{}}
	}
	items := map[string]string{}
	if err := json.NewDecoder(f).Decode(&items); err != nil {
		fmt.Println("could not decode, creating fresh database...", err.Error())
		return &Database{items: map[string]string{}}
	}
	return &Database{items: items}
}
