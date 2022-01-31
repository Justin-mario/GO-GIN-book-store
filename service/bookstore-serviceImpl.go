package service

import "github.com/Justin-mario/GO-GIN-book-store/model"

type bookService struct {
	books []model.Bookstore
}

func New() BookService {
	return &bookService{}
}

func (service *bookService) Save(book model.Bookstore) model.Bookstore {
	service.books = append(service.books, book)
	return book
}

func (service *bookService) FindAllBook() []model.Bookstore {
	return service.books
}
