package manager

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type SessionManager struct {
	SessionData map[string]*SessionData
	Lock        sync.RWMutex
}

const (
	SessionCookieName  = "session_id"
	SessionContextName = "session_data"
)

var (
	Manager *SessionManager
)

func InitSessionManager() {
	Manager = &SessionManager{
		SessionData: make(map[string]*SessionData, 1024),
	}
}

func (m *SessionManager) GetSessionData(key string) (data *SessionData, err error) {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	data, ok := m.SessionData[key]
	if !ok {
		err = fmt.Errorf("invalid session id")
	}
	return
}

func (m *SessionManager) CreateSession() (sd *SessionData) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	uuid := uuid.NewV4().String()
	sd = NewSessionData(uuid)
	m.SessionData[uuid] = sd
	return
}
