import service from '@/utils/request'

// @Tags AvfOrder
// @Summary 创建AvfOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrder true "创建AvfOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfOrder/createAvfOrder [post]
export const createAvfOrder = (data) => {
     return service({
         url: "/avfOrder/createAvfOrder",
         method: 'post',
         data
     })
 }


// @Tags AvfOrder
// @Summary 删除AvfOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrder true "删除AvfOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfOrder/deleteAvfOrder [delete]
 export const deleteAvfOrder = (data) => {
     return service({
         url: "/avfOrder/deleteAvfOrder",
         method: 'delete',
         data
     })
 }

// @Tags AvfOrder
// @Summary 删除AvfOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfOrder/deleteAvfOrder [delete]
 export const deleteAvfOrderByIds = (data) => {
     return service({
         url: "/avfOrder/deleteAvfOrderByIds",
         method: 'delete',
         data
     })
 }

// @Tags AvfOrder
// @Summary 更新AvfOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrder true "更新AvfOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfOrder/updateAvfOrder [put]
 export const updateAvfOrder = (data) => {
     return service({
         url: "/avfOrder/updateAvfOrder",
         method: 'put',
         data
     })
 }


// @Tags AvfOrder
// @Summary 用id查询AvfOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrder true "用id查询AvfOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfOrder/findAvfOrder [get]
 export const findAvfOrder = (params) => {
     return service({
         url: "/avfOrder/findAvfOrder",
         method: 'get',
         params
     })
 }


// @Tags AvfOrder
// @Summary 分页获取AvfOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AvfOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfOrder/getAvfOrderList [get]
 export const getAvfOrderList = (params) => {
     return service({
         url: "/avfOrder/getAvfOrderList",
         method: 'get',
         params
     })
 }