package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jqs7/demo/app"
	"github.com/jqs7/demo/models"
)

func Hello(c *gin.Context) {
	data := models.Hello{Hello: app.Config().GetString("hello")}
	c.JSON(http.StatusOK, data)
}
