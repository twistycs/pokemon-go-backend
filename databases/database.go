package database

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

//SetUpDb Database
func SetUpDb() *DBConfig {
	viper.SetConfigFile("./dbconfig.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Can not read config file.")
	}
	var configDBs DBConfig
	viper.Unmarshal(&configDBs)
	return &configDBs
}

//ReadConfigDBFile print
func (db *DBConfig) ReadConfigDBFile() string {
	fmt.Printf("%+v", db)
	return fmt.Sprintf(db.User + ":" + db.Password + "@tcp(" + db.Host + ":" + strconv.Itoa(db.Port) + ")/" + db.DBName + "?charset=utf8&parseTime=True&loc=Local")
}
