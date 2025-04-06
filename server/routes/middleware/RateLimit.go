package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var globalIpsMapArray = make([]*sync.Map, 0)

// CleanAllMap 计算所有 sync.Map 中嵌套 sync.Map 里 int64 值的总和并清空最外层 sync.Map
func CleanAllMap() int {
	var sum int
	for _, outerMap := range globalIpsMapArray {
		outerMap.Range(func(_, value interface{}) bool {
			// 检查值是否为 *sync.Map 类型
			if nestedMap, ok := value.(*sync.Map); ok {
				nestedMap.Range(func(_, nestedValue interface{}) bool {
					// 检查嵌套值是否为 int64 类型
					if v, ok := nestedValue.(int); ok {
						sum += v
					}
					return true
				})
			}
			return true
		})
		// 清空当前外层 sync.Map
		outerMap.Range(func(key, _ interface{}) bool {
			outerMap.Delete(key)
			return true
		})
	}
	return sum
}

type ipRateLimiter struct {
	ips sync.Map
	r   int
}

// newIPRateLimiter 创建一个新的 IP 访问频率限制器实例
func newIPRateLimiter(reqPerHour int) *ipRateLimiter {
	rl := &ipRateLimiter{
		r: reqPerHour,
	}
	globalIpsMapArray = append(globalIpsMapArray, &rl.ips)
	go rl.cleanupOldHours()
	return rl
}

// getCurrentHour 获取当前小时数
func getCurrentHour() int64 {
	return time.Now().Unix() / 3600
}

// incrementCount 增加指定 IP 在当前小时的访问次数
func (i *ipRateLimiter) incrementCount(ip string) bool {
	currentHour := getCurrentHour()
	hourMap, _ := i.ips.LoadOrStore(currentHour, &sync.Map{})
	count, loaded := hourMap.(*sync.Map).LoadOrStore(ip, 1)
	if loaded {
		newCount := count.(int) + 1
		if newCount > i.r {
			return false
		}
		hourMap.(*sync.Map).Store(ip, newCount)
	}
	return true
}

// cleanupOldHours 每小时清理一次旧的小时数据
func (i *ipRateLimiter) cleanupOldHours() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		currentHour := getCurrentHour()
		i.ips.Range(func(key, value interface{}) bool {
			hour := key.(int64)
			if hour < currentHour {
				i.ips.Delete(hour)
			}
			return true
		})
	}
}

// RateLimitMiddleware 创建一个 Gin 中间件来实现 IP 访问频率限制
func RateLimit(reqPerHour int) gin.HandlerFunc {
	limiter := newIPRateLimiter(reqPerHour)
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		if !limiter.incrementCount(clientIP) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code":    http.StatusTooManyRequests,
				"message": "Too many requests, please try again later.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
