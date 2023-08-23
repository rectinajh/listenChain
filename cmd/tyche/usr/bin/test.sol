// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.6.0;

contract SimpleContract {
    uint256 public value;

    function setValue(uint256 newValue) public {
        value = newValue;
    }

    function getValue() public view returns (uint256) {
        return value;
    }
}