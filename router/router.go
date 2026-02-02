package router

import (
	"demo02/controller"
	"demo02/midleWare"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{

		user := v1.Group("/user")
		{
			user.POST("/userRegistrateIn", controller.NewUserCTL().UserRegistIn)
			user.POST("/userLogIn", controller.NewUserCTL().UserLogIn)

		}

		auth := user.Group("")
		{
			//中间件实现鉴权
			auth.Use(midleWare.AdminAuthMiddleware())
			{
				auth.GET("/profile", controller.NewUserCTL().GetUserProfile)    //获取用户信息
				auth.PUT("/profile", controller.NewUserCTL().UpdateUserProfile) //修改用户信息
				auth.PUT("/password", controller.NewUserCTL().ChangePassword)   //修改密码
				auth.DELETE("/logout", controller.NewUserCTL().Logout)          //登出
			}
		}

		book := v1.Group("/books")
		{
			book.GET("/popularbooks", controller.NewBookController().GetPopBooks)
			book.GET("/searchbooks", controller.NewBookController().SearchBooks)
		}
		favourite := v1.Group("/favourite")
		{
			favourite.Use(midleWare.JWTAuthMiddleware())
			{
				favourite.POST("/favourite", controller.NewFavouriteController().AddFavourite)
				favourite.DELETE("/favourite", controller.NewFavouriteController().RemoveFavourite)
				favourite.GET("/list", controller.NewFavouriteController().GetUserFavourite)
			}

		}
	}
	return r
}
