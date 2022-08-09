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

func (db *Database) Set(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.items[key] = value
}

func (db *Database) Get(key string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, found := db.items[key]
	return value, found
}
