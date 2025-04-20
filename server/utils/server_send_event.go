package utils

import (
	"sync"
)

// SSEManager 用户 SSE 连接管理
type SSEManager struct {
	clients map[uint]chan SSEMessage
	mu      sync.Mutex
}

type SSEMessage struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Link    *string `json:"link"`
}

// NewSSEManager 初始化 SSE 管理器
func NewSSEManager() *SSEManager {
	return &SSEManager{
		clients: make(map[uint]chan SSEMessage),
	}
}

// RegisterClient 注册 SSE 客户端
func (m *SSEManager) RegisterClient(userID uint) chan SSEMessage {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.clients[userID]; !exists {
		m.clients[userID] = make(chan SSEMessage)
	}
	return m.clients[userID]
}

// UnregisterClient 注销 SSE 客户端
func (m *SSEManager) UnregisterClient(userID uint) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, exists := m.clients[userID]; exists {
		close(client)
		delete(m.clients, userID)
	}
}

// SendMessageToUser 向指定用户发送消息
func (m *SSEManager) SendMessageToUser(userID uint, sseMessage SSEMessage) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, exists := m.clients[userID]; exists {
		client <- sseMessage
	}
}

func (m *SSEManager) SendMessageToUsers(userIDs []uint, sseMessage SSEMessage) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, userID := range userIDs {
		if client, exists := m.clients[userID]; exists {
			client <- sseMessage
		}
	}
}

// SendMessageToAllUsers 向所有用户发送消息
func (m *SSEManager) SendMessageToAllUsers(sseMessage SSEMessage) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, client := range m.clients {
		client <- sseMessage
	}
}
