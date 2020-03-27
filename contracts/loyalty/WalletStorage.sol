/*
    @author: DuHD
    @version: 1.0
    @date: 09/04/2019

    Contract duoc deploy boi Business Contract, no human access 
*/

pragma solidity ^0.6.4;

import "./Owned.sol";


contract WalletStorage is Owned {
    int version;
    uint createTime;
    
    constructor(int _version) public {
        owner = msg.sender;
        version = _version;
        createTime = now;
    }

    //------------------------------------------------------------------------------------------------------------    
    //Luu giu danh sach acc ETH 
    mapping(address => bool) private memberApi;
    address[] private memberApiIdx;
    
    function getMemberApi(address _member) view external returns (bool){
        return memberApi[_member];
    }
    
    function setMemberApi(address _member) onlyOwner external {
        memberApi[_member] = true;
        memberApiIdx.push(_member);
    }

    function unSetMemberApi(address _member) onlyOwner external {
        memberApi[_member] = false;
    }

    function getMemberApiIdx() view external returns (address[] memory){
        return memberApiIdx;
    }
    
    function getMemberApiLength() view external returns (int16) {
        int16 memberLength = 0;
        for (uint i = 0; i < memberApiIdx.length; i++) {
            if (memberApi[memberApiIdx[i]] == true) {
                memberLength++;
            }
        }
        return memberLength;
    }    

    //------------------------------------------------------------------------------------------------------------
    // Bang luu giu address cua account, mapping giữa mã số tải khoản và address contract  - account_id => contract address
    mapping(bytes32 => address) private stashRegistry;
    bytes32[] private stashNames;
    
    function getStashRegistry(bytes32 _stashName) view external returns (address){
        require(stashRegistry[_stashName] > address(0x0), 'TAI KHOAN KHONG TON TAI');
        return stashRegistry[_stashName];
    }
    
    function getStashRegistryByIdx(uint _idx) view external returns (bytes32, address){
        return (stashNames[_idx], stashRegistry[stashNames[_idx]]);
    }
    
    function getStashByIdx(uint _idx) view external returns (address){
        return stashRegistry[stashNames[_idx]];
    }    
    
    function setStashRegistry(bytes32 _stashName, address _stash) onlyOwner external {
        require(!(stashRegistry[_stashName] > address(0x0)), 'TAI KHOAN DA TON TAI');
        stashRegistry[_stashName] = _stash;
        stashNames.push(_stashName);
    }
    
    function getStashNamesLength() view external returns (uint){
        return stashNames.length;
    }
    
    function getStashNames(uint _idx) view external returns (bytes32){
        return stashNames[_idx];
    }

    function getAllStashRegistry() view external returns (bytes32[] memory, address[] memory) {
        uint length = stashNames.length;
        address[] memory stashAddr = new address[](length);
        for (uint i = 0; i < length; i++) {
           stashAddr[i] = stashRegistry[stashNames[i]];
        }
        return (stashNames,stashAddr);
    } 
    function isStashRegistry(bytes32 _stashName) view external returns (bool){
        return (stashRegistry[_stashName] > address(0x0));
    }
}