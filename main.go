package main

import (
	"box/controller/calculator/cost"
	"box/controller/ofy/account"
	"box/controller/ofy/medical"
	"box/middleware"
	"box/preload"
	"github.com/gin-gonic/gin"
)

func main() {
	preload.InitConfig()
	preload.InitLog()
	preload.InitMySQL()
	//preload.GenerateModel(preload.DB)

	r := gin.Default()
	r.Use(gin.Recovery())

	ofy := r.Group("/ofy")
	accountGroup := ofy.Group("/account")
	{
		accountGroup.POST("/register", account.Register)
		accountGroup.POST("/login", account.Login)
	}
	ofy.Use(middleware.JWT())
	//scheduleGroup := ofy.Group("/schedule")
	//{
	//	scheduleGroup.POST("/create", schedule.Add)
	//	scheduleGroup.GET("/retrieve", schedule.Retrieve)
	//	scheduleGroup.POST("/update", account.Login)
	//	scheduleGroup.POST("/delete", account.Login)
	//}
	medicalGroup := ofy.Group("/medical")
	{
		//添加成长记录
		medicalGroup.POST("/add", medical.Add)
		//获取成长记录
		medicalGroup.GET("/retrieve", medical.Retrieve)
	}
	calculator := r.Group("/calculator")
	{
		calculator.GET("/subwayCost", cost.SubwayCost)
	}

	err := r.Run()
	if err != nil {
		panic("http engine run fail")
	}
}
