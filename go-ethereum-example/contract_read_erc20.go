package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"    // 提供以太坊地址相关功能
	"github.com/ethereum/go-ethereum/ethclient" // 提供以太坊客户端连接功能
	"log"
	"math"
	"math/big"

	token "go-learning/go-ethereum-example/contracts_erc20" // 导入生成的合约绑定代码
)

// 运行步骤:
//  1. 确保已安装Go环境
//  2. 安装依赖:
//     go get github.com/ethereum/go-ethereum
//  3. 编译合约生成Go绑定代码:
//     abigen --abi=erc20.abi --pkg=token --out=contracts_erc20/erc20.go
//  4. 运行程序:
//     go run contract_read_erc20.go
//
// 本程序功能:
// - 连接以太坊主网
// - 读取Golem(GNT)代币合约信息
// - 查询指定地址的代币余额
func main() {
	// 连接以太坊节点
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Println("连接以太坊节点失败")
		log.Fatal(err)
	}

	// 设置GNT代币合约地址
	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	// 创建合约实例
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Println("创建合约实例失败")
		log.Fatal(err)
	}

	// 设置要查询余额的钱包地址
	address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
	// 查询代币余额
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Println("查询余额失败")
		log.Fatal(err)
	}

	// 查询代币名称
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// 查询代币符号
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// 查询代币精度
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// 打印代币基本信息
	fmt.Printf("name: %s\n", name)         // 输出代币名称
	fmt.Printf("symbol: %s\n", symbol)     // 输出代币符号
	fmt.Printf("decimals: %v\n", decimals) // 输出代币精度

	fmt.Printf("wei: %s\n", bal) // 输出原始余额(wei)

	// 将余额转换为可读格式
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	// 根据代币精度进行转换
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value) // 输出转换后的余额
}
