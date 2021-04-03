package middleware

import (
    "fmt"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func LifeCycle() gin.HandlerFunc {
	return func(c *gin.Context) {

	    //1. 生成 Request-Id
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			id, err := uuid.NewV4()
			if err != nil {
				xlog.Infof("requeset ID err :%v", err)
				return
			}
			requestId = fmt.Sprintf("%s", id)
		}
		c.Set("X-Request-Id", requestId)
		c.Writer.Header().Set("X-Request-Id", requestId)

        //2. 记录事件生命周期
        /*
            action
            uri
            method
            instanceId
            Body
            Response
            IsSuccess
            createUser
            execState
         */

		c.Next()

	}
}
