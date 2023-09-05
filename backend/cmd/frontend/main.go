package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("starting server")
	r := gin.Default()

	r.Static("/", "web")

	r.Run(":8080")
}
