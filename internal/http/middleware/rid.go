package middleware

import (
	"fmt"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func SetRId() gin.HandlerFunc {
	return func(c *gin.Context) {

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
		c.Next()

	}
}
