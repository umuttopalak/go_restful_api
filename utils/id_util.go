package utils

import "sync"

var idCounter int
var idMutex sync.Mutex

func GetNextID() int {
	idMutex.Lock()
	defer idMutex.Unlock()
	idCounter++
	return idCounter
}
