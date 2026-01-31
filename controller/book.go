package controller

import (
	"demo02/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookController struct {
	BookCTL *service.BookService
}

func NewBookController() *BookController {
	return &BookController{
		BookCTL: service.NewBookService(),
	}
}

func (b *BookController) GetPopBooks(c *gin.Context) {

	err, books := b.BookCTL.GetPopBooks()
	if err != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"msg":  "获取畅销书榜失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": books,
		"msg":  "获取热销书榜成功",
	})
	return
}

func (b *BookController) SearchBooks(c *gin.Context) {
	keyword := c.Query("key")
	if keyword == "" {
		c.JSON(400, gin.H{
			"code": -1,
			"msg":  "关键字不可以为空",
		})
		return
	}
	//分页逻辑
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10")) //将从query参数中获取到的字符类型数据转为整形
	books, total, err := b.BookCTL.SearchBooks(keyword, page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{
			"code":  -1,
			"msg":   "搜索书籍失败",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"data": gin.H{
			"books":     books,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
		"msg": "搜索书籍成功",
	})
	return
}
