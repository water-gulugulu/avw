import service from '@/utils/request'

// @Tags AvfCard
// @Summary 创建AvfCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCard true "创建AvfCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfCard/createAvfCard [post]
export const createAvfCard = (data) => {
     return service({
         url: "/avfCard/createAvfCard",
         method: 'post',
         data
     })
 }


// @Tags AvfCard
// @Summary 删除AvfCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCard true "删除AvfCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfCard/deleteAvfCard [delete]
 export const deleteAvfCard = (data) => {
     return service({
         url: "/avfCard/deleteAvfCard",
         method: 'delete',
         data
     })
 }

// @Tags AvfCard
// @Summary 删除AvfCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AvfCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /avfCard/deleteAvfCard [delete]
 export const deleteAvfCardByIds = (data) => {
     return service({
         url: "/avfCard/deleteAvfCardByIds",
         method: 'delete',
         data
     })
 }

// @Tags AvfCard
// @Summary 更新AvfCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCard true "更新AvfCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /avfCard/updateAvfCard [put]
 export const updateAvfCard = (data) => {
     return service({
         url: "/avfCard/updateAvfCard",
         method: 'put',
         data
     })
 }


// @Tags AvfCard
// @Summary 用id查询AvfCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AvfCard true "用id查询AvfCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /avfCard/findAvfCard [get]
 export const findAvfCard = (params) => {
     return service({
         url: "/avfCard/findAvfCard",
         method: 'get',
         params
     })
 }


// @Tags AvfCard
// @Summary 分页获取AvfCard列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AvfCard列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /avfCard/getAvfCardList [get]
 export const getAvfCardList = (params) => {
     return service({
         url: "/avfCard/getAvfCardList",
         method: 'get',
         params
     })
 }