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

package web_tools

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/blockchian"
	"log"
	"strconv"
	"strings"
	"time"
)

type Manager struct {
	client        *blockchian.ClientManage
	OrderList     []model.AvfOrder
	OrderCardList []model.AvfOrderCard
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
			fmt.Printf("[%s]Reset Timer success!\n", time.Now())
		}
	}
}
func (c *Manager) getOrder() {
	Order := model.AvfOrder{
		Status: 2,
	}

	list, err := Order.FindListByStatus(global.GVA_DB)
	if err != nil {
		fmt.Printf("[%s]Query Order Failed! error:%e\n", time.Now(), err)
		return
	}

	c.OrderList = list

	return
}
func Init() *Manager {
	set := time.Second * 10
	client, err := blockchian.NewClient()
	if err != nil {
		log.Printf("[%s]Failed to client RPC error:%e\n", time.Now(), err)
		return nil
	}
	data := Manager{
		client:        client,
		timer:         time.NewTimer(set),
		OrderList:     make([]model.AvfOrder, 0),
		OrderCardList: make([]model.AvfOrderCard, 0),
		second:        set,
	}
	go data.getOrder()
	go data.timeOut()
	fmt.Printf("[%s]Init Manager success\n", time.Now())
	return &data
}

// 循环读取哈希来改变订单状态
func (c *Manager) LoopOrderStatus() {
	fmt.Printf("order:%s\n", c.OrderList)
	for {
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
				P := 0.001 * 100000000000000000
				Price = int64(P)
			}
			price := strconv.Itoa(int(Price))
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

		time.Sleep(2)
	}
}