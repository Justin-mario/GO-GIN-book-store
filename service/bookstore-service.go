package service

import "github.com/Justin-mario/GO-GIN-book-store/model"

type BookService interface {
	Save(model.Bookstore) model.Bookstore
	FindAllBook() []model.Bookstore
}
