package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
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
		log.Fatal("failed to init database")
	}
	defer db.Close()
	server := gin.Default()
	SetRouter(server)
	var realPort = os.Getenv("PORT")
	if realPort == "" {
		realPort = strconv.Itoa(*port)
	}
	err = server.Run(":" + realPort)
	if err != nil {
		log.Print(err)
	}
}
