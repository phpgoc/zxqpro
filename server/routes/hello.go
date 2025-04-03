package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/

// HelloWorld  godoc
// @Summary hello example
// @Schemes
// @Description do hello
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} HelloWorld
// @Router /api/hello_world [get]
func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "hello world")
}
