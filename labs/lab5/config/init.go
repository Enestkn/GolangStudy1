package config

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := "sqlserver://sa:123456@localhost:1433?database=HumanResources&trusted+connection=yes&app+name=Human_Resources"
	if DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB Connected")
	}

}
