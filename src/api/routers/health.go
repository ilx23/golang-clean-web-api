package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilx23/golang-clean-web-api/src/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
	r.GET("/test", handler.Test)
	r.GET("/test/:id", handler.TestId)
}