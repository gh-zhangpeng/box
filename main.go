package main

import (
	"box/base/validator"
	"box/controller/account"
	"box/controller/calculator/cost"
	"box/controller/medical"
	"box/controller/schedule"
	"box/middleware"
	"box/preload"
	"github.com/gin-gonic/gin"
	"reflect"
)

func main() {
	preload.InitConfig()
	preload.InitLog()
	preload.InitMySQL()
	preload.GenerateModel(preload.DB)

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

	r.Use(middleware.JWT())
	medicalGroup := r.Group("/medical")
	{
		//添加成长记录
		medicalGroup.POST("/create", medical.Create)
		//查找成长记录
		medicalGroup.GET("/retrieve", medical.Retrieve)
		//更新成长记录
		medicalGroup.POST("/update", medical.Update)
		//删除成长记录
		medicalGroup.POST("/delete", medical.Delete)
	}
	scheduleGroup := r.Group("/schedule")
	{
		scheduleGroup.POST("/create", schedule.Create)
		scheduleGroup.GET("/retrieve", schedule.Retrieve)
		scheduleGroup.POST("/update", schedule.Update)
		scheduleGroup.POST("/delete", schedule.Delete)
	}
	err := engine.Run()
	if err != nil {
		panic("http engine run fail")
	}
}

func StructToMap(obj interface{}) map[string]interface{} {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < typ.NumField(); i++ {
		data[typ.Field(i).Name] = val.Field(i).Interface()
	}
	return data
}
