package http

import (
	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-ops/docs"
	"go-ops/internal/api/message"
	"go-ops/internal/api/multi_cloud"
	"go-ops/internal/api/scheduler"
	"go-ops/internal/http/middleware"
)

func InitRouters(server *xgin.Server) {

	// set message center router
	send := server.Group("/api/v1/message")

	// set Cors middleware
	//send.Use(middleware.Cors())
	//send.Use(cors.Default())

	// set ReQueSetId middleware
	send.Use(middleware.SetRId())

	// set JWTauth middleware
	//send.Use(middleware.JWTAuth())

	{
		send.POST("/ding/add", message.AddDing)
		send.DELETE("/ding/delete", message.DelDing)
		send.PUT("/ding/update", message.UpdateDing)
		send.GET("/ding/query", message.QueryDing)
		send.POST("/ding/send", message.GobalSend.SendDing)

		send.POST("/mail/add", message.AddMail)
		send.DELETE("/mail/delete", message.DelMail)
		send.PUT("/mail/update", message.UpdateMail)
		send.GET("/mail/query", message.QueryMail)
		send.POST("/mail/send", message.GobalSend.SendDing)
	}

	// api swagger doc
	url := ginSwagger.URL("http://127.0.0.1:9900/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// task scheduler
	ts := server.Group("/api/v1/task-scheduler/dag")
	{
		ts.POST("/", scheduler.TriggerDagRun)
		ts.GET("/", scheduler.ListDagRun)
		ts.GET("/tasks", scheduler.TaskInstance)
	}

	// multi-cloud router
	template := server.Group("/api/v1/multi-cloud/template")
	{
		template.POST("", multi_cloud.CreateTemplateView)
		template.GET("/", multi_cloud.CetTemplateByAppKeyView)
		template.DELETE("/:id", multi_cloud.DeleteTemplateView)
	}
	instance := server.Group("/api/v1/multi-cloud/instance")
	{
		instance.POST("/create", multi_cloud.CreateInstanceView)
		instance.POST("/start", multi_cloud.StartInstanceView)
		instance.POST("/stop", multi_cloud.StopInstanceView)
		instance.POST("/reboot", multi_cloud.RebootInstanceView)
		instance.POST("/destroy", multi_cloud.DestroyInstanceView)
		instance.POST("/modify_instance_name", multi_cloud.ModifyInstanceNameView)
		instance.GET("/", multi_cloud.ListInstancesView)
		instance.GET("/:id", multi_cloud.InstanceDetailView)
	}
	region := server.Group("/api/v1/multi-cloud/regions")
	{
		region.GET("/", multi_cloud.ListRegionView)
	}
	account := server.Group("/api/v1/multi-cloud/account")
	{
		account.POST("", multi_cloud.CreateAccountView)
		account.GET("/", multi_cloud.ListAccountView)
		account.DELETE("/:id", multi_cloud.DeleteAccountView)
	}
}
