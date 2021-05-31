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
	"encoding/json"
	"fmt"
	"gin-vue-admin/api/web"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils/blockchian"
	"gin-vue-admin/utils/rabbitmq"
	"gorm.io/gorm"
	"log"
	"time"
)

type Transfer struct {
	Uid     int     `json:"uid"`
	CardId  int     `json:"card_id"`
	Address string  `json:"address"`
	Type    int     `json:"type"`
	Price   float64 `json:"price"`
	Detail  string  `json:"detail"`
	Hash    string  `json:"hash"`
}

var (
	queueExchange = rabbitmq.QueueExchange{
		QuName: "earnings_queue",
		RtKey:  "#.",
		ExName: "earnings",
		ExType: "topic",
		Dns:    "amqp://root:123123@127.0.0.1:5672",
	}

	Client *blockchian.ClientManage
	err    error
)

func Start() {
	Client, err = blockchian.NewClient()
	if err != nil {
		log.Printf("[%s]client blockchian failed error:%e\n", time.Now(), err)
		return
	}
	forever := make(chan bool)
	t := &TestPro{}
	// for {
	rabbitmq.Recv(queueExchange, t, 10)
	// }
	<-forever
	// return &manager
}

type TestPro struct {
	msgContent string
}

// 实现发送者
func (t *TestPro) FailAction(dataByte []byte) error {
	rabbitMQ := model.AvfRabbitmqError{
		Message:    string(dataByte),
		CreateTime: time.Now(),
	}
	_ = rabbitMQ.Create(global.GVA_DB)
	// fmt.Println(dataByte)
	return nil
}

// 实现接收者
func (t *TestPro) Consumer(dataByte []byte) error {
	// fmt.Printf("data:%s\n", dataByte)
	stu := &web.UserData{}
	if err := json.Unmarshal(dataByte, &stu); err != nil {
		return err
	}
	hash, err := Client.TransferToAddress(stu.WalletAddress, 0.0001)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("[%s]用户地址：%s,发放每日挖矿收益失败，金额：%v,:%e\n", time.Now(), stu.WalletAddress, stu.Money+stu.Direct, err))
		return err
	}
	DB := global.GVA_DB
	return DB.Transaction(func(tx *gorm.DB) error {
		for _, item := range stu.UserBill {
			UserBill := model.AvfUserBill{
				GVA_MODEL:  global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
				Uid:        item.Uid,
				CardId:     item.CardId,
				Pid:        item.Pid,
				Address:    item.Address,
				Type:       item.Type,
				Money:      item.Money,
				Payment:    item.Payment,
				PayType:    item.PayType,
				Detail:     item.Detail,
				TxHash:     hash,
				CreateTime: item.CreateTime,
			}
			if err = UserBill.Create(DB); err != nil {
				log.Printf("[%s]用户地址：%s,发放每日挖矿收益账单保存失败，金额：%v,:%e\n", time.Now(), item.Address, item.Money, err)
				return err
			}
		}
		return nil
	})
}

