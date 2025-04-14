package routes

import (
	"log"
	"net/http"

	"github.com/phpgoc/zxqpro/routes/middleware"
	"github.com/phpgoc/zxqpro/utils"

	"github.com/gin-gonic/gin"
)

// SSE 处理程序
func ServerSideEvent(c *gin.Context) {
	userId := middleware.GetUserIdFromAuthMiddleware(c)
	manager := c.MustGet("sseManager").(*utils.SSEManager)
	// 设置响应头
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// 保持连接打开
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		http.Error(c.Writer, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// 注册客户端
	client := manager.RegisterClient(userId)
	defer manager.UnregisterClient(userId)

	for {
		select {
		case <-c.Request.Context().Done():
			return
		case message, ok := <-client:
			if !ok {
				return
			}
			// 发送消息
			sseMessage := "data: " + message + "\n\n"
			_, err := c.Writer.Write([]byte(sseMessage))
			if err != nil {
				log.Println("Error writing message:", err)
				return
			}
			flusher.Flush()
		}
	}
}

func TestSendSelf(c *gin.Context) {
	userId := middleware.GetUserIdFromAuthMiddleware(c)
	manager := c.MustGet("sseManager").(*utils.SSEManager)
	manager.SendMessageToUser(userId, "test message")
	c.JSON(http.StatusOK, gin.H{"message": "Message sent"})
}
