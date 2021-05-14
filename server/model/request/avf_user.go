package request

import "gin-vue-admin/model"

type AvfUserSearch struct{
    model.AvfUser
    PageInfo
}