pragma solidity ^0.6.4;

contract Inbox {

    string public message;

    function doInbox(string memory initialMessage) public {
        message = initialMessage;
    }

    function setMessage(string  memory newMessage) public {
        message = newMessage;
    }
}
