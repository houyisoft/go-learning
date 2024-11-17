// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    // 定义一个状态变量，用于存储整数
    uint256 private storedValue;

    // 设置存储的值
    function set(uint256 _value) public {
        storedValue = _value;
    }

    // 获取当前存储的值
    function get() public view returns (uint256) {
        return storedValue;
    }
}