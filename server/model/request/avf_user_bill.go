package request

import "gin-vue-admin/model"

type AvfUserBillSearch struct {
	model.AvfUserBill
	PageInfo
}