//
// func (c *Manager) Transfer() {
// 	CardRecord := model.AvfOrderCard{
// 		Status: 1,
// 	}
// 	DB := global.GVA_DB
// 	list, err2 := CardRecord.GetListByMining(DB)
// 	if err2 != nil {
// 		log.Printf("[%s]query card list failed error:%e\n", time.Now(), err2)
// 		return
// 	}
// 	l := make(map[int]*model.AvfOrderCard, 0)
// 	var allStar int
// 	for key, item := range list {
// 		l[key] = item
// 		allStar = allStar + item.Star
// 	}
// 	starExchange := global.GVA_CONFIG.CollectionAddress.MaxExchange
// 	oneStarExchange := web_tools.IntToFloat(starExchange) / web_tools.IntToFloat(allStar)
// 	// fmt.Printf("avw:%s\n", oneStarExchange)
// 	c.LoopList = l
// 	if len(c.LoopList) == 0 {
// 		log.Println("列表空的")
// 		return
// 	}
// 	var err error
// 	var Parent *model.AvfUser
// 	var UserBill model.AvfUserBill
// 	var hash, parentHash string
// 	Direct := global.GVA_CONFIG.CollectionAddress.Direct
// 	// e, _ := oneStarExchange.Float64()
// 	d := web_tools.IntToFloat(Direct)
//
// 	UserList, err2 := new(model.AvfUser).GetListAll(DB)
// 	if err2 != nil {
// 		log.Printf("[%s]query user list failed error:%e\n", time.Now(), err2)
// 		return
// 	}
// 	UserMap := make(map[string]*model.AvfUser, 0)
// 	for _, item := range UserList {
// 		UserMap[item.WalletAddress] = item
// 	}
// 	for key, item := range c.LoopList {
// 		Price := oneStarExchange * web_tools.IntToFloat(item.Star)
// 		// fmt.Printf("price:%v\n", Price)
// 		ParentPrice := Price * d / 100
// 		ParentPrice = web_tools.FormatFloat(ParentPrice, 4)
//
// 		// fmt.Printf("price:%v,parentPrice:%v\n", Price, ParentPrice)
// 		// return
// 		hash, err = c.BlockChain.TransferToAddress(item.User.WalletAddress, Price)
// 		if err != nil {
// 			log.Printf("[%s]用户地址：%s,发放每日挖矿收益失败，金额：%v,:%e\n", time.Now(), item.User.WalletAddress, Price, err)
// 			continue
// 		}
// 		UserBill = model.AvfUserBill{
// 			GVA_MODEL: global.GVA_MODEL{
// 				UpdatedAt: time.Now(),
// 			},
// 			Uid:        item.Uid,
// 			CardId:     int(item.ID),
// 			Address:    item.User.WalletAddress,
// 			Pid:        item.CardId,
// 			Type:       1,
// 			Money:      Price,
// 			Payment:    1,
// 			PayType:    1,
// 			Detail:     fmt.Sprintf("每日挖矿收益：%v", Price),
// 			TxHash:     hash,
// 			CreateTime: int(time.Now().Unix()),
// 		}
// 		if err = UserBill.Create(DB); err != nil {
// 			log.Printf("[%s]用户地址：%s,发放每日挖矿收益账单保存失败，金额：%v,:%e\n", time.Now(), item.User.WalletAddress, Price, err)
// 			return
// 		}
// 		Parent = UserMap[item.User.Pid]
// 		if len(item.User.Pid) != 0 && Parent != nil {
// 			parentHash, err = c.BlockChain.TransferToAddress(item.User.Pid, ParentPrice)
// 			if err != nil {
// 				log.Printf("[%s]用户地址：%s,发放每日挖矿直推收益失败，金额：%v,:%e\n", time.Now(), item.User.Pid, ParentPrice, err)
// 				continue
// 			}
//
// 			UserBill = model.AvfUserBill{
// 				GVA_MODEL: global.GVA_MODEL{
// 					UpdatedAt: time.Now(),
// 				},
// 				Uid:        int(Parent.ID),
// 				CardId:     int(item.ID),
// 				Address:    Parent.WalletAddress,
// 				Pid:        item.CardId,
// 				Type:       5,
// 				Money:      ParentPrice,
// 				Payment:    1,
// 				PayType:    1,
// 				Detail:     fmt.Sprintf("直推下级：%v,挖矿直推收益：%v", item.User.WalletAddress, ParentPrice),
// 				TxHash:     parentHash,
// 				CreateTime: int(time.Now().Unix()),
// 			}
// 			if err = UserBill.Create(DB); err != nil {
// 				log.Printf("[%s]用户地址：%s,发放每日挖矿收益账单保存失败，金额：%v,:%e\n", time.Now(), item.User.WalletAddress, Price, err2)
// 				return
// 			}
// 		}
// 		delete(c.LoopList, key)
// 	}
// }
