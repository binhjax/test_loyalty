/*
    @author: DuHD
    @version: 1.0
    @date: 09/04/2019

    Contract duoc deploy boi Business Contract, no human access
*/

pragma solidity ^0.6.4;

import "./Owned.sol";


contract WalletHistory is Owned {
    int version;
    uint createTime;

    constructor(int _version) public {
        owner = msg.sender;
        createTime = now;
        version = _version;
    }

    //------------------------------------------------------------------------------------------------------------
    // Luu giu lich su giao dich createStash
    struct CreateStash {
        bytes32 txRef;
        bytes32 stashName;
        address stashAddr;
        uint timestamp_onchain; // Thoi diem thuc hien giao dich tai onchain
        uint timestamp_offchain; // Thoi diem thuc hien giao dich tai offchain
    }
    bytes32[] private createStashIdx;
    mapping(bytes32 => CreateStash) private createStashHistorys;

    function setCreateStashHistory(bytes32 _txRef, bytes32 _stashName, address _stashAddr, uint _timestamp_onchain, uint timestamp_offchain) onlyOwner external {
        require(!(createStashHistorys[_txRef].txRef > 0x0), 'DA TON TAI GIAO DICH TAO TAI KHOAN VOI TRACE_ID NAY');
        createStashIdx.push(_txRef);
        createStashHistorys[_txRef].txRef = _txRef;
        createStashHistorys[_txRef].stashName = _stashName;
        createStashHistorys[_txRef].stashAddr = _stashAddr;
        createStashHistorys[_txRef].timestamp_onchain = _timestamp_onchain;
        createStashHistorys[_txRef].timestamp_offchain = timestamp_offchain;
    }
    function isCreateStashHistory(bytes32 _txRef) view external returns (bool){
        return (createStashHistorys[_txRef].txRef > 0x0);
    }

/**Đếm số Stash được tạo trong khoảng thời gian *******************************************************************/
    function countCreateStash(uint _from_timestamp, uint _to_timestamp) view public returns(int)
    {
        require(_from_timestamp < _to_timestamp, 'YEU CAU FROM_TIME < TO_TIME');
        int count = 0;
        uint cur_time;
        uint i = createStashIdx.length;
        do {
            cur_time = createStashHistorys[createStashIdx[i]].timestamp_offchain;
            if (cur_time <= _to_timestamp) count++;
            i--;
        } while (cur_time < _from_timestamp && i >= 1);
        return (count);
    }

    //------------------------------------------------------------------------------------------------------------
    // Luu giu lich su giao dich debit
    struct Debit {
        bytes32 txRef;
        bytes32 stashName;
        int amount;
        uint timestamp_onchain; // Thoi diem thuc hien giao dich tai onchain
        uint timestamp_offchain; // Thoi diem thuc hien giao dich tai offchain
    }
    bytes32[] private debitIdx;
    mapping(bytes32 => Debit) private debits;

    function getDebitHistoryIdxLength() view external returns (uint){
        return debitIdx.length;
    }

    function setDebitHistory(bytes32 _txRef, bytes32 _stashName, int _amount, uint _timestamp_onchain, uint timestamp_offchain) onlyOwner external {
        require(!(debits[_txRef].txRef > 0x0), 'DA TON TAI GIAO DICH DEBIT VOI TRACE_ID NAY');
        debitIdx.push(_txRef);
        debits[_txRef].txRef = _txRef;
        debits[_txRef].stashName = _stashName;
        debits[_txRef].amount = _amount;
        debits[_txRef].timestamp_onchain = _timestamp_onchain;
        debits[_txRef].timestamp_offchain = timestamp_offchain;
    }

    function getDebitHistoryByTxRef(bytes32 _txRef) view external returns (bytes32, bytes32, int, uint, uint){
        require(debits[_txRef].txRef > 0x0, 'KHONG TON TAI GIAO DICH DEBIT VOI TRACE_ID NAY');
        return (debits[_txRef].txRef, debits[_txRef].stashName, debits[_txRef].amount, debits[_txRef].timestamp_onchain, debits[_txRef].timestamp_offchain);
    }

    function isDebitHistory(bytes32 _txRef) view external returns (bool){
        return (debits[_txRef].txRef > 0x0);
    }

/**Đếm và Cộng gd Debit được tạo trong khoảng thời gian *******************************************************************/
    function countAndSumDebit(uint _from_timestamp, uint _to_timestamp) view public returns(int, int){
        require(_from_timestamp < _to_timestamp, 'YEU CAU FROM_TIME < TO_TIME');
        int count = 0;
        int sum = 0;
        uint cur_time;
        uint i = debitIdx.length;
        do {
            cur_time = debits[debitIdx[i]].timestamp_offchain;
            if (cur_time <= _to_timestamp) {
                count++;
                sum += debits[debitIdx[i]].amount;
            }
            i--;
        } while (cur_time < _from_timestamp && i >= 1);
        return (count, sum);
    }

    //------------------------------------------------------------------------------------------------------------
    // Luu giu lich su giao dich credit
    struct Credit {
        bytes32 txRef;
        bytes32 stashName;
        int amount;
        uint timestamp_onchain; // Thoi diem thuc hien giao dich tai onchain
        uint timestamp_offchain; // Thoi diem thuc hien giao dich tai offchain
    }
    bytes32[] private creditIdx;
    mapping(bytes32 => Credit) private credits;

    function getCreditHistoryIdxLength() view external returns (uint){
        return creditIdx.length;
    }

    function setCreditHistory(bytes32 _txRef, bytes32 _stashName, int _amount, uint _timestamp_onchain, uint timestamp_offchain) onlyOwner external {
        require(!(credits[_txRef].txRef > 0x0), 'DA TON TAI GIAO DICH CREDIT VOI TRACE_ID NAY');
        creditIdx.push(_txRef);
        credits[_txRef].txRef = _txRef;
        credits[_txRef].stashName = _stashName;
        credits[_txRef].amount = _amount;
        credits[_txRef].timestamp_onchain = _timestamp_onchain;
        credits[_txRef].timestamp_offchain = timestamp_offchain;
    }

    function getCreditHistoryByTxRef(bytes32 _txRef) view external returns (bytes32, bytes32, int, uint, uint){
        require(credits[_txRef].txRef > 0x0, 'KHONG TON TAI GIAO DICH CREDIT VOI TRACE_ID NAY');
        return (credits[_txRef].txRef, credits[_txRef].stashName, credits[_txRef].amount, credits[_txRef].timestamp_onchain, credits[_txRef].timestamp_offchain);
    }

    function isCreditHistory(bytes32 _txRef) view external returns (bool){
        return (credits[_txRef].txRef > 0x0);
    }

    function countAndSumCredit(uint _from_timestamp, uint _to_timestamp) view public returns(int, int){
        require(_from_timestamp < _to_timestamp, 'YEU CAU FROM_TIME < TO_TIME');
        int count = 0;
        int sum = 0;
        uint cur_time;
        uint i = creditIdx.length;
        do {
            cur_time = credits[creditIdx[i]].timestamp_offchain;
            if (cur_time <= _to_timestamp) {
                count++;
                sum += credits[creditIdx[i]].amount;
            }
            i--;
        } while (cur_time < _from_timestamp && i >= 1);
        return (count, sum);
    }

    //------------------------------------------------------------------------------------------------------------
    // Luu giu lich su giao dich transfer
    struct Transfer {
        bytes32 txRef;
        bytes32 sender;
        bytes32 receiver;
        int amount; // > 0
        string note;
        uint timestamp_onchain; // Thoi diem thuc hien giao dich tai onchain
        uint timestamp_offchain; // Thoi diem thuc hien giao dich tai offchain
    }

    bytes32[] private transferIdx;                  // @private (list of all-trans)
    mapping(bytes32 => Transfer) private transfers; // @private

    function getTransferHistoryIdxLength() view external returns (uint){
        return transferIdx.length;
    }

    function setTransferHistory(bytes32 _txRef, bytes32 _sender, bytes32 _receiver, int _amount, string memory _note, uint _timestamp_onchain, uint _timestamp_offchain) onlyOwner public {
        require(!(transfers[_txRef].txRef > 0x0), 'DA TON TAI GIAO DICH TRANSFER VOI TRACE_ID NAY');
        transferIdx.push(_txRef);
        transfers[_txRef].txRef = _txRef;
        transfers[_txRef].sender = _sender;
        transfers[_txRef].receiver = _receiver;
        transfers[_txRef].amount = _amount;
        transfers[_txRef].note = _note;
        transfers[_txRef].timestamp_onchain = _timestamp_onchain;
        transfers[_txRef].timestamp_offchain = _timestamp_offchain;
    }

    function getTransferHistoryByTxRef(bytes32 _txRef) view external returns (bytes32, bytes32, int, string memory, uint, uint){
        require(transfers[_txRef].txRef > 0x0, 'KHONG TON TAI GIAO DICH TRANSFER VOI TRACE_ID NAY');
        string memory note = transfers[_txRef].note;
        int amount = transfers[_txRef].amount;
        return (transfers[_txRef].sender, transfers[_txRef].receiver, amount, note, transfers[_txRef].timestamp_onchain, transfers[_txRef].timestamp_offchain);
    }

    function isTransferHistory(bytes32 _txRef) view external returns (bool){
        return (transfers[_txRef].txRef > 0x0);
    }

    function countAndSumTransfer(uint _from_timestamp, uint _to_timestamp) view public returns(int, int){
        require(_from_timestamp < _to_timestamp, 'YEU CAU FROM_TIME < TO_TIME');
        int count = 0;
        int sum = 0;
        uint cur_time;
        uint i = transferIdx.length;
        do {
            cur_time = transfers[transferIdx[i]].timestamp_offchain;
            if (cur_time <= _to_timestamp) {
                count++;
                sum += transfers[transferIdx[i]].amount;
            }
            i--;
        } while (cur_time < _from_timestamp && i >= 1);
        return (count, sum);
    }

    //------------------------------------------------------------------------------------------------------------
    // Luu giu lich su giao dich reverts
    struct Revert {
        bytes32 txRef;
        bytes32 txRef_org;
        string note;
        int amount;
        uint timestamp_onchain; // Thoi diem thuc hien giao dich tai onchain
        uint timestamp_offchain; // Thoi diem thuc hien giao dich tai offchain
    }

    bytes32[] private revertIdx;                  // @private (list of all-trans)
    mapping(bytes32 => Revert) private reverts; // @private

    function getRevertHistoryIdxLength() view external returns (uint){
        return revertIdx.length;
    }

    function setRevertHistory(bytes32 _txRef, bytes32 _txRef_org, int _amount, string memory _note,  uint _timestamp_onchain, uint _timestamp_offchain) onlyOwner public {
        require(!(reverts[_txRef].txRef > 0x0), 'DA TON TAI GIAO DICH REVERT VOI TRACE_ID NAY');
        revertIdx.push(_txRef);
        reverts[_txRef].txRef = _txRef;
        reverts[_txRef].txRef_org = _txRef_org;
        reverts[_txRef].note = _note;
        reverts[_txRef].amount = _amount;
        reverts[_txRef].timestamp_onchain = _timestamp_onchain;
        reverts[_txRef].timestamp_offchain = _timestamp_offchain;
    }

    function getRevertHistoryByTxRef(bytes32 _txRef) view external returns (bytes32, string memory, uint, uint){
        require(reverts[_txRef].txRef > 0x0, 'KHONG TON TAI GIAO DICH REVERT VOI TRACE_ID NAY');
        bytes32 txRef_org = reverts[_txRef].txRef_org;
        string memory note = reverts[_txRef].note;
        return (txRef_org, note, reverts[_txRef].timestamp_onchain, reverts[_txRef].timestamp_offchain);
    }

    function isRevertHistory(bytes32 _txRef) view external returns (bool){
        return (reverts[_txRef].txRef > 0x0);
    }

    function countAndSumRevert(uint _from_timestamp, uint _to_timestamp) view public returns(int, int){
        require(_from_timestamp < _to_timestamp, 'YEU CAU FROM_TIME < TO_TIME');
        int count = 0;
        int sum = 0;
        uint cur_time;
        uint i = revertIdx.length;
        do {
            cur_time = reverts[revertIdx[i]].timestamp_offchain;
            if (cur_time <= _to_timestamp) {
                count++;
                sum += reverts[revertIdx[i]].amount;
            }
            i--;
        } while (cur_time < _from_timestamp && i >= 1);
        return (count, sum);
    }
}
