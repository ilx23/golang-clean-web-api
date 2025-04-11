package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {

}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Working !")
	return
}

func (h *HealthHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, "Now, you are in the test page")
	return
}
func (h *HealthHandler) TestId(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("Working post by id: %s", id))
	return
}