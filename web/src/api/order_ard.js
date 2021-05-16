import service from '@/utils/request'

// @Tags AvfOrderCard
// @Summary 创建AvfOrderCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrderCard true "创建AvfOrderCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfOrderCard/createAvfOrderCard [post]
export const createAvfOrderCard = (data) => {
     return service({
         url: "/avfOrderCard/createAvfOrderCard",
         method: 'post',
         data
     })
 }


// @Tags AvfOrderCard
// @Summary 删除AvfOrderCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrderCard true "删除AvfOrderCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfOrderCard/deleteAvfOrderCard [delete]
 export const deleteAvfOrderCard = (data) => {
     return service({
         url: "/avfOrderCard/deleteAvfOrderCard",
         method: 'delete',
         data
     })
 }

// @Tags AvfOrderCard
// @Summary 删除AvfOrderCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfOrderCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfOrderCard/deleteAvfOrderCard [delete]
 export const deleteAvfOrderCardByIds = (data) => {
     return service({
         url: "/avfOrderCard/deleteAvfOrderCardByIds",
         method: 'delete',
         data
     })
 }

// @Tags AvfOrderCard
// @Summary 更新AvfOrderCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrderCard true "更新AvfOrderCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfOrderCard/updateAvfOrderCard [put]
 export const updateAvfOrderCard = (data) => {
     return service({
         url: "/avfOrderCard/updateAvfOrderCard",
         method: 'put',
         data
     })
 }


// @Tags AvfOrderCard
// @Summary 用id查询AvfOrderCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfOrderCard true "用id查询AvfOrderCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfOrderCard/findAvfOrderCard [get]
 export const findAvfOrderCard = (params) => {
     return service({
         url: "/avfOrderCard/findAvfOrderCard",
         method: 'get',
         params
     })
 }


// @Tags AvfOrderCard
// @Summary 分页获取AvfOrderCard列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AvfOrderCard列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfOrderCard/getAvfOrderCardList [get]
 export const getAvfOrderCardList = (params) => {
     return service({
         url: "/avfOrderCard/getAvfOrderCardList",
         method: 'get',
         params
     })
 }