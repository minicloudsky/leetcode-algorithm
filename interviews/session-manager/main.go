package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Session 设计一个会话管理（SessionManager）系统，要求：
// 1. 可以添加和获取用户会话（Session）
// 2. 当会话不再被外部使用时自动从管理器中移除
// 3. 避免内存泄漏
// 4. 支持并发访问
type Session struct {
	SessionId string
}

type SessionManageImpl struct {
	SessionMap  map[string]*Session
	SessionList []string
	Capacity    int64
	lock        sync.Mutex
}

type SessionManager interface {
	// 添加或更新会话，返回会话ID
	AddSession(s *Session) string

	// 获取会话，如果不存在返回nil
	GetSession(sessionID string) *Session

	// 返回当前活跃会话数量
	ActiveSessions() int
	Delete(sessionId string)
}

func (s *SessionManageImpl) AddSession(session *Session) string {
	s.lock.Lock()
	if session.SessionId == "" {
		source := rand.NewSource(10000)
		session.SessionId = fmt.Sprintf("%v", source.Int63())
	}
	s.SessionMap[session.SessionId] = session
	s.lock.Unlock()
	return session.SessionId
}

func (s *SessionManageImpl) GetSession(sessionID string) *Session {
	session, ok := s.SessionMap[sessionID]
	if ok && session != nil {
		return session
	}
	return nil
}

func (s *SessionManageImpl) ActiveSessions() int {
	return 0
}
