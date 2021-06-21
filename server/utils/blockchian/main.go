// @File  : main.go
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

package blockchian

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	Token "gin-vue-admin/utils/blockchian/token"
	"gin-vue-admin/utils/blockchian/tools"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math"
	"math/big"
	"strings"
	"time"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

const (
	// contract = "0x21fd0FBE5Fb40B9A86FF21f223dCbCB2A308c3E5" // 旧的
	// contract = "0x3c94f2bb9a35e38a827feae443bc63d5a80a409f"                                                               // 女优币 1万亿
	contract = "0x8429937eaD794f4B82009B4aCf18Db52E2171235"                                                               // 女优币 1万亿
	key      = "./utils/blockchian/wallets/UTC--2021-05-11T06-48-26.264188000Z--8f2b1cea616b837b74ae5b5e31054a36cd2fd380" // 火币链 0x8f2b1CeA616b837b74Ae5B5E31054A36cd2FD380
	// key = "./wallets/UTC--2021-05-10T16-31-35.638486000Z--d7940959ec892652f2042fcb0f9feef3e498e724" // 私链     0xd7940959ec892652f2042fcb0f9feef3e498e724
	RawUrl = "https://http-mainnet-node.huobichain.com"
)

var Now = time.Now().Format("2006-01-02 15:04:05")

type ClientManage struct {
	rpcConn *rpc.Client
	client  *ethclient.Client
	token   *Token.Token
	auth    *bind.TransactOpts
}

func NewClient() (*ClientManage, error) {
	rpcDial, err := rpc.Dial(RawUrl)
	if err != nil {
		log.Printf("[%s]Failed to init rpc client error:%v\n", Now, err)
		return nil, err
	}
	client := ethclient.NewClient(rpcDial)
	token, e := Token.NewToken(common.HexToAddress(contract), client)
	if e != nil {
		log.Printf("[%s]Failed to init token error:%v\n", Now, e)
		return nil, e
	}
	return &ClientManage{
		rpcConn: rpcDial,
		client:  client,
		token:   token,
	}, nil
}

func (c *ClientManage) CreateAccount(password string) (account string, err error) {
	ks := keystore.NewKeyStore("./wallets/", keystore.StandardScryptN, keystore.StandardScryptP)
	addressJson, _ := ks.NewAccount(password)
	a, e := ks.Export(addressJson, password, password)
	if e != nil {
		log.Printf("[%s]Failed to create new account error:%v\n", Now, e)
		return "", e
	}
	address := tools.Address{}
	if err := json.Unmarshal(a, &address); err != nil {
		log.Printf("[%s]Failed to parse json error:%v\n", Now, err)
		return "", err
	}
	// fmt.Printf("account:%s\n", account)
	return address.Address, nil
}

// 测试地址 0x5029C7e715cB5FAA4c17E6f503f6a1ea8b3870A5
// 查询钱包余额
func (c *ClientManage) SelectBalance(address string) (balance *big.Float, err error) {
	b, e := c.token.BalanceOf(nil, common.HexToAddress(address))
	if e != nil {
		log.Printf("[%s]Failed to select wallet balance error:%v\n", Now, e)
		return nil, e
	}
	BigFloat := new(big.Float)
	FloatBalance := BigFloat
	FloatBalance.SetString(b.String())
	balance = BigFloat.Quo(FloatBalance, big.NewFloat(math.Pow10(18)))
	return
}

// 初始化链
func (c *ClientManage) NewTransactorChainID() error {
	privateKey, err := crypto.HexToECDSA("83d10f228b1a7aa65164a8fc425c3af00d6577be6d5060fa26f992949682b849")
	if err != nil {
		log.Fatal(err)
		return err
	}
	auth, err2 := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(128))
	if err2 != nil {
		log.Printf("[%s]Init Transactor chainId error:%v\n", Now, err2)
		return err2
	}
	c.auth = auth
	return nil
}

