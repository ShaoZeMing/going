package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App   app
	Mysql mysql
	Log   log
	Extra extra
}

type app struct {
	AppName    string `json:"app_name"`    //应用名称
	AppEnv     string `json:"app_env"`     //应用环境
	RunMode    string `json:"run_mode"`    //运行模式
	HttpListen string `json:"http_listen"` //应用监听HTTP服务=IP:端口
	GrpcListen string `json:"grpc_listen"` //应用监听grpc= IP:端口
}

type mysql struct {
	Dsn             string `json:"dsn"`
	MaxIdleConn     int    `json:"maxIdleConn"`
	MaxOpenConn     int    `json:"maxOpenConn"`
	ConnMaxLifetime int    `json:"connMaxLifetime"`
}

type extra struct {
	DingRobotKey string `json:"ding_robot_key"`
}

type log struct {
	Channel string `json:"channel"`  //日志记录渠道
	DirPath string `json:"dir_path"` //日志记录目录地址，绝对路径
	Days    int    `json:"days"`     //日志记录天数
}

func ConfInit() *Config {

	var conf Config
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
