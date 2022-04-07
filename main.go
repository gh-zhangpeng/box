package main

import (
	"box/controller/calculator/cost"
	"box/controller/ofy/account"
	"box/controller/ofy/medical"
	"box/controller/ofy/schedule"
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
	scheduleGroup := ofy.Group("/schedule")
	{
		scheduleGroup.POST("/create", schedule.Create)
		scheduleGroup.GET("/retrieve", schedule.Retrieve)
		scheduleGroup.POST("/update", account.Login)
		scheduleGroup.POST("/delete", account.Login)
	}
	medicalGroup := ofy.Group("/medical")
	{
		//创建成长记录
		medicalGroup.POST("/addMedical", medical.AddRecord)
		//获取成长记录
		medicalGroup.GET("/getMedicals", medical.GetRecords)
		//growthGroup.POST("/records", account.Login)
		//growthGroup.POST("/records", account.Login)
	}

	calculator := r.Group("/calculator")
	{
		calculator.GET("/getSubwayCost", cost.SubwayCost)
	}

	err := r.Run()
	if err != nil {
		panic("http engine run fail")
	}
}
