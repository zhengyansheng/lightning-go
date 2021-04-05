package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lightning-go/internal/models/multi_cloud"
	"lightning-go/pkg/tools"
	"strings"

	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

/*
	创建
	变更
		开机
		关机
		重启
		升配
		降配
	下线
*/

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LifeCycle() gin.HandlerFunc {
	return func(c *gin.Context) {

		//1. 生成 Request-Id
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			id, err := uuid.NewV4()
			if err != nil {
				xlog.Infof("request ID err :%v", err)
				return
			}
			requestId = fmt.Sprintf("%s", id)
		}
		c.Set("X-Request-Id", requestId)
		c.Writer.Header().Set("X-Request-Id", requestId)

		//2. 记录事件生命周期
		//cCp := c.Copy()
		rawData, _ := c.GetRawData()
		// 重新赋值
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))

		cycle := multi_cloud.InstanceLifeCycle{
			Uri:      c.Request.URL.Path,
			Method:   c.Request.Method,
			Query:    c.Request.URL.RawQuery,
			Body:     rawData,
			RemoteIp: c.ClientIP(),
		}

		headerAuthorization := c.Request.Header.Get("Authorization")
		if headerAuthorization != "" {
			authArr := strings.Split(headerAuthorization, " ")
			if len(authArr) == 2 {
				cycle.CreateUser = authArr[1]
			}
		}

		// https://github.com/gin-gonic/gin/issues/1363
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		c.Next()

		// 4. attach field response
		cycle.Response = []byte(w.body.String())

		// 5. attach field IsSuccess
		var jsonResult tools.JSONResult
		err := json.Unmarshal([]byte(w.body.String()), &jsonResult)
		if err == nil {
			if jsonResult.Code == tools.MSG_OK {
				cycle.IsSuccess = true
				if vInstanceInfo, ok := jsonResult.Data.(map[string]interface{}); ok {
					cycle.InstanceId = vInstanceInfo["instance_id"].(string)
				} else {
					bodyInfo, err := tools.StringToMap(rawData)
					if err != nil {
						xlog.Infof("string to map err :%v", err)
					}
					if instanceId, ok := bodyInfo["instance_id"].(string); ok {
						cycle.InstanceId = instanceId
					}
					fmt.Println(4)

				}
			}
		} else {
			xlog.Infof("middleware attach instance_id and is_success field err :%v", err)
		}

		// 6. save
		_ = cycle.Create()

	}
}
