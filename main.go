package main

import (
	"github.com/gin-gonic/gin"
	"todo/dao"
	"todo/models"
	"todo/routers"
)

func main() {
	if err := dao.ConnectMysql(); err != nil {
		panic(err)
	}
	if err := dao.Db.AutoMigrate(&models.Todo{}); err != nil {
		panic(err)
	}
	defer dao.Close()
	r := gin.Default()
	routers.SetupRouter(r)
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