// 转账到地址
// Address 收款地址
func (c *ClientManage) TransferToAddress(Address string, Number float64) (string, error) {
	toAddress := common.HexToAddress(Address)
	// val, err := c.SelectBalance(Address)
	// if err != nil {
	// 	log.Printf("[%s]Failed to select balance: %v\n", Now, err)
	// 	return err
	// }
	// fmt.Printf("[%s]before transfer :%s\n", Now, val)

	// Create an authorized transactor and spend 1 unicorn
	if c.auth == nil {
		if err := c.NewTransactorChainID(); err != nil {
			log.Printf("[%s]Failed to create authorized transactor: %v\n", Now, err)
			return "", err
		}
	}

	// 每个代币都会有相应的位数，例如eos是18位，那么我们转账的时候，需要在金额后面加18个0
	decimal, err3 := c.token.Decimals(nil)
	if err3 != nil {
		log.Printf("[%s]Failed to create decimal: %v\n", Now, err3)
		return "", err3
	}

	tenDecimal := big.NewFloat(math.Pow(10, float64(decimal.Int64())))
	convertAmount, _ := new(big.Float).Mul(tenDecimal, big.NewFloat(Number)).Int(&big.Int{})
	tx, txErr := c.token.Transfer(c.auth, toAddress, convertAmount)

	if txErr != nil {
		log.Printf("[%s]Failed to request token transfer: %v\n", Now, txErr)
		return "", txErr
	}
	ctx := context.Background()
	receipt, WaitErr := bind.WaitMined(ctx, c.client, tx)

	if WaitErr != nil {
		log.Printf("[%s]tx mining error:%v\n", Now, WaitErr)
		return "", WaitErr
	}
	if receipt.Status != 1 {
		return "", errors.New("交易失败")
	}
	// fmt.Printf("tx is :%s\n", tx.Hash().String())
	// fmt.Printf("receipt is :%s\n", receipt)

	return tx.Hash().String(), nil
}

// 读取最新的头部区块数
func (c *ClientManage) ReadNewHeaderBlock() (int64, error) {
	header, err := c.client.HeaderByNumber(context.Background(), nil)

	if err != nil {
		log.Printf("[%s]Failed to read header block error:%e", Now, err)
		return 0, err
	}

	fmt.Println(header.Number.Int64()) // 5671744
	block := header.Number.Int64()
	return block, nil
}

// 读取指定区块的信息
func (c *ClientManage) ReadBlockInfo(blockNumber int64) (*types.Block, error) {
	number := big.NewInt(blockNumber)
	block, err := c.client.BlockByNumber(context.Background(), number)

	if err != nil {
		log.Printf("[%s]Failed to read block info error:%e", Now, err)
		return nil, err
	}

	return block, nil
}
func (c *ClientManage) ReadLogs() {
	query := ethereum.FilterQuery{}
	c.client.FilterLogs(context.Background(), query)
}

// 测试 0x92927e603a0b31a2009d82182eca1eca343b80d049910eb4e1f3a7f2d6a2285c
// 通过事务hash来获取交易事务内容
func (c *ClientManage) QueryTransactionByTxHash(hash string) (res tools.TransactionResponse, err error) {
	// 上下文
	ctx := context.Background()
	// 转hash处理
	TxHash := common.HexToHash(hash)
	// 通过hash查询交易事务信息
	tx, _, e := c.client.TransactionByHash(ctx, TxHash)
	if e != nil {
		log.Printf("[%s]Failed to query transaction error:%e", Now, e)
		return res, e
	}
	// 获取最新的链ID
	chainID, e3 := c.client.NetworkID(context.Background())
	if e3 != nil {
		log.Printf("[%s]Failed to query transaction error:%e", Now, e3)
		return res, e3
	}
	// 通过链ID来获取消息
	msg, e4 := tx.AsMessage(types.NewEIP155Signer(chainID))
	if e4 != nil {
		log.Printf("[%s]Failed to query transaction error:%e", Now, e4)
		return res, e4
	}
	// 通过单个hash来获取状态等信息
	receipt, e5 := c.client.TransactionReceipt(context.Background(), tx.Hash())
	if e5 != nil {
		log.Printf("[%s]Failed to query transaction error:%e", Now, e5)
		return res, e5
	}
	// 返回结果
	res = tools.TransactionResponse{
		TxHash:     hash,
		Block:      receipt.BlockNumber,
		From:       msg.From().Hex(),
		To:         msg.To().Hex(),
		GasPrice:   msg.GasPrice(),
		Value:      msg.Value(),
		Gas:        msg.Gas(),
		Nonce:      msg.Nonce(),
		Data:       msg.Data(),
		CheckNonce: msg.CheckNonce(),
		Status:     receipt.Status,
	}
	// fmt.Printf("Status:%s\n", receipt.Status)
	return
}

