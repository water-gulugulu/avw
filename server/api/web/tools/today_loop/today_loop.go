// @File  : today_loop.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/21
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

package today_loop

import (
	"fmt"
	web_tools "gin-vue-admin/api/web/tools"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/blockchian"
	"log"
	"strconv"
	"time"
)

type Manager struct {
	List       map[int]Transfer
	LoopList   map[int]*model.AvfOrderCard
	BlockChain *blockchian.ClientManage
	InChan     chan Transfer
}

type Transfer struct {
	Uid     int     `json:"uid"`
	CardId  int     `json:"card_id"`
	Address string  `json:"address"`
	Type    int     `json:"type"`
	Price   float64 `json:"price"`
	Detail  string  `json:"detail"`
	Hash    string  `json:"hash"`
}

func Start() *Manager {
	client, err := blockchian.NewClient()
	if err != nil {
		log.Printf("[%s]client blockchian failed error:%e\n", time.Now(), err)
		return nil
	}
	manager := Manager{
		List:       make(map[int]Transfer, 0),
		LoopList:   make(map[int]*model.AvfOrderCard, 0),
		InChan:     make(chan Transfer, 128),
		BlockChain: client,
	}

	// go manager.GetLoopList()
	return &manager
}

func (c *Manager) GetLoopList() {
	CardRecord := model.AvfOrderCard{
		Status: 1,
	}
	DB := global.GVA_DB
	list, err := CardRecord.GetListByMining(DB)
	if err != nil {
		log.Printf("[%s]query card list failed error:%e\n", time.Now(), err)
		return
	}
	l := make(map[int]*model.AvfOrderCard, 0)
	for key, item := range list {
		l[key] = item
	}
	c.LoopList = l
}

func (c *Manager) Transfer() {
	CardRecord := model.AvfOrderCard{
		Status: 1,
	}
	DB := global.GVA_DB
	list, err2 := CardRecord.GetListByMining(DB)
	if err2 != nil {
		log.Printf("[%s]query card list failed error:%e\n", time.Now(), err2)
		return
	}
	l := make(map[int]*model.AvfOrderCard, 0)
	for key, item := range list {
		l[key] = item
	}
	c.LoopList = l
	if len(c.LoopList) == 0 {
		fmt.Println("列表空的")
		return
	}
	var err error
	var Parent *model.AvfUser
	var UserBill model.AvfUserBill
	var hash, parentHash string
	Exchange := global.GVA_CONFIG.CollectionAddress.Exchange
	Direct := global.GVA_CONFIG.CollectionAddress.Direct

	e, _ := strconv.ParseFloat(Exchange, 64)
	d, _ := strconv.ParseFloat(Direct, 64)

	UserList, err2 := new(model.AvfUser).GetListAll(DB)
	if err2 != nil {
		log.Printf("[%s]query user list failed error:%e\n", time.Now(), err2)
		return
	}
	UserMap := make(map[string]*model.AvfUser, 0)
	for _, item := range UserList {
		UserMap[item.WalletAddress] = item
	}
	for key, item := range c.LoopList {
		Price := e * float64(item.Star)
		ParentPrice := d * Price / 100
		Price = web_tools.FormatFloat(Price, 4)
		ParentPrice = web_tools.FormatFloat(ParentPrice, 4)

		hash, err = c.BlockChain.TransferToAddress(item.User.WalletAddress, Price)
		if err != nil {
			log.Printf("[%s]用户地址：%s,发放每日挖矿收益失败，金额：%v,:%e\n", time.Now(), item.User.WalletAddress, Price, err)
			continue
		}
		UserBill = model.AvfUserBill{
			GVA_MODEL: global.GVA_MODEL{
				UpdatedAt: time.Now(),
			},
			Uid:        item.Uid,
			CardId:     int(item.ID),
			Address:    item.User.WalletAddress,
			Pid:        item.CardId,
			Type:       1,
			Money:      Price,
			Payment:    1,
			PayType:    1,
			Detail:     fmt.Sprintf("每日挖矿收益：%v", Price),
			TxHash:     hash,
			CreateTime: int(time.Now().Unix()),
		}
		if err = UserBill.Create(DB); err != nil {
			log.Printf("[%s]用户地址：%s,发放每日挖矿收益账单保存失败，金额：%v,:%e\n", time.Now(), item.User.WalletAddress, Price, err)
			return
		}
		Parent = UserMap[item.User.Pid]
		if len(item.User.Pid) != 0 && Parent != nil {
			parentHash, err = c.BlockChain.TransferToAddress(item.User.Pid, ParentPrice)
			if err != nil {
				log.Printf("[%s]用户地址：%s,发放每日挖矿直推收益失败，金额：%v,:%e\n", time.Now(), item.User.Pid, ParentPrice, err)
				continue
			}

			UserBill = model.AvfUserBill{
				GVA_MODEL: global.GVA_MODEL{
					UpdatedAt: time.Now(),
				},
				Uid:        int(Parent.ID),
				CardId:     int(item.ID),
				Address:    Parent.WalletAddress,
				Pid:        item.CardId,
				Type:       5,
				Money:      ParentPrice,
				Payment:    1,
				PayType:    1,
				Detail:     fmt.Sprintf("直推下级：%v,挖矿直推收益：%v", item.User.WalletAddress, ParentPrice),
				TxHash:     parentHash,
				CreateTime: int(time.Now().Unix()),
			}
			if err = UserBill.Create(DB); err != nil {
				log.Printf("[%s]用户地址：%s,发放每日挖矿收益账单保存失败，金额：%v,:%e\n", time.Now(), item.User.WalletAddress, Price, err2)
				return
			}
		}
		delete(c.LoopList, key)
	}
}
