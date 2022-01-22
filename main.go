package main

import (

   "ToDoAPP/handlers"
   "fmt"
   "log"
   "net/http"
   "github.com/astaxie/beego/orm"
   _ "github.com/go-sql-driver/mysql"
   "github.com/spf13/cast"
   "github.com/spf13/viper"
   "time"
   "github.com/akrylysov/algnhsa"

)

var (
	ENV string
)

func init() {
	setUpViper()
	registerDatabase()

}

func main() {

    http.HandleFunc("/todo/user",handlers.UserHandler)
	http.HandleFunc("/todo/task",handlers.TaskHandler)


	fmt.Println("Start...")
	if ENV =="dev"{
		_ = http.ListenAndServe(":3000", nil)

	}
	algnhsa.ListenAndServe(http.DefaultServeMux, nil)
}


//function to register the database to beego orm
func registerDatabase() {
	ENV = cast.ToString(viper.Get("runmode"))
	mysql := viper.Get(ENV + ".mysql").(map[string]interface{})
	mysqlConf := mysql["user"].(string) + ":" + mysql["password"].(string) + "@tcp(" + mysql["host"].(string) + ")/" + mysql["database"].(string)
	log.Println("conf", mysqlConf)
	orm.RegisterDataBase("default", "mysql", mysqlConf)
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Kolkata")
	orm.Debug = true
}

//set up config file from conf folder
func setUpViper() {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetEnvPrefix("global")
}