package main

import (
	"github.com/Justin-mario/GO-GIN-book-store/model"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	model.ConnectDatabase()
	err := server.Run()
	if err != nil {
		return
	}

}
