package commons

import (
	"strings"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"conf"
	"log"
	"os"
)

var dbInstance *gorm.DB
func  init() {
	section := conf.GetConfigSection("db")
	db_connect_string := strings.Join([]string{section["user"], ":", section["password"], "@tcp(", section["host"], ":", section["port"], ")/", section["name"], "?charset=", section["charset"], "&parseTime=", section["parsetime"]}, "")

	instance, err := gorm.Open(section["driver"], db_connect_string)
	if err != nil {
		ERROR("db connect failed:", err.Error())
	}
	//defer MY_DB.Close()
	instance.SetLogger(log.New(os.Stdout, "\r\n", 0))
	dbInstance = instance
}

func GetDbInstance()  *gorm.DB{
	return dbInstance;
}