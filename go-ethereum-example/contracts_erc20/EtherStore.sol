// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract EtherStore {

    // 初始化互斥锁
    bool reEntrancyMutex = false;
    uint256 public withdrawalLimit = 1 ether;
    mapping(address => uint256) public lastWithdrawTime;
    mapping(address => uint256) public balances;

    function depositFunds() public payable {
        balances[msg.sender] += msg.value;
    }

    function withdrawFunds(uint256 _weiToWithdraw) public {
        require(!reEntrancyMutex);
        require(balances[msg.sender] >= _weiToWithdraw);
        // 限制取款金额
        require(_weiToWithdraw <= withdrawalLimit);
        // 限制取款时间
        require(block.timestamp >= lastWithdrawTime[msg.sender] + 1 weeks);
        balances[msg.sender] -= _weiToWithdraw;
        lastWithdrawTime[msg.sender] = block.timestamp;
        // 在外部调用前设置互斥锁
        reEntrancyMutex = true;
        payable(msg.sender).transfer(_weiToWithdraw);
        // 在外部调用后释放互斥锁
        reEntrancyMutex = false;
    }
}
