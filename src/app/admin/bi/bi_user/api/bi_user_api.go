// Package api  BiUserApi 自动生成模板
// @description <TODO description class purpose>
// @author Administrator
// @File: bi_user
// @version 1.0.0
// @create 2023-08-16 14:53:36
package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"manager-gin/src/app/admin/bi/bi_user/service"
	"manager-gin/src/app/admin/bi/bi_user/service/view"
	"manager-gin/src/common"
	response "manager-gin/src/common/response"
	"manager-gin/src/global"
	"manager-gin/src/utils"
	"net/http"
	"strings"
)

type BiUserApi struct {
}

var biUserService = service.BiUserServiceApp

// Create 创建BiUser
// @Summary 创建BiUser
// @Router /biUser/create [post]
func (biUserApi *BiUserApi) Create(c *gin.Context) {
	var biUserView view.BiUserView
	_ = c.ShouldBindJSON(&biUserView)
	biUserView.CreateTime = utils.GetCurTimeStr()
	biUserView.UpdateTime = utils.GetCurTimeStr()
	biUserView.Id = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	if err := biUserService.Create(&biUserView); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除BiUser
// @Summary 删除BiUser
// @Router /biUser/delete [delete]
func (biUserApi *BiUserApi) Delete(c *gin.Context) {
	var id string
	_ = c.ShouldBindJSON(&id)
	if err := biUserService.Delete(id); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除BiUser
// @Summary 批量删除BiUser
// @Router /biUser/deleteByIds [delete]
func (biUserApi *BiUserApi) DeleteByIds(c *gin.Context) {
	var ids []string
	_ = c.ShouldBindJSON(&ids)
	if err := biUserService.DeleteByIds(ids); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新BiUser
// @Summary 更新BiUser
// @Router /biUser/update [put]
func (biUserApi *BiUserApi) Update(c *gin.Context) {
	var biUserView view.BiUserView
	err := c.ShouldBindJSON(&biUserView)
	//id := c.Query("id")
	id := biUserView.Id
	// atoi, err := strconv.Atoi(id)
	if err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	//biUserViewJson := c.Query("biUserView")

	//err = json.Unmarshal([]byte(biUserViewJson), &biUserView)
	biUserView.UpdateTime = utils.GetCurTimeStr()
	if err := biUserService.Update(id, &biUserView); err != nil {
		global.Logger.Error("更新解析上报数据失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err = biUserService.Update(id, &biUserView); err != nil {
		global.Logger.Error("更新持久化失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Get 用id查询BiUser
// @Summary 用id查询BiUser
// @Router /biUser/get [get]
func (biUserApi *BiUserApi) Get(c *gin.Context) {
	id := c.Query("id")
	// atoi, err := strconv.Atoi(id)
	if id == "" {
		// global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	}
	if err, biUserView := biUserService.Get(id); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"biUserView": biUserView}, c)
	}
}

// Find 分页获取BiUser列表
// @Summary 分页获取BiUser列表
// @Router /biUser/find [get]
func (biUserApi *BiUserApi) Find(c *gin.Context) {
	var pageInfo common.PageInfoV2
	// 绑定查询参数到 pageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用 Calculate 方法自动计算 Limit 和 Offset
	pageInfo.Calculate()
	if err := biUserService.Find(&pageInfo); err != nil {
		global.Logger.Error("获取分页信息失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(pageInfo, "获取成功", c)
	}
}
