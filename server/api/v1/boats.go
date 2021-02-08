package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Boats
// @Summary 创建Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "创建Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /boats/createBoats [post]
func CreateBoats(c *gin.Context) {
	var boats model.Boats
	_ = c.ShouldBindJSON(&boats)
	if err := service.CreateBoats(boats); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Boats
// @Summary 删除Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "删除Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /boats/deleteBoats [delete]
func DeleteBoats(c *gin.Context) {
	var boats model.Boats
	_ = c.ShouldBindJSON(&boats)
	if err := service.DeleteBoats(boats); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Boats
// @Summary 批量删除Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /boats/deleteBoatsByIds [delete]
func DeleteBoatsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteBoatsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Boats
// @Summary 更新Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "更新Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /boats/updateBoats [put]
func UpdateBoats(c *gin.Context) {
	var boats model.Boats
	_ = c.ShouldBindJSON(&boats)
	if err := service.UpdateBoats(boats); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Boats
// @Summary 用id查询Boats
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Boats true "用id查询Boats"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /boats/findBoats [get]
func FindBoats(c *gin.Context) {
	var boats model.Boats
	_ = c.ShouldBindQuery(&boats)
	if err, reboats := service.GetBoats(boats.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reboats": reboats}, c)
	}
}

// @Tags Boats
// @Summary 分页获取Boats列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BoatsSearch true "分页获取Boats列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /boats/getBoatsList [get]
func GetBoatsList(c *gin.Context) {
	var pageInfo request.BoatsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetBoatsInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}