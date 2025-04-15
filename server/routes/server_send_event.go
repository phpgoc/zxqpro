package routes

import (
	"encoding/json"
	"net/http"

	"github.com/phpgoc/zxqpro/interfaces"
	"github.com/phpgoc/zxqpro/pro_types"

	"github.com/phpgoc/zxqpro/routes/middleware"
	"github.com/phpgoc/zxqpro/utils"

	"github.com/gin-gonic/gin"
)

func joinMessage(sseMessage utils.SSEMessage) (res []byte) {
	jsonData, _ := json.Marshal(sseMessage)
	res = []byte("data: " + string(jsonData) + "\n\n")
	return
}

// SSE 处理程序
func ServerSideEvent(c *gin.Context) {
	manager := c.MustGet("sseManager").(*utils.SSEManager)
	// 设置响应头
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))

	// 保持连接打开
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		http.Error(c.Writer, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	cookie, err := c.Request.Cookie(utils.CookieName)
	if err != nil {
		_, _ = c.Writer.Write(joinMessage(utils.SSEMessage{
			Code:    401,
			Message: err.Error(),
		}))
		flusher.Flush()
		return
	}
	cookieValue := cookie.Value
	var cookieData pro_types.Cookie
	has := interfaces.Cache.Get(cookieValue, &cookieData)
	if !has {
		_, _ = c.Writer.Write(joinMessage(utils.SSEMessage{
			Code:    401,
			Message: "Unauthorized",
		}))
		flusher.Flush()
		return
	}

	userId := cookieData.ID

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

			_, err = c.Writer.Write(joinMessage(message))
			if err != nil {
				utils.LogError("Error writing message:" + err.Error())
				return
			}
			flusher.Flush()
		}
	}
}

func TestSendSelf(c *gin.Context) {
	userId := middleware.GetUserIdFromAuthMiddleware(c)
	sseManager := c.MustGet("sseManager").(*utils.SSEManager)
	sseManager.SendMessageToUser(userId, utils.SSEMessage{
		Code:    0,
		Message: "test message",
	})
	c.JSON(http.StatusOK, gin.H{"message": "Message sent"})
}