func (c *ClientManage) CloseClient() {
	c.client.Close()
	c.rpcConn.Close()
}

type BatchTransfer struct {
	Address common.Address `json:"address"`
	Amount  float64        `json:"amount"`
}

// 批量转账到地址
func (c *ClientManage) BatchTransferToAddress(list []BatchTransfer) (string, error) {
	if list == nil {
		return "", errors.New("地址为空")
	}
	if c.auth == nil {
		if err := c.NewTransactorChainID(); err != nil {
			log.Printf("[%s]Failed to create authorized transactor: %v\n", Now, err)
			return "", err
		}
	}
	address := make([]common.Address, 0, len(list))
	amount := make([]*big.Int, 0, len(list))

	// 每个代币都会有相应的位数，例如eos是18位，那么我们转账的时候，需要在金额后面加18个0
	decimal, err3 := c.token.Decimals(nil)
	if err3 != nil {
		log.Printf("[%s]Failed to create decimal: %v\n", Now, err3)
		return "", err3
	}

	tenDecimal := big.NewFloat(math.Pow(10, float64(decimal.Int64())))
	for _, item := range list {
		address = append(address, item.Address)
		convertAmount, _ := new(big.Float).Mul(tenDecimal, big.NewFloat(item.Amount)).Int(&big.Int{})
		amount = append(amount, convertAmount)
	}

	tx, err := c.token.BatchTransfer(c.auth, address, amount)
	if err != nil {
		log.Printf("[%s]Failed to request token transfer: %v\n", Now, err)
		return "", err
	}
	ctx := context.Background()
	receipt, WaitErr := bind.WaitMined(ctx, c.client, tx)

	if WaitErr != nil {
		log.Printf("[%s]tx mining error:%v\n", Now, WaitErr)
		return "", WaitErr
	}
	if receipt.Status != 1 {
		return "", errors.New("交易失败")
	}

	return tx.Hash().String(), nil
}

// 检查usd的hash
func (c *ClientManage) CheckUsdHash(hash string) (res *tools.TransactionResponse, err error) {
	hashReply, err := c.client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		log.Printf("[%v]query hash failed error:%e\n", time.Now().Format("2006-05-04 15:02:01"), err)
		return nil, err
	}
	if len(hashReply.Logs) == 0 {
		return nil, errors.New("哈希地址日志为空")
	}
	if len(hashReply.Logs[0].Topics) == 0 {
		return nil, errors.New("充值地址日志主题为空")
	}

	hashFrom := strings.ToLower(common.HexToAddress(hashReply.Logs[0].Topics[1].Hex()).Hex())
	hashTo := strings.ToLower(common.HexToAddress(hashReply.Logs[0].Topics[2].Hex()).Hex())
	hashMoney := common.BytesToHash(hashReply.Logs[0].Data).Big().String()

	fmt.Printf("to:%s,money:%s\n", hashTo, hashMoney)

	if hashReply.Status != 1 {
		return nil, errors.New("交易未完成")
	}
	res = &tools.TransactionResponse{
		TxHash: hash,
		Block:  hashReply.BlockNumber,
		From:   hashFrom,
		To:     hashTo,
		Status: hashReply.Status,
	}
	return res, nil
}
