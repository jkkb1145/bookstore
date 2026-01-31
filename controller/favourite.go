package controller

import (
	"demo02/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavouriteController struct {
	FavouriteCTL *service.FavouriteService
}

func NewFavouriteController() *FavouriteController {
	return &FavouriteController{
		FavouriteCTL: service.NewFavouriteService(),
	}
}

func GetUserIDFromHeader(c *gin.Context) int {
	userID, exists := c.Get("userID")
	if !exists {
		return 0
	}
	return userID.(int)
}

func (f *FavouriteController) AddFavourite(c *gin.Context) {
	userID := GetUserIDFromHeader(c)
	if userID==0{
		c.JSON(http.StatusUnauthorized,gin.H{
			"code":-1,
			"msg":"未登录",
		})
		return
	}

}
