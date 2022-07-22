package preload

import (
	"fmt"
	box_lib "github.com/gh-zhangpeng/box-lib"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

var (
	DB *gorm.DB
)

func InitMySQL() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.Get("db.user"),
		viper.Get("db.password"),
		viper.Get("db.address"),
		viper.GetInt("db.port"),
		viper.Get("db.db"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Errorf("open db fail: %w \n", err))
	}
	DB = db
}

func GenerateModel(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		Mode:              gen.WithDefaultQuery,
		OutPath:           "./dal/query",
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldCoverable:    true,
		WithUnitTest:      true,
	})
	g.UseDB(db)
	g.WithModelNameStrategy(func(tableName string) (modelName string) {
		if strings.HasPrefix(tableName, "tbl") {
			return box_lib.FirstUpper(tableName[3:])
		}
		if strings.HasPrefix(tableName, "table") {
			return box_lib.FirstUpper(tableName[5:])
		}
		return box_lib.FirstUpper(tableName)
	})
	g.WithFileNameStrategy(func(tableName string) (fileName string) {
		if strings.HasPrefix(tableName, "tbl") {
			return box_lib.FirstLower(tableName[3:])
		}
		if strings.HasPrefix(tableName, "table") {
			return box_lib.FirstLower(tableName[5:])
		}
		return tableName
	})

	g.WithDataTypeMap(map[string]func(detailType string) (dataType string){
		"int": func(detailType string) (dataType string) {
			if strings.Contains(detailType, "unsigned") {
				return "int64"
			}
			return "int64"
		},
	})
	//g.GenerateModelAs("tblUser", "User")
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
