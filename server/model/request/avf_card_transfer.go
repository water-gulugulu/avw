package request

import "gin-vue-admin/model"

type AvfCardTransferSearch struct {
	model.AvfCardTransfer
	PageInfo
}
