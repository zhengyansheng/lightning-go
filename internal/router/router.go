package router

import (
	_ "lightning-go/docs"
	"lightning-go/internal/api/multi_cloud"
	"lightning-go/internal/api/scheduler"
	"lightning-go/internal/router/middleware"

	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	// api swagger doc
	url := ginSwagger.URL("http://127.0.0.1:9900/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// task scheduler
	ts := server.Group("/api/v1/task-scheduler/dag")
	{
		ts.POST("/:dagName", scheduler.TriggerDagRun)
		ts.GET("/", scheduler.ListDagRun)
	}

	// multi-cloud router
	template := server.Group("/api/v1/multi-cloud/template")
	{
		template.POST("", multi_cloud.CreateTemplateView)
		template.GET("/", multi_cloud.CetTemplateByAppKeyView)
		template.DELETE("/:id", multi_cloud.DeleteTemplateView)
	}
	instance := server.Group("/api/v1/multi-cloud/instance")
	instance.Use(middleware.LifeCycle()) // 记录生命周期事件
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
