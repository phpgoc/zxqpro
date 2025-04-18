package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateJson(g *gin.Context, req interface{}) bool {
	if err := g.ShouldBindJSON(req); err != nil {
		g.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return false
	}
	return true
}

func ValidateQuery(g *gin.Context, req interface{}) bool {
	if err := g.ShouldBindQuery(req); err != nil {
		g.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return false
	}
	return true
}
