import service from '@/utils/request'

// @Tags AvfEventLog
// @Summary 创建AvfEventLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfEventLog true "创建AvfEventLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfEventLog/createAvfEventLog [post]
export const createAvfEventLog = (data) => {
     return service({
         url: "/avfEventLog/createAvfEventLog",
         method: 'post',
         data
     })
 }


// @Tags AvfEventLog
// @Summary 删除AvfEventLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfEventLog true "删除AvfEventLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfEventLog/deleteAvfEventLog [delete]
 export const deleteAvfEventLog = (data) => {
     return service({
         url: "/avfEventLog/deleteAvfEventLog",
         method: 'delete',
         data
     })
 }

// @Tags AvfEventLog
// @Summary 删除AvfEventLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfEventLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfEventLog/deleteAvfEventLog [delete]
 export const deleteAvfEventLogByIds = (data) => {
     return service({
         url: "/avfEventLog/deleteAvfEventLogByIds",
         method: 'delete',
         data
     })
 }

// @Tags AvfEventLog
// @Summary 更新AvfEventLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfEventLog true "更新AvfEventLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfEventLog/updateAvfEventLog [put]
 export const updateAvfEventLog = (data) => {
     return service({
         url: "/avfEventLog/updateAvfEventLog",
         method: 'put',
         data
     })
 }


// @Tags AvfEventLog
// @Summary 用id查询AvfEventLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfEventLog true "用id查询AvfEventLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfEventLog/findAvfEventLog [get]
 export const findAvfEventLog = (params) => {
     return service({
         url: "/avfEventLog/findAvfEventLog",
         method: 'get',
         params
     })
 }


// @Tags AvfEventLog
// @Summary 分页获取AvfEventLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AvfEventLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfEventLog/getAvfEventLogList [get]
 export const getAvfEventLogList = (params) => {
     return service({
         url: "/avfEventLog/getAvfEventLogList",
         method: 'get',
         params
     })
 }