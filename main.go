package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"rest-server/model"
	"strconv"
)

var (
	port  = flag.Int("port", 3000, "specify the server listening port.")
	Token = flag.String("token", "token", "specify the private token.")
)

func SetRouter(router *gin.Engine) {
	router.GET("/*path", Get)
	router.POST("/*path", Post)
	router.PUT("/*path", Put)
	router.DELETE("/*path", Delete)
}

func main() {
	flag.Parse()
	db, err := model.InitDB()
	if err != nil {
		_ = fmt.Errorf("failed to init database")
	}
	defer db.Close()
	server := gin.Default()
	SetRouter(server)
	var realPort = os.Getenv("PORT")
	if realPort == "" {
		realPort = strconv.Itoa(*port)
	}
	_ = server.Run(":" + realPort)
}
