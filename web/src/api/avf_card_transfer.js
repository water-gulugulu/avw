import service from '@/utils/request'

// @Tags AvfCardTransfer
// @Summary 创建AvfCardTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCardTransfer true "创建AvfCardTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfCardTransfer/createAvfCardTransfer [post]
export const createAvfCardTransfer = (data) => {
     return service({
         url: "/avfCardTransfer/createAvfCardTransfer",
         method: 'post',
         data
     })
 }


// @Tags AvfCardTransfer
// @Summary 删除AvfCardTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCardTransfer true "删除AvfCardTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfCardTransfer/deleteAvfCardTransfer [delete]
 export const deleteAvfCardTransfer = (data) => {
     return service({
         url: "/avfCardTransfer/deleteAvfCardTransfer",
         method: 'delete',
         data
     })
 }

// @Tags AvfCardTransfer
// @Summary 删除AvfCardTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfCardTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfCardTransfer/deleteAvfCardTransfer [delete]
 export const deleteAvfCardTransferByIds = (data) => {
     return service({
         url: "/avfCardTransfer/deleteAvfCardTransferByIds",
         method: 'delete',
         data
     })
 }

// @Tags AvfCardTransfer
// @Summary 更新AvfCardTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCardTransfer true "更新AvfCardTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfCardTransfer/updateAvfCardTransfer [put]
 export const updateAvfCardTransfer = (data) => {
     return service({
         url: "/avfCardTransfer/updateAvfCardTransfer",
         method: 'put',
         data
     })
 }


// @Tags AvfCardTransfer
// @Summary 用id查询AvfCardTransfer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCardTransfer true "用id查询AvfCardTransfer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfCardTransfer/findAvfCardTransfer [get]
 export const findAvfCardTransfer = (params) => {
     return service({
         url: "/avfCardTransfer/findAvfCardTransfer",
         method: 'get',
         params
     })
 }


// @Tags AvfCardTransfer
// @Summary 分页获取AvfCardTransfer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AvfCardTransfer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfCardTransfer/getAvfCardTransferList [get]
 export const getAvfCardTransferList = (params) => {
     return service({
         url: "/avfCardTransfer/getAvfCardTransferList",
         method: 'get',
         params
     })
 }