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
func (m *SSEManager) RegisterClient(userId uint) chan SSEMessage {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.clients[userId]; !exists {
		m.clients[userId] = make(chan SSEMessage)
	}
	return m.clients[userId]
}

// UnregisterClient 注销 SSE 客户端
func (m *SSEManager) UnregisterClient(userId uint) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, exists := m.clients[userId]; exists {
		close(client)
		delete(m.clients, userId)
	}
}

// SendMessageToUser 向指定用户发送消息
func (m *SSEManager) SendMessageToUser(userId uint, sseMessage SSEMessage) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, exists := m.clients[userId]; exists {
		client <- sseMessage
	}
}

func (m *SSEManager) SendMessageToUsers(userIds []uint, sseMessage SSEMessage) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, userId := range userIds {
		if client, exists := m.clients[userId]; exists {
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
