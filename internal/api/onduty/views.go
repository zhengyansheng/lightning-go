package onduty

import (
	"go-ops/internal/service/onduty"
	"go-ops/pkg/tools"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary 创建
// @Description 创建树结构数据
// @Accept  application/json
// @Produce application/json
// @Param   data   body service.OnDutySerializer  true   "Body数据格式"
// @Success 200 {string}  {"code": 0, "data": "", "message": ""}
// @Failure 500 {string}  {"code": -1, "data": "", "message": ""}
// @Router /api/v1/on-duty [post]
func CreateOnDuty(c *gin.Context) {
	var (
		err        error
		serializer service.OnDutySerializer
	)
	// 参数验证
	err = c.ShouldBindWith(&serializer, binding.JSON)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// 格式化打印验证的数据
	tools.PrettyPrint(serializer)

	// 保存
	err = serializer.Create()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}

	// 返回
	tools.JSONOk(c, "Created successfully")
}

// @Summary 修改
// @Description 修改节点
// @Accept  application/json
// @Produce application/json
// @Param   id  path   int true     "id"
// @Param   data   body service.OnDutySerializer  true   "Body数据格式"
// @Success 200 {string}  {"code": 0, "data": "", "message": ""}
// @Failure 500 {string}  {"code": -1, "data": "", "message": ""}
// @Router /api/v1/on-duty/{id} [put]
func UpdateOnDuty(c *gin.Context) {
	var (
		err        error
		serializer service.OnDutySerializer
	)
	// 参数验证
	pk := c.Param("id")
	err = c.ShouldBindWith(&serializer, binding.JSON)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}

	// 赋值
	id, err := tools.StringToUint(pk)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	serializer.ID = id

	// 格式化打印验证的数据
	tools.PrettyPrint(serializer)

	err = serializer.Update()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}

	// 返回
	tools.JSONOk(c, "Updated successfully")
}

// @Summary 删除
// @Description 删除节点
// @Accept  application/json
// @Produce application/json
// @Param   id  path   int true     "id"
// @Success 200 {string}  {"code": 0, "data": "", "message": ""}
// @Failure 500 {string}  {"code": -1, "data": "", "message": ""}
// @Router /api/v1/on-duty/{id} [delete]
func DeleteOnDuty(c *gin.Context) {
	var (
		err        error
		serializer service.OnDutySerializer
	)
	// 参数验证
	pk := c.Param("id")

	// 赋值
	id, err := tools.StringToUint(pk)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	serializer.ID = id

	// 格式化打印验证的数据
	tools.PrettyPrint(serializer)

	err = serializer.Delete()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}

	// 返回
	tools.JSONOk(c, "Delete successfully")

}
