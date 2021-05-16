package request

import "gin-vue-admin/model"

type AvfOrderSearch struct {
	model.AvfOrder
	PageInfo
}
