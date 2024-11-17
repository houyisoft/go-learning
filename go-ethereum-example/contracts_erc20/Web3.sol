// 导入 Web3 库
const Web3 = require('web3');

// 连接到以太坊网络的节点（本地节点或者 Infura 等提供的节点）
const web3 = new Web3('http://localhost:8545');

// 合约地址和 ABI（在部署合约时生成的 ABI）
const contractAddress = '0x1234567890abcdef1234567890abcdef12345678';
const abi = [
{
"constant": false,
"inputs": [
{
"name": "x",
"type": "uint256"
}
],
"name": "set",
"outputs": [],
"payable": false,
"stateMutability": "nonpayable",
"type": "function"
},
{
"constant": true,
"inputs": [],
"name": "get",
"outputs": [
{
"name": "",
"type": "uint256"
}
],
"payable": false,
"stateMutability": "view",
"type": "function"
}
];

// 使用合约地址和 ABI 创建合约实例
const contract = new web3.eth.Contract(abi, contractAddress);

// 调用合约的 get 函数（读取当前存储的值）
contract.methods.get().call()
.then(value => {console.log('Current value:', value);
})
.catch(error => {console.error('Error:', error);
});

// 调用合约的 set 函数（设置新的值）const newValue = 42;
contract.methods.set(newValue).send({ from: '0xabcdef1234567890abcdef1234567890abcdef12' })
.then(receipt => {console.log('Transaction receipt:', receipt);
})
.catch(error => {console.error('Error:', error);
});