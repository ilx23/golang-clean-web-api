package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilx23/golang-clean-web-api/src/api/handlers"
)


func Test(r *gin.RouterGroup) {
	handler := handlers.NewTestHandler()

	r.GET("/", handler.Test)
	r.GET("/users", handler.Users)
	r.GET("/user/:id", handler.UserById)
	r.GET("/user/get-user-by-username/:username", handler.UserByUsername)
	r.POST("/add-user", handler.AddUser)

	r.POST("/binder/header1", handler.HeaderBinder1)
	r.POST("/binder/header2", handler.HeaderBinder2)


	r.POST("/binder/query1", handler.QueryBinder1)
	r.POST("/binder/query2", handler.QueryBinder2)

	r.POST("/binder/url/:id/:name", handler.UrlBinder)

	r.POST("/binder/body", handler.BodyBinder)
	r.POST("/binder/form", handler.FormBinder)
}