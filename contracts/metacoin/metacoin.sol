pragma solidity ^0.6.4;

contract metaCoin {
      mapping (address => uint) public balances;

    	constructor(uint initialBalance) public {
    		balances[tx.origin] = initialBalance;
    	}
    	function sendToken(address receiver, uint amount) public returns(bool successful){
    		if (balances[tx.origin] < amount) return false;
     		balances[tx.origin] -= amount;
     		balances[receiver] += amount;
     		return true;
     	}
      function balanceOf(address _account) public view returns (uint) {
            return balances[_account];
        }
}
