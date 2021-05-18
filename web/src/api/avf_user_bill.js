import service from '@/utils/request'

// @Tags AvfUserBill
// @Summary 创建AvfUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUserBill true "创建AvfUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfUserBill/createAvfUserBill [post]
export const createAvfUserBill = (data) => {
     return service({
         url: "/avfUserBill/createAvfUserBill",
         method: 'post',
         data
     })
 }


// @Tags AvfUserBill
// @Summary 删除AvfUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUserBill true "删除AvfUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfUserBill/deleteAvfUserBill [delete]
 export const deleteAvfUserBill = (data) => {
     return service({
         url: "/avfUserBill/deleteAvfUserBill",
         method: 'delete',
         data
     })
 }

// @Tags AvfUserBill
// @Summary 删除AvfUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfUserBill/deleteAvfUserBill [delete]
 export const deleteAvfUserBillByIds = (data) => {
     return service({
         url: "/avfUserBill/deleteAvfUserBillByIds",
         method: 'delete',
         data
     })
 }

// @Tags AvfUserBill
// @Summary 更新AvfUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUserBill true "更新AvfUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfUserBill/updateAvfUserBill [put]
 export const updateAvfUserBill = (data) => {
     return service({
         url: "/avfUserBill/updateAvfUserBill",
         method: 'put',
         data
     })
 }


// @Tags AvfUserBill
// @Summary 用id查询AvfUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfUserBill true "用id查询AvfUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfUserBill/findAvfUserBill [get]
 export const findAvfUserBill = (params) => {
     return service({
         url: "/avfUserBill/findAvfUserBill",
         method: 'get',
         params
     })
 }


// @Tags AvfUserBill
// @Summary 分页获取AvfUserBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AvfUserBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfUserBill/getAvfUserBillList [get]
 export const getAvfUserBillList = (params) => {
     return service({
         url: "/avfUserBill/getAvfUserBillList",
         method: 'get',
         params
     })
 }