package manager

import (
	"fmt"
	"sync"
)

type SessionData struct {
	Id   string
	Data map[string]interface{}
	Lock sync.RWMutex
}

func NewSessionData(id string) *SessionData {
	return &SessionData{
		Id:   id,
		Data: make(map[string]interface{}, 16),
	}
}

func (s *SessionData) Get(key string) (val interface{}, err error) {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	val, ok := s.Data[key]
	if !ok {
		err = fmt.Errorf("invalid key")
	}
	return
}

func (s *SessionData) Set(key string, val interface{}) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Data[key] = val
}

func (s *SessionData) Del(key string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	delete(s.Data, key)
}
