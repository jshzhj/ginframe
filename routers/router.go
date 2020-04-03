package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jshzhj/ginframe/controller/api/v1"
	"github.com/jshzhj/ginframe/routers/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger()) //gin自带log
	r.Use(gin.Recovery()) //gin自带recover

	//告诉gin去哪里找静态文件,前面是路由,后面是文件路径
	r.Static("/static", "asset/static")

	//告诉gin去哪里找模版文件(html)
	r.LoadHTMLGlob("asset/templates/*")

	//下面是api路由(局部中间件)
	r.GET("/", middleware.SpendTimeMiddleware(),controller.Indexhandler)

	//路由分组(分组中间件)
	v1Group := r.Group("/v1")
	v1Group.Use(middleware.SpendTimeMiddleware())
	{
		//添加
		v1Group.POST("/todo", controller.CreateATodo)

		//查看
		v1Group.GET("/todo", controller.GetTodoList)

		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return r
}
