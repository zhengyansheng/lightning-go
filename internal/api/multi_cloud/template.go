package multi_cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-ops/internal/models/multi_cloud"
	"go-ops/pkg/tools"
)

// @Summary 创建
// @Description 创建多云模版
// @Accept  application/json
// @Produce application/json
// @Param   data   body multi_cloud_template.CloudTemplate  true   "Body数据格式"
// @Success 200 {string}  {"code": 0, "data": "", "message": ""}
// @Failure 500 {string}  {"code": -1, "data": "", "message": ""}
// @Router /api/v1/multi-cloud/template [post]
func CreateTemplateView(c *gin.Context) {
	var (
		err      error
		template multi_cloud.CloudTemplate
	)
	// 参数验证
	err = c.ShouldBindWith(&template, binding.JSON)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// 格式化打印验证的数据
	tools.PrettyPrint(template)

	// 保存
	err = template.Create()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}

	// 返回
	tools.JSONOk(c, "Created successfully")
}

// @Summary 创建
// @Description 查询模版基于app_key
// @Accept  application/json
// @Produce application/json
// @Param   data   body multi_cloud_template.CloudTemplate  true   "Body数据格式"
// @Success 200 {string}  {"code": 0, "data": "", "message": ""}
// @Failure 500 {string}  {"code": -1, "data": "", "message": ""}
// @Router /api/v1/multi-cloud/template [get]
func CetTemplateByAppKeyView(c *gin.Context) {
	var (
		err       error
		template  multi_cloud.CloudTemplate
		templates []multi_cloud.CloudTemplate
	)
	// 参数验证
	appKey, ok := c.GetQuery("app_key")
	if !ok {
		tools.JSONFailed(c, tools.MSG_ERR, "app_key is required.")
		return
	}

	// 保存

	template.AppKey = appKey

	if v, ok := c.GetQuery("account"); ok {
		template.Account = v
	}

	if v, ok := c.GetQuery("region_id"); ok {
		template.RegionId = v
	}

	if v, ok := c.GetQuery("env"); ok {
		template.Env = v
	}

	// 格式化打印验证的数据
	tools.PrettyPrint(template)

	templates, err = template.ListByAppKey()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}

	// 返回
	tools.JSONOk(c, templates)
}


func DeleteTemplateView(c *gin.Context) {
	var template  multi_cloud.CloudTemplate

	pk := c.Param("id")
	pkUint, _ := tools.StringToUint(pk)
	template.ID = pkUint
	err := template.Delete()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}

	// 返回
	tools.JSONOk(c, "Delete ok.")
}