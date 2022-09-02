// Package form 
package form

type (
    LoginParam struct {
        Email      string `json:"email" binding:"required,email"`
        Password   string `json:"password" binding:"required_without=UserApiKey"`
        UserApiKey string `json:"user_api_key" binding:"omitempty,min=1"`
    }
    UserInfo struct {
        Email      string
        UserKey    string
        UserApiKey string
    }
    CloudflareResponse struct {
        Msg     interface{} `json:"msg"`
        Request struct {
            Act string `json:"act"`
        } `json:"request"`
        Response struct {
            CloudflareEmail string      `json:"cloudflare_email"`
            UniqueId        interface{} `json:"unique_id"`
            UserApiKey      string      `json:"user_api_key"`
            UserKey         string      `json:"user_key"`
        } `json:"response"`
        Result string `json:"result"`
    }
)
