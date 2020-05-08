package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"rest-server/model"
	"strings"
)

func Get(c *gin.Context) {
	path := strings.ToLower(c.Param("path"))
	data, err := model.Query(path)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    data.Data,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}

func Post(c *gin.Context) {
	path := strings.ToLower(c.Param("path"))
	var data model.Data
	data.Path = path
	body, _ := ioutil.ReadAll(c.Request.Body)
	data.Data = string(body)
	var err = model.Create(&data)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}

func Put(c *gin.Context) {
	//path := strings.ToLower(c.Param("path"))
	var data model.Data
	_ = c.BindJSON(&data)
	var err = model.Update(&data)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}

func Delete(c *gin.Context) {
	path := strings.ToLower(c.Param("path"))
	err := model.Delete(path)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
