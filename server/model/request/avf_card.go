package request

import "gin-vue-admin/model"

type AvfCardSearch struct{
    model.AvfCard
    PageInfo
}