package utils

import (
	"fmt"
	"sync"
)

// SSEManager 用户 SSE 连接管理
type SSEManager struct {
	clients map[uint]map[string]chan SSEMessage // 用户
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
		clients: make(map[uint]map[string]chan SSEMessage),
	}
}

// RegisterClient 注册 SSE 客户端
func (m *SSEManager) RegisterClient(userID uint, uuid string) chan SSEMessage {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.clients[userID]; !exists {
		m.clients[userID] = make(map[string]chan SSEMessage)
	}
	m.clients[userID][uuid] = make(chan SSEMessage)
	return m.clients[userID][uuid]
}

// UnregisterClient 注销 SSE 客户端
func (m *SSEManager) UnregisterClient(userID uint, uuid string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, exists := m.clients[userID]; exists {
		close(client[uuid])
		if len(client) == 1 {
			delete(m.clients, userID)
		} else {
			delete(m.clients[userID], uuid)
		}
	}
	// 视乎会比较频繁的出现，当tab页不活跃，浏览器会主动断开连接，这是现代浏览器的优化，正常
	LogInfo(fmt.Sprintf("UnregisterClient: userID=%d, uuid=%s", userID, uuid))
}

func (m *SSEManager) UnregisterUser(userID uint) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, exists := m.clients[userID]; exists {
		for uuid, c := range client {
			close(c)
			delete(client, uuid)
		}
		delete(m.clients, userID)
	}
}

// SendMessageToUser 向指定用户发送消息
func (m *SSEManager) SendMessageToUser(userID uint, sseMessage SSEMessage) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, exists := m.clients[userID]; exists {
		for _, c := range client {
			c <- sseMessage
		}
	}
}

func (m *SSEManager) SendMessageToUsers(userIDs []uint, sseMessage SSEMessage) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, userID := range userIDs {
		if client, exists := m.clients[userID]; exists {
			for _, c := range client {
				c <- sseMessage
			}
		}
	}
}

// SendMessageToAllUsers 向所有用户发送消息
func (m *SSEManager) SendMessageToAllUsers(sseMessage SSEMessage) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, client := range m.clients {
		for _, c := range client {
			c <- sseMessage
		}
	}
}
