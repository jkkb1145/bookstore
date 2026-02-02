package controller

import (
	"demo02/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FavouriteController struct {
	FavouriteCTL *service.FavouriteService
}

func NewFavouriteController() *FavouriteController {
	return &FavouriteController{
		FavouriteCTL: service.NewFavouriteService(),
	}
}

// 从请求头中获取用户ID
func GetUserIDFromHeader(c *gin.Context) int {
	userID, exists := c.Get("userID")
	if !exists {
		return 0
	}
	return userID.(int)
}

func (f *FavouriteController) AddFavourite(c *gin.Context) {
	userID := GetUserIDFromHeader(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "未登录",
		})
		return
	}
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  -1,
			"msg":   "无效的书籍ID",
			"error": err,
		})
		return
	}
	err = f.FavouriteCTL.AddFavourite(userID, bookID)
	if err != nil {
		c.JSON(500, gin.H{
			"code":  -1,
			"msg":   "添加书籍到收藏夹失败",
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  0,
		"msg":   "书籍已添加至收藏夹",
		"error": nil,
	})
}

func (f *FavouriteController) RemoveFavourite(c *gin.Context) {
	userID := GetUserIDFromHeader(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "未登录",
		})
		return
	}

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"code":  -1,
			"msg":   "无效书籍ID",
			"error": err,
		})
		return
	}
	affectedRows, err := f.FavouriteCTL.RemoveFavourite(userID, bookID)
	if err != nil {
		c.JSON(500, gin.H{
			"code":  -1,
			"msg":   "移除收藏失败",
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":          0,
		"msg":           "已移除收藏",
		"affected_rows": affectedRows,
		"error":         nil,
	})
}
