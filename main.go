package main

import (
	"LargeDataWrite/global"
	"LargeDataWrite/initialize"
	"LargeDataWrite/run"
)

func Init() {
	//初始化配置
	initialize.InitConfig("./config/config.toml")

	//初始化mysql
	global.DB = initialize.InitMysql()
}

func main() {
	//初始化
	Init()

	//注册用户
	run.Register()
}
