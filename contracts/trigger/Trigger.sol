pragma solidity ^0.6.4;

contract Trigger {
  address owner;

  constructor() public {
      owner = msg.sender;
  }
  event TriggerEvt(address _sender, uint _trigger);


  function Test(uint _trigger) public {
    require(_trigger > 0 , 'CHI OWNER CONTRACT MOI DUOC GOI HAM');
    emit  TriggerEvt(msg.sender, _trigger);
  }

  function getOwner() external returns  (address)  {
    return owner;
  }
}
