package main

import (
	"box/base/validator"
	"box/controller/account"
	"box/controller/calculator/cost"
	"box/controller/medical"
	"box/middleware"
	"box/preload"
	"github.com/gin-gonic/gin"
)

func main() {
	preload.InitConfig()
	preload.InitLog()
	preload.InitMySQL()
	//preload.GenerateModel(preload.DB)

	//初始化 validator 错误翻译器
	validator.Init()

	engine := gin.Default()
	r := engine.Group("/api", gin.Recovery())

	accountGroup := r.Group("/account")
	{
		accountGroup.POST("/register", account.Register)
		accountGroup.POST("/login", account.Login)
	}

	calculator := r.Group("/calculator")
	{
		calculator.GET("/subwayCost", cost.SubwayCost)
	}

	//ofy := r.Group("/ofy", middleware.JWT())
	r.Use(middleware.JWT())
	medicalGroup := r.Group("/medical")
	{
		//添加成长记录
		medicalGroup.POST("/create", medical.Add)
		//查找成长记录
		medicalGroup.GET("/retrieve", medical.Retrieve)
		//更新成长记录
		medicalGroup.GET("/update", medical.Update)
		//删除成长记录
		medicalGroup.GET("/delete", medical.Retrieve)
	}
	//scheduleGroup := ofy.Group("/schedule")
	//{
	//	scheduleGroup.POST("/create", schedule.Add)
	//	scheduleGroup.GET("/retrieve", schedule.Retrieve)
	//	scheduleGroup.POST("/update", account.Login)
	//	scheduleGroup.POST("/delete", account.Login)
	//}
	err := engine.Run()
	if err != nil {
		panic("http engine run failed")
	}
}
