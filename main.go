package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo/dao"
	"todo/models"
	"todo/routers"
)

func main() {
	if err := dao.ConnectMysql(); err != nil {
		fmt.Println(err)
		return
	}
	dao.Db.AutoMigrate(&models.Todo{})
	defer dao.Close()
	r := gin.Default()
	routers.SetupRouter(r)
	r.Run(":8081")
}