package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type header struct {
	UserId string
	Browser string
}

type personData struct {
	FirstName string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName string `json:"last_name" binding:"required,alpha,min=5,max=20"`
	PhoneNumber string `json:"phone_number" binding:"required,number,min=11,max=11"`
}

type TestHandler struct {

}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, "Test Page")
}

func (t *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, "Users List")
	return
}

func (t *TestHandler) UserById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("User with %s id is online now !", id))
	return
}


func (t *TestHandler) UserByUsername(c *gin.Context) {
	username := c.Params.ByName("username")
	c.JSON(http.StatusOK, fmt.Sprintf("%s is now online ", username))
	return
}


func (t *TestHandler) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, "User has been successfuly added")
	return
}

func (t *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("userId")
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	})
	return
}
func (t *TestHandler) HeaderBinder2(c *gin.Context) {
	h := header{}
	c.BindHeader(&h)
	

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"header": h,
	})
	return
}

func (t *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"id": id,
		"name": name,
	})
	return
}

func (t *TestHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"id": ids,
		"name": name,
	})
	return
}

func (t *TestHandler) UrlBinder(c *gin.Context) {
	id := c.Params.ByName("id")
	name := c.Params.ByName("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"id": id,
		"name": name,
	})
	return
}

func (t *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validationError": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"person": p,
	})
	return
}

func (t *TestHandler) FormBinder(c *gin.Context) {
	p := personData{}
	c.Bind(&p)

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"person": p,
	})
	return
}