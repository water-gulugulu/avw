// @File  : loop_order.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/17
// @slogan: 又是不想写代码的一天，神兽保佑，代码无BUG！
//         ┏┓      ┏┓
//        ┏┛┻━━━━━━┛┻┓
//        ┃     ღ    ┃
//        ┃  ┳┛   ┗┳ ┃
//        ┃     ┻    ┃
//        ┗━┓      ┏━┛
//          ┃      ┗━━━┓
//          ┃ 神兽咆哮!  ┣┓
//          ┃         ┏┛
//          ┗┓┓┏━━━┳┓┏┛
//           ┃┫┫   ┃┫┫
//           ┗┻┛   ┗┻┛

package order_loop

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/blockchian"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
	"time"
)

type Manager struct {
	client        *blockchian.ClientManage
	OrderList     []*model.AvfOrder
	OrderPayList  []*model.AvfCardTransfer
	OrderFeesList []*model.AvfCardTransfer
	timer         *time.Timer
	second        time.Duration
}

func (c *Manager) timeOut() {
	for {
		select {
		case <-c.timer.C:
			c.timer.Reset(c.second)
			if c.client == nil {
				client, _ := blockchian.NewClient()
				if client != nil {
					c.client = client
				}
			}
			go c.getOrder()
			go c.ChangeOrderStatus()
			go c.LoopOrderStatus()
			go c.LoopFeesOrder()
			go c.LoopPayOrder()
			fmt.Printf("[%s]Reset Timer success!\n", time.Now())
		}
	}
}
func (c *Manager) getOrder() {
	Order := model.AvfOrder{
		Status: 2,
	}
	DB := global.GVA_DB
	list, err := Order.FindListByStatus(DB)
	if err != nil {
		fmt.Printf("[%s]Query Order Failed! error:%e\n", time.Now(), err)
		return
	}
	FeesOrder := model.AvfCardTransfer{
		Status: 2,
	}
	PayOrder := model.AvfCardTransfer{
		Status: 5,
	}

	feesList, err2 := FeesOrder.GetByStatus(DB)

	if err2 != nil {
		fmt.Printf("[%s]Query fees_order Failed! error:%e\n", time.Now(), err2)
		return
	}
	payList, err3 := PayOrder.GetByStatus(DB)

	if err3 != nil {
		fmt.Printf("[%s]Query pay_order Failed! error:%e\n", time.Now(), err3)
		return
	}
	c.OrderList = list
	c.OrderFeesList = feesList
	c.OrderPayList = payList

	return
}
func Init() *Manager {
	set := time.Second * 15
	client, err := blockchian.NewClient()
	if err != nil {
		log.Printf("[%s]Failed to client RPC error:%e\n", time.Now(), err)
		return nil
	}
	data := Manager{
		client:        client,
		timer:         time.NewTimer(set),
		OrderList:     make([]*model.AvfOrder, 0),
		OrderPayList:  make([]*model.AvfCardTransfer, 0),
		OrderFeesList: make([]*model.AvfCardTransfer, 0),
		second:        set,
	}
	go data.getOrder()
	go data.timeOut()

	fmt.Printf("[%s]Init Manager success\n", time.Now())
	return &data
}

