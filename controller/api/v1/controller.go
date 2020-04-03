package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jshzhj/ginframe/models"
	"net/http"
	"time"
)

//url-->controller(接收数据,参数验证,调用logic层,返回数据)
//->logic(逻辑层:数据整合,逻辑判断)->models(模型层:mysql的增删改查)

func Indexhandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func CreateATodo(c *gin.Context) {

	//1.从请求中拿出数据,绑定到结构体
	var todo models.Todos
	c.BindJSON(&todo)

	//2.存入数据库
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func GetTodoList(c *gin.Context) {
	var todoList []models.Todos

	if err := models.GetTodoList(&todoList); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var todo models.Todos
	time.Sleep(10*time.Second)
	//先查再更新
	if err := models.GetATodo(id, &todo); err != nil {
		c.JSON(500, err.Error())
		return
	}

	//赋值
	c.ShouldBind(&todo)
	fmt.Printf("%v\n", todo)

	//更新
	if err := models.UpdateATodo(id, &todo); err != nil {
		//更新失败
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		//更新成功
		c.JSON(200, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var todo models.Todos
	if err := models.DeleteATodo(id, &todo); err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "id:" + id + "删除成功",
		})
	}
}
