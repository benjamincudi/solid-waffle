package main

import (
	"log"
	"solid-waffle/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.AttachRouteHandlers(r)
	if err := r.Run(":80"); err != nil {
		log.Fatal(err)
	}
}