// 循环读取哈希来改变订单状态
func (c *Manager) LoopOrderStatus() {
	// for {
	for _, item := range c.OrderList {
		res, err2 := c.client.QueryTransactionByTxHash(item.TxHash)
		if err2 != nil {
			log.Printf("[%s]Failed to query transaction error:%e\n", time.Now(), err2)
			continue
		}
		if res.Status != 1 {
			log.Printf("[%s]Failed to status not 1\n", time.Now())
			continue
		}
		res.From = strings.ToUpper(res.From)
		item.From = strings.ToUpper(item.From)
		if res.From != item.From {
			log.Printf("[%s]Failed to form:%s orderForm:%s\n", time.Now(), res.From, item.From)
			continue
		}
		Price := item.Price * 100000000000000000

		if global.GVA_CONFIG.CollectionAddress.Debug == "1" {
			Price = 100000000000000
			// Price = int64(P)
		}
		// orderPrice:10000000000000,
		// tx_price:  100000000000000
		price := strconv.Itoa(int(Price))
		// fmt.Printf("orderPrice:%v,tx_price:%v\n", price, res.Value.String())
		if price != res.Value.String() {
			log.Printf("[%s]Failed to money not same money:%v,%v \n", time.Now(), Price, res.Value)
			continue
		}

		res.To = strings.ToUpper(res.To)
		item.To = strings.ToUpper(global.GVA_CONFIG.CollectionAddress.Address)
		if res.To != item.To {
			log.Printf("[%s]Failed to to:%s orderTo:%s\n", time.Now(), res.To, item.To)
			continue
		}

		Order := model.AvfOrder{
			GVA_MODEL: global.GVA_MODEL{
				ID:        item.ID,
				UpdatedAt: time.Now(),
			},
			TxHash:   item.TxHash,
			Status:   3,
			PayTime:  int(time.Now().Unix()),
			Block:    res.Block.String(),
			Gas:      strconv.Itoa(int(res.Gas)),
			GasPrice: res.GasPrice.String(),
			From:     res.From,
			To:       res.To,
		}
		if err := Order.UpdateOrder(global.GVA_DB); err != nil {
			log.Printf("[%s]Failed to update Order error:%e\n", time.Now(), err)
			continue
		}

		Log := model.AvfTransactionLog{
			OrderId:    int(item.ID),
			Block:      res.Block.String(),
			TxHash:     item.TxHash,
			Form:       res.From,
			To:         res.To,
			Gas:        strconv.Itoa(int(res.Gas)),
			GasPrice:   res.GasPrice.String(),
			Value:      res.Value.String(),
			Nonce:      string(res.Nonce),
			Data:       string(res.Data),
			Status:     int64(res.Status),
			CreateDate: time.Now(),
			CreateTime: time.Now().Unix(),
			Type:       1,
		}
		if err := Log.CreateLog(global.GVA_DB); err != nil {
			log.Printf("[%s]Failed to update Order error:%e,Log:%s\n", time.Now(), err, Log)
		}
	}
	// }
}

// 循环读取哈希来改变手续费订单
func (c *Manager) LoopFeesOrder() {
	// for {
	for _, item := range c.OrderFeesList {
		res, err2 := c.client.QueryTransactionByTxHash(item.FeesHash)
		if err2 != nil {
			log.Printf("[%s]Failed to query fees transaction error:%e\n", time.Now(), err2)
			continue
		}
		if res.Status != 1 {
			log.Printf("[%s]Failed to fees status not 1\n", time.Now())
			continue
		}
		res.From = strings.ToUpper(res.From)
		item.From = strings.ToUpper(item.From)
		if res.From != item.From {
			log.Printf("[%s]Failed to fees form:%s orderForm:%s\n", time.Now(), res.From, item.From)
			continue
		}
		Price := item.Price * 100000000000000000

		if global.GVA_CONFIG.CollectionAddress.Debug == "1" {
			P := 0.0001 * 100000000000000000
			Price = int(P)
		}
		price := strconv.Itoa(Price)
		if price != res.Value.String() {
			log.Printf("[%s]Failed to fees money not same money:%v,%v \n", time.Now(), Price, res.Value)
			continue
		}

		res.To = strings.ToUpper(res.To)
		item.To = strings.ToUpper(global.GVA_CONFIG.CollectionAddress.Address)
		if res.To != item.To {
			log.Printf("[%s]Failed to fees to:%s orderTo:%s\n", time.Now(), res.To, item.To)
			continue
		}

		Order := model.AvfCardTransfer{
			GVA_MODEL: global.GVA_MODEL{
				ID:        item.ID,
				UpdatedAt: time.Now(),
			},
			FeesHash: item.TxHash,
			Status:   3,
			Block:    res.Block.String(),
			From:     res.From,
			System:   res.To,
		}
		if err := Order.Update(global.GVA_DB); err != nil {
			log.Printf("[%s]Failed to fees update Order error:%e\n", time.Now(), err)
			continue
		}

		Log := model.AvfTransactionLog{
			OrderId:    int(item.ID),
			Block:      res.Block.String(),
			TxHash:     item.FeesHash,
			Form:       res.From,
			To:         res.To,
			Gas:        strconv.Itoa(int(res.Gas)),
			GasPrice:   res.GasPrice.String(),
			Value:      res.Value.String(),
			Nonce:      string(res.Nonce),
			Data:       string(res.Data),
			Status:     int64(res.Status),
			CreateDate: time.Now(),
			CreateTime: time.Now().Unix(),
			Type:       2,
		}
		if err := Log.CreateLog(global.GVA_DB); err != nil {
			log.Printf("[%s]Failed to fees update Order error:%e,Log:%s\n", time.Now(), err, Log)
		}
	}
	// }
}

