package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

var (
	port  = flag.Int("port", 3000, "specify the server listening port.")
	Token = flag.String("token", "token", "specify the private token.")
)

func SetRouter(router *gin.Engine) {
	router.Use(cors.Default())
	router.GET("/*path", Get)
	router.POST("/*path", Post)
	router.PUT("/*path", Put)
	router.DELETE("/*path", Delete)
}

func main() {
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	flag.Parse()
	db, err := InitDB()
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
		log.Println(err)
	}
}
