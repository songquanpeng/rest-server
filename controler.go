package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(c *gin.Context) {
	path := strings.ToLower(c.Param("path"))
	data, err := Query(path)
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
	var data Data
	data.Path = path
	body, _ := ioutil.ReadAll(c.Request.Body)
	data.Data = string(body)
	var err = Create(&data)
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
	var data Data
	data.Path = strings.ToLower(c.Param("path"))
	body, _ := ioutil.ReadAll(c.Request.Body)
	data.Data = string(body)
	var err = Update(&data)
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
	err := Remove(path)
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
