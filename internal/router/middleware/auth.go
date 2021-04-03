package middleware

import (
	"fmt"

	"github.com/douyu/jupiter/pkg/conf"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"

	"lightning-go/pkg/tools"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		authSuccess := tools.JSONResult{}
		jwt := c.Request.Header.Get("authorization")
		// 处理请求
		if len(jwt) < 1 {
			tools.JSONFailed(c, 403, "The token is empty")
			c.Abort()
			return
		}

		client := resty.New()
		_, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("authorization", jwt).
			SetBody(map[string]interface{}{
				"path":   c.Request.URL.Path,
				"method": c.Request.Method,
			}).
			SetResult(&authSuccess).
			Post(conf.GetString("go-ops.hosts.auth"))

		if err != nil {
			tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("auth err :%v ", err))
			c.Abort()
			return
		}

		if authSuccess.Code != 0 {
			tools.JSONFailed(c, authSuccess.Code, authSuccess.Message)
			c.Abort()
			return
		}
		c.Next()
	}
}
