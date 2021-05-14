import service from '@/utils/request'

// @Tags AvfUser
// @Summary 创建AvfUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUser true "创建AvfUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfUser/createAvfUser [post]
export const createAvfUser = (data) => {
     return service({
         url: "/avfUser/createAvfUser",
         method: 'post',
         data
     })
 }


// @Tags AvfUser
// @Summary 删除AvfUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUser true "删除AvfUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfUser/deleteAvfUser [delete]
 export const deleteAvfUser = (data) => {
     return service({
         url: "/avfUser/deleteAvfUser",
         method: 'delete',
         data
     })
 }

// @Tags AvfUser
// @Summary 删除AvfUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfUser/deleteAvfUser [delete]
 export const deleteAvfUserByIds = (data) => {
     return service({
         url: "/avfUser/deleteAvfUserByIds",
         method: 'delete',
         data
     })
 }

// @Tags AvfUser
// @Summary 更新AvfUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUser true "更新AvfUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfUser/updateAvfUser [put]
 export const updateAvfUser = (data) => {
     return service({
         url: "/avfUser/updateAvfUser",
         method: 'put',
         data
     })
 }


// @Tags AvfUser
// @Summary 用id查询AvfUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUser true "用id查询AvfUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfUser/findAvfUser [get]
 export const findAvfUser = (params) => {
     return service({
         url: "/avfUser/findAvfUser",
         method: 'get',
         params
     })
 }


// @Tags AvfUser
// @Summary 分页获取AvfUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AvfUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfUser/getAvfUserList [get]
 export const getAvfUserList = (params) => {
     return service({
         url: "/avfUser/getAvfUserList",
         method: 'get',
         params
     })
 }