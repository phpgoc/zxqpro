package utils

import (
	"sync"
)

// 用户 SSE 连接管理
type SSEManager struct {
	clients map[uint]chan string
	mu      sync.Mutex
}

// 初始化 SSE 管理器
func NewSSEManager() *SSEManager {
	return &SSEManager{
		clients: make(map[uint]chan string),
	}
}

// 注册 SSE 客户端
func (m *SSEManager) RegisterClient(userId uint) chan string {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.clients[userId]; !exists {
		m.clients[userId] = make(chan string)
	}
	return m.clients[userId]
}

// 注销 SSE 客户端
func (m *SSEManager) UnregisterClient(userId uint) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, exists := m.clients[userId]; exists {
		close(client)
		delete(m.clients, userId)
	}
}

// 向指定用户发送消息
func (m *SSEManager) SendMessageToUser(userId uint, message string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, exists := m.clients[userId]; exists {
		client <- message
	}
}
