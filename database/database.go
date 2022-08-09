package database

import (
	"sync"
)

type Database struct {
	items map[string]string
	mu    sync.RWMutex
}
