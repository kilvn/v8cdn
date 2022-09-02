package config

import (
    "fmt"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
)

var Env Config

type Config struct {
    AppName    string `mapstructure:"app_name"`
    Debug      bool   `mapstructure:"debug"`
    Listen     string `mapstructure:"listen"`
    Static     string `mapstructure:"static"`
    Cloudflare struct {
        Email         string `mapstructure:"email"`
        HostKey       string `mapstructure:"host_key"`
        DefaultRecord string `mapstructure:"default_record"`
    }
    JwtSecret string `mapstructure:"jwt_secret"`
    RateLimit struct {
        Enabled      bool  `mapstructure:"enabled"`
        FillInterval int64 `mapstructure:"fill_interval"`
        Capacity     int64 `mapstructure:"capacity"`
    }
}

func InitConfig() {
    viper.SetConfigFile("./config.yaml") // 指定配置文件路径

    // 查找并读取配置文件
    if err := viper.ReadInConfig(); err != nil { // 处理读取配置文件的错误
        log.Panic(fmt.Errorf("读取配置出错: %s \n", err))
    }

    if err := viper.Unmarshal(&Env); err != nil {
        log.Panic(fmt.Errorf("解析配置出错: %s \n", err))
    }
}
