// @File  : struct.go
// @Author: JunLong.Liao&此处不应有BUG!
// @Date  : 2021/5/11
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

package tools

import "math/big"

type Address struct {
	Address string `json:"address"`
	Crypto  Crypto `json:"crypto"`
	ID      string `json:"id"`
	Version int    `json:"version"`
}
type Cipherparams struct {
	Iv string `json:"iv"`
}
type Kdfparams struct {
	Dklen int    `json:"dklen"`
	N     int    `json:"n"`
	P     int    `json:"p"`
	R     int    `json:"r"`
	Salt  string `json:"salt"`
}
type Crypto struct {
	Cipher       string       `json:"cipher"`
	Ciphertext   string       `json:"ciphertext"`
	Cipherparams Cipherparams `json:"cipherparams"`
	Kdf          string       `json:"kdf"`
	Kdfparams    Kdfparams    `json:"kdfparams"`
	Mac          string       `json:"mac"`
}

// TransactionResponse 查询交易返回结果
type TransactionResponse struct {
	TxHash     string   `json:"tx_hash"`
	Block      *big.Int `json:"block"`
	From       string   `json:"from"`
	To         string   `json:"to"`
	GasPrice   *big.Int `json:"gas_price"`
	Value      *big.Int `json:"value"`
	Gas        uint64   `json:"gas"`
	Nonce      uint64   `json:"nonce"`
	Data       []byte   `json:"data"`
	CheckNonce bool     `json:"check_nonce"`
	Status     uint64   `json:"status"`
}
