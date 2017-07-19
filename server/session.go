package server

import (
	"github.com/gorilla/sessions"
)

type SessionStore struct {
	key   []byte
	Store *sessions.CookieStore
}

var sessionStore *SessionStore

func GetSessionStore() *SessionStore {
	if sessionStore == nil {
		sessionStore = &SessionStore{key: []byte("super-secret-key")}
		sessionStore.Store = sessions.NewCookieStore(sessionStore.key)
	}
	return sessionStore
}
