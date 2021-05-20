package main

import (
	"encoding/json"
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"io/ioutil"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {

	global.GVA_VP = core.Viper()         // 初始化Viper
	global.GVA_LOG = core.Zap()          // 初始化zap日志库
	global.GVA_ERRCODE = readErrorCode() // 读取错误码
	global.GVA_DB = initialize.Gorm()    // gorm连接数据库
	initialize.Timer()
	// if global.GVA_DB != nil {
	// 	initialize.MysqlTables(global.GVA_DB) // 初始化表
	// 	// 程序结束前关闭数据库链接
	db, _ := global.GVA_DB.DB()
	defer db.Close()
	// }

	core.RunWindowsServer()
}

func readErrorCode() map[string]map[string]string {
	filename := "./docs/error_code.json"
	// 读取文件
	f, e := ioutil.ReadFile(filename)
	if e != nil {
		panic(filename + e.Error())
	}
	// 解析文件
	m := make(map[string]map[string]string)
	if err := json.Unmarshal(f, &m); err != nil {
		panic(filename + err.Error())
	}

	return m
}
