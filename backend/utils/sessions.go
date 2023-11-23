package utils

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)
var sessionKey = struct{}{}
// Session represents a user session
type Session struct {
	ID        string
	Values    map[string]interface{}
	ExpireAt  time.Time
}

var sessions = make(map[string]*Session)
var mu sync.Mutex

// NewSession creates a new session
func NewSession() *Session {
	mu.Lock()
	defer mu.Unlock()

	id := uuid.New().String()
	session := &Session{
		ID:       id,
		Values:   make(map[string]interface{}),
		ExpireAt: time.Now().Add(24 * time.Hour), // Например, на сутки
	}

	sessions[id] = session

	// Очистка старых сессий
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			mu.Lock()
			for id, session := range sessions {
				if time.Now().After(session.ExpireAt) {
					delete(sessions, id)
				}
			}
			mu.Unlock()
		}
	}()

	return session
}

// GetSession retrieves a session by ID
func GetSession(id string) *Session {
	mu.Lock()
	defer mu.Unlock()
	return sessions[id]
}

// SetSession stores a session
func SetSession(session *Session) {
	mu.Lock()
	defer mu.Unlock()
	sessions[session.ID] = session
}

// DeleteSession removes a session
func DeleteSession(id string) {
	mu.Lock()
	defer mu.Unlock()
	delete(sessions, id)
}

// CleanUpSessions removes expired sessions
func CleanUpSessions() {
	mu.Lock()
	defer mu.Unlock()
	for id, session := range sessions {
		if time.Now().After(session.ExpireAt) {
			delete(sessions, id)
		}
	}
}

// CleanUpSessionsScheduler periodically calls CleanUpSessions
func CleanUpSessionsScheduler() {
	for {
		time.Sleep(1 * time.Hour)
		CleanUpSessions()
	}
}

// Middleware для проверки сессии
func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err != nil {
			// Нет куки с сессией
			next.ServeHTTP(w, r)
			return
		}

		session := GetSession(sessionID.Value)
		if session == nil {
			// Сессия не найдена
			next.ServeHTTP(w, r)
			return
		}

		// Продлеваем срок действия сессии
		session.ExpireAt = time.Now().Add(24 * time.Hour) // Например, на сутки
		SetSession(session)

		// Присоединяем сессию к запросу
		r = r.WithContext(SessionContext(r.Context(), session))

		next.ServeHTTP(w, r)
	})
}

// SessionContext возвращает копию контекста с добавленной сессией
func SessionContext(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}

// GetSessionFromContext получает сессию из контекста
func GetSessionFromContext(ctx context.Context) *Session {
	session, _ := ctx.Value(sessionKey).(*Session)
	return session
}