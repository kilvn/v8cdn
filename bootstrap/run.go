package bootstrap

import (
    "github.com/aoaostar/v8cdn_panel/config"
    "github.com/gin-contrib/static"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
)

func Run() {
    config.InitConfig()
    InitLog()
    InitRateLimit()
    InitValidator()

    engine := gin.Default()
    if !config.Env.Debug {
        //记录panic错误
        engine.Use(gin.RecoveryWithWriter(log.StandardLogger().Writer()))
    }
    engine.Use(static.Serve("/", static.LocalFile(config.Env.Static, false)))
    engine.NoRoute(func(c *gin.Context) {
        c.File(config.Env.Static + "/index.html")
    })

    //初始化路由
    InitRouter(engine)

    //初始化数据库
    //初始化缓存

    err := engine.Run(config.Env.Listen)
    if err != nil {
        log.Error("Gin启动失败：%+v\n", err)
        return
    }
}
