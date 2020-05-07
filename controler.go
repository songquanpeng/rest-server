package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-server/model"
)

func Get(c *gin.Context) {
	path := c.Param("path")
	data, _ := model.Get(path)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func Post(c *gin.Context) {

}

func Put(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