// 循环读取哈希来改变购买卡牌订单
func (c *Manager) LoopPayOrder() {
	// for {
	for _, item := range c.OrderPayList {
		res, err2 := c.client.QueryTransactionByTxHash(item.TxHash)
		if err2 != nil {
			log.Printf("[%s]Failed to pay transaction error:%e\n", time.Now(), err2)
			continue
		}
		if res.Status != 1 {
			log.Printf("[%s]Failed to pay status not 1\n", time.Now())
			continue
		}
		res.From = strings.ToUpper(res.From)
		item.From = strings.ToUpper(item.From)
		res.To = strings.ToUpper(res.To)
		item.To = strings.ToUpper(item.To)
		if res.From != item.To {
			log.Printf("[%s]Failed to pay form:%s orderForm:%s\n", time.Now(), res.From, item.From)
			continue
		}
		Price := item.Price * 100000000000000000

		if global.GVA_CONFIG.CollectionAddress.Debug == "1" {
			P := 0.0001 * 100000000000000000
			Price = int(P)
		}
		price := strconv.Itoa(Price)
		if price != res.Value.String() {
			log.Printf("[%s]Failed to pay money not same money:%v,%v \n", time.Now(), Price, res.Value)
			continue
		}

		if res.To != item.From {
			log.Printf("[%s]Failed to pay to:%s orderTo:%s\n", time.Now(), res.To, item.To)
			continue
		}
		_ = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			Order := model.AvfCardTransfer{
				GVA_MODEL: global.GVA_MODEL{
					ID:        item.ID,
					UpdatedAt: time.Now(),
				},
				TxHash: item.TxHash,
				Status: 6,
				Block:  res.Block.String(),
				From:   res.To,
				To:     res.From,
			}
			if err := Order.Update(tx); err != nil {
				log.Printf("[%s]Failed to pay update Order error:%e\n", time.Now(), err)
				return err
			}
			OrderRecord := model.AvfOrderCard{
				GVA_MODEL: global.GVA_MODEL{
					ID:        uint(item.RecordId),
					UpdatedAt: time.Now(),
				},
				Uid:        item.BuyId,
				Status:     1,
				GiveType:   2,
				UpdateTime: int(time.Now().Unix()),
			}
			if err := OrderRecord.Update(tx); err != nil {
				return err
			}

			Log := model.AvfTransactionLog{
				OrderId:    int(item.ID),
				Block:      res.Block.String(),
				TxHash:     item.TxHash,
				Form:       res.From,
				To:         res.To,
				Gas:        strconv.Itoa(int(res.Gas)),
				GasPrice:   res.GasPrice.String(),
				Value:      res.Value.String(),
				Nonce:      string(res.Nonce),
				Data:       string(res.Data),
				Status:     int64(res.Status),
				CreateDate: time.Now(),
				CreateTime: time.Now().Unix(),
				Type:       3,
			}
			if err := Log.CreateLog(tx); err != nil {
				log.Printf("[%s]Failed to pay update Order error:%e,Log:%s\n", time.Now(), err, Log)
				return err
			}
			return nil
		})
	}
	// }
}

// 循环取消点击购买未付款的订单
func (c *Manager) ChangeOrderStatus() {
	PayOrder := model.AvfCardTransfer{
		Status:     4,
		ExpireTime: int(time.Now().Unix()),
	}

	_ = PayOrder.ChangeStatusByExpireTime(global.GVA_DB)
	return
}
