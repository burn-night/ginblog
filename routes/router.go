package routes

import (
	v1 "ginblo/api/v1"
	"ginblo/middleware"
	"ginblo/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{

		//User模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("uer/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		//上传文件
		auth.POST("upload", v1.UpLoad)

	}

	router := r.Group("api/v1")
	{

		//User模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("user/", v1.GetUsers)
		router.POST("login/", v1.Login)

		//分类模块的路由接口
		router.GET("category/", v1.GetCategorys)

		//文章模块的路由接口
		router.GET("article/", v1.GetArticles)
		router.GET("article/cate/:id", v1.GetCateArticle)
		router.GET("article/one/:id", v1.GetArticle)

	}

	r.Run()

}
