package controllers

import (
    "github.com/aoaostar/v8cdn_panel/app/form"
    "github.com/aoaostar/v8cdn_panel/app/services/svc_auth"
    "github.com/aoaostar/v8cdn_panel/app/util"
    "github.com/gin-gonic/gin"
)

type Auth struct {
}

func (i *Auth) Login(ctx *gin.Context) {
    params := &form.LoginParam{}

    if err := ctx.ShouldBind(&params); err != nil {
        e, _ := util.FomateValidateError(err)

        util.JSON(ctx, "error", "无效的参数"+e, nil)
    }

    token, err := svc_auth.Instance(ctx).GetToken(ctx, params)
    if err != nil {
        util.JSON(ctx, "error", "无效的参数"+err.Error(), nil)
    }

    util.JSON(ctx, "ok", "success", gin.H{
        "token": token,
    })
}
