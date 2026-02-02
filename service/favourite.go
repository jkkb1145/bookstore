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

func (f *FavouriteService) AddFavourite(userID, bookID int) error {
	err := f.FavouriteSVC.AddFavourite(userID, bookID)
	if err != nil {
		return err
	}
	return nil
}

func (f *FavouriteService) RemoveFavourite(userID, bookID int) (bool, error) {
	return f.FavouriteSVC.RemoveFavourite(userID, bookID)
}
