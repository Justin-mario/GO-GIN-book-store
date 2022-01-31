package main

import (
	"github.com/Justin-mario/GO-GIN-book-store/model"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	FindAll() []model.Bookstore
	Save(ctx gin.Context)
}
