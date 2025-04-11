package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/ilx23/golang-clean-web-api/src/api/middlewares"
	"github.com/ilx23/golang-clean-web-api/src/api/routers"
	"github.com/ilx23/golang-clean-web-api/src/api/validations"
	"github.com/ilx23/golang-clean-web-api/src/config"
)


func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleWare()*/, middlewares.LimitByRequest())

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("number", validations.IranianMobilePhoneNumberValidator, true)
	}

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{

		health := v1.Group("/health")
		test_router := v1.Group("/test")
		
		routers.Health(health)
		routers.Test(test_router)

	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}