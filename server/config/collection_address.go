// @File  : collection-address.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/15
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

package config

type CollectionAddress struct {
	Address    string `json:"address"`    // 收款地址
	Fees       string `json:"fees"`       // 转卡手续费百分比
	Proportion string `json:"proportion"` // 卡牌卖出原价比例
	Exchange   string `json:"exchange"`   // 算力每日释放avw比例
	Direct     string `json:"direct"`     // 直推奖金比例
	Debug      string `json:"debug"`
}
