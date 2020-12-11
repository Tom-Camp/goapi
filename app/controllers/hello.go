package controllers

import "github.com/gin-gonic/gin"

type HelloController struct{}

func (i *HelloController) Default(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello world, climate change is real"})
}
