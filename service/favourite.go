package service

import "demo02/repository"

type FavouriteService struct {
	FavouriteSVC *repository.FavouriteDAO
}

func NewFavouriteService() *FavouriteService {
	return &FavouriteService{
		FavouriteSVC: repository.NewFavouriteDAO(),
	}
}
