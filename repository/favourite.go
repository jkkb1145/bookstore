package repository

import (
	"database/sql"
	"demo02/global"
)

type FavouriteDAO struct {
	FavouriteDB *sql.DB
}

func NewFavouriteDAO() *FavouriteDAO {
	return &FavouriteDAO{
		FavouriteDB: global.GetDb(),
	}
}
