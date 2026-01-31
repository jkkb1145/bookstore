package service

import (
	"demo02/model"
	"demo02/repository"
)

type BookService struct {
	BookSVC *repository.BookDAO
}

func NewBookService() *BookService {
	return &BookService{
		BookSVC: repository.NewBookDAO(),
	}
}

func (b *BookService) GetPopBooks() (error, *[]model.Book) {
	err, books := b.BookSVC.GetPopBooks()
	if err != nil {
		return err, nil
	}
	return nil, books
}

func (b *BookService) SearchBooks(keywords string, page, pageSize int) (*[]model.Book, int, error) {
	return b.BookSVC.SearchBooks(keywords, page, pageSize)
}
