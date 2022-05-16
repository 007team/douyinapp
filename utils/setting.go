package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode		string
	HttpPort	string
	JwtKey		string

	Db 			string
	DbHost 		string
	DbPort 		string
	DbUser 		string
	DbPassWord 	string
	DbName 		string

	AccessKey 	string
	SecretKey 	string
	Bucket 		string
	QiniuServer	string
)

func init(){
	file,err:=ini.Load("config/config.ini")
	if err!=nil{
		log.Println("配置文件读取错误")
	}

	LoadServer(file)
	LoadDatabase(file)
	// LoadQiniu(file)
}

func LoadServer(file *ini.File){
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	JwtKey = file.Section("server").Key(
		"JwtKey").MustString("zhaoyf")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
}


func LoadDatabase(file *ini.File){
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("lroot")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuSever").String()
}