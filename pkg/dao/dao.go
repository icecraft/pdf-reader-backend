package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/config"
)

var (
	Db           *gorm.DB
	htmlTemplate *HTMLTemplate
)

func InitDAO(cf config.MySqlConf, tplFolder string) error {
	var err error

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", cf.USER, cf.PASSWD, cf.HOST, cf.DB)

	Db, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN: mysqlInfo,
		}),
		&gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Silent),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	if err != nil {
		fmt.Println("Error connecting to database : error=%w", err)
		return nil
	}

	htmlTemplate = Default(tplFolder, ".sql")
	RegisterSqlTemplate(htmlTemplate)

	return nil
}

func Sync2() error {
	return nil
}


func TruncateTables() error {
	/*
	if err := Db.Exec("truncate dataset").Error; err != nil {
		return err
	}
	*/
	return nil
}
