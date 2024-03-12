package conf

import (
	"go_mall/dao"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	// RedisDb     string
	// RedisAddr   string
	// RedisPw     string
	// RedisDbName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		panic("读取配置文件出错")
	}
	LoadServer(file)
	LoadMysqlData(file)
	//LoadRedis(file)
	LoadEmail(file)
	LoadPhotoPath(file)

	//mysql 读(80%)主
	pathRead := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	//mysql 写(20%)从
	pathWrite := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	dao.Database(pathRead, pathWrite)

}

// func LoadRedis(file *ini.File) {
// 	RedisDb = file.Section("redis").Key("RedisDb").String()
// 	RedisAddr = file.Section("redis").Key("RedisAddr").String()
// 	RedisPw = file.Section("redis").Key("RedisPw").String()
// 	RedisDbName = file.Section("redis").Key("RedisDbName").String()
// }

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}

func LoadPhotoPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}
