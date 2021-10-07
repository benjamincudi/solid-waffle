package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	if err := r.Run(":80"); err != nil {
		log.Fatal(err)
	}
}
