package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
