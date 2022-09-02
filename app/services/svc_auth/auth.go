// Package svc_auth 
package svc_auth

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/aoaostar/v8cdn_panel/app/enum"
    "github.com/aoaostar/v8cdn_panel/app/form"
    "github.com/aoaostar/v8cdn_panel/app/util"
    "github.com/aoaostar/v8cdn_panel/pkg"
    "github.com/cloudflare/cloudflare-go"
    "github.com/sirupsen/logrus"
    "net/url"
)

// GetToken
// https://api.cloudflare.com/host-gw.html
func (svc *authService) GetToken(ctx context.Context, params *form.LoginParam) (token string, err error) {
    userInfo := &util.User{}
    if params.UserApiKey == "" {
        if pkg.Conf.Cloudflare.HostKey == "" {
            err = errors.New("HostKey有误，无法使用密码登录")
            return
        }

        userInfo, err = svc.authByPartner(params)
    } else {
        userInfo, err = svc.authByKey(ctx, params)
    }

    if err != nil {
        return
    }

    token, err = util.GenerateToken(userInfo)

    return
}

func (svc *authService) authByKey(ctx context.Context, params *form.LoginParam) (userInfo *util.User, err error) {
    api, err := cloudflare.New(params.UserApiKey, params.Email)
    if err != nil {
        return
    }

    user, err := api.UserDetails(ctx)
    if err != nil {
        err = errors.New("key无效")
        return
    }

    if user.TwoFA {
        err = errors.New("您开启了两步登录验证，无法在此登陆")
        return
    }

    //accounts, _, err := api.Accounts(c, cloudflare.PaginationOptions{
    //	Page:    1,
    //	PerPage: 1,
    //})
    //if err != nil || len(accounts) <= 0 {
    //	return nil, errors.New("key无效")
    //}

    userInfo = &util.User{}
    userInfo.ID = user.ID
    userInfo.Email = params.Email
    userInfo.UserKey = user.APIKey
    userInfo.UserApiKey = params.UserApiKey
    userInfo.AuthType = string(enum.EnumUserAuthTypeApiKey)

    return
}

func (svc *authService) authByPartner(params *form.LoginParam) (userInfo *util.User, err error) {
    v8cdnPost := util.V8cdnPostForm("https://api.cloudflare.com/host-gw.html", url.Values{
        "act":              {"user_auth"},
        "host_key":         {pkg.Conf.Cloudflare.HostKey},
        "cloudflare_email": {params.Email},
        "cloudflare_pass":  {params.Password},
    })

    logrus.WithFields(logrus.Fields{
        "data": v8cdnPost,
    }).Debug()

    data := &form.CloudflareResponse{}
    err = json.Unmarshal([]byte(v8cdnPost), &data)
    if err != nil {
        err = errors.New("请求失败" + err.Error())
        return
    }

    if data.Result != "success" {
        message := "未知异常"
        if data.Msg != nil {
            message = fmt.Sprintf("%v", data.Msg)
        }

        err = errors.New(message)
        return
    }

    userInfo = &util.User{}
    userInfo.Email = data.Response.CloudflareEmail
    userInfo.UserKey = data.Response.UserKey
    userInfo.UserApiKey = data.Response.UserApiKey
    userInfo.AuthType = string(enum.EnumUserAuthTypePartner)

    return
}
