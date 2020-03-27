/*
    @author: DuHD
    @version: 1.0
    @date: 25/03/2020
    @project LOYALTY
*/

pragma solidity ^0.6.4;

import "./Owned.sol";
import "./WalletStash.sol";
import "./WalletStorage.sol";
import "./WalletHistory.sol";


contract WalletBusiness is Owned {
    event ErrorOccurred(string code, string reason, string uuid);
    
    //Deploy WalletStorage contract va luu address
    int walletStorageCtx_ver = 0;
    address walletStorageCtx_addr = address(new WalletStorage(walletStorageCtx_ver));
    WalletStorage walletStorageCtx = WalletStorage(walletStorageCtx_addr);
    
    //Deploy WalletHistory contract va luu address
    int walletHistoryCtx_ver = 0;
    address walletHistoryCtx_addr = address(new WalletHistory(walletHistoryCtx_ver));
    WalletHistory walletHistoryCtx = WalletHistory(walletHistoryCtx_addr);
    
    constructor() public {
        owner = msg.sender;
    }
    
    // TXS - Hàm để tạo walletHistoryCtx Contract mới khi cần cutoff contract cũ
    function newWalletHistoryCtx(int _version) onlyMember public
    {
        //Deploy WalletHistory contract new va luu address
        walletHistoryCtx_addr = address(new WalletHistory(_version));
        walletHistoryCtx = WalletHistory(walletHistoryCtx_addr);
    }
    // CALL - Hàm lấy địa chỉ contract walletHistoryCtx hiện tại
    function getWalletHistoryCtx_addr() view public returns (address) {
        return walletHistoryCtx_addr;
    }

    // TXS - Hàm tạo tài khoản
    function createStash(bytes32 _txRef, bytes32 _nameStash, int8 _typeStash, int8 _stateStash, int8 _levelStash, uint _createTime) onlyMember public
    {
        if (walletHistoryCtx.isCreateStashHistory(_txRef)) {
            emit ErrorOccurred('ERR471', 'DA TON TAI GIAO DICH TAO ACC VOI TRACE_ID NAY', bytes32ToString(_txRef));
        } else if (walletStorageCtx.isStashRegistry(_nameStash)) {
            emit ErrorOccurred('ERR473','TAI KHOAN DA TON TAI', bytes32ToString(_txRef));
        } else {
            address stash = address(new WalletStash(_nameStash, _typeStash, _stateStash, _levelStash, _createTime));
            walletStorageCtx.setStashRegistry(_nameStash, stash);
            walletHistoryCtx.setCreateStashHistory(_txRef,_nameStash, stash, now, _createTime);
        }
    }

   // CALL - Lay thong tin ACC
    function getStashInfo(bytes32 _nameStash) view public returns (int, int8, int8, int8, uint)
    {
        require(walletStorageCtx.isStashRegistry(_nameStash),'ERR404: TAI KHOAN KHONG TON TAI');
        int bal = 0;
        int8 type_ = 0;
        int8 state = 0;
        int8 level = 0;
        uint createTime = 0;
        WalletStash stash = WalletStash(walletStorageCtx.getStashRegistry(_nameStash));
        bal = stash.getBalance();
        type_ = stash.getType();
        state = stash.getState();
        level = stash.getLevel();
        createTime = stash.getCreateTime();
        return (bal, type_, state, level, createTime);
    }

    // TXS - Hàm thay đổi thông tin tài khoản, nếu tham số vào <= 0 nghĩa là không thay đổi
    function updateStash(bytes32 _txRef, bytes32 _nameStash, int8 _typeStash, int8 _stateStash, int8 _levelStash) onlyMember public
    {
        if (!walletStorageCtx.isStashRegistry(_nameStash)) {
            emit ErrorOccurred('ERR404','TAI KHOAN KHONG TON TAI', bytes32ToString(_txRef));
        } else {       
            WalletStash stashCtx = WalletStash(walletStorageCtx.getStashRegistry(_nameStash));
            if (_typeStash > 0) stashCtx.setType(_typeStash); //nếu tham số vào <= 0 nghĩa là không thay đổi
            if (_stateStash > 0) stashCtx.setState(_stateStash); //nếu tham số vào <= 0 nghĩa là không thay đổi
            if (_levelStash > 0) stashCtx.setLevel(_levelStash); //nếu tham số vào <= 0 nghĩa là không thay đổi
        }
    }
    
    // TXS - Nạp điểm vào ACC
    function credit(bytes32 _txRef, bytes32 _nameStash, int _amount, uint _timestamp_offchain) onlyMember public
    {   
        if (walletHistoryCtx.isCreditHistory(_txRef)) {
            emit ErrorOccurred('ERR471', 'GIAO DICH CREDIT CO TRACE_ID NAY DA TON TAI', bytes32ToString(_txRef));
        } else if (!walletStorageCtx.isStashRegistry(_nameStash)) {
            emit ErrorOccurred('ERR404','TAI KHOAN KHONG TON TAI', bytes32ToString(_txRef));
        } else {
            WalletStash stash = WalletStash(walletStorageCtx.getStashRegistry(_nameStash));
            if (stash.getState() != 1) {
                 emit ErrorOccurred('ERR472','TAI KHOAN KHONG O TRANG THAI ACTIVE', bytes32ToString(_txRef));
            } else if (_amount <= 0) {
                 emit ErrorOccurred('ERR461','AMOUNT PHAI > 0', bytes32ToString(_txRef));
            } else {
                stash.credit(_amount);
                uint timestamp = now;
                walletHistoryCtx.setCreditHistory(_txRef, _nameStash, _amount, timestamp, _timestamp_offchain);
            }
        }
    }

    // TXS - Rút điểm khỏi ACC
    function debit(bytes32 _txRef, bytes32 _nameStash, int _amount, uint _timestamp_offchain) onlyMember public
    {   
        if (walletHistoryCtx.isDebitHistory(_txRef)) {
            emit ErrorOccurred('ERR471', 'GIAO DICH DEBIT CO TRACE_ID NAY DA TON TAI', bytes32ToString(_txRef));
        } else if (!walletStorageCtx.isStashRegistry(_nameStash)) {
            emit ErrorOccurred('ERR404','TAI KHOAN KHONG TON TAI', bytes32ToString(_txRef));
        } else {
            WalletStash stash = WalletStash(walletStorageCtx.getStashRegistry(_nameStash));
            if (stash.getState() != 1) {
                 emit ErrorOccurred('ERR472','TAI KHOAN KHONG O TRANG THAI ACTIVE', bytes32ToString(_txRef));
            } else if (stash.getBalance() < _amount) {
                 emit ErrorOccurred('ERR473','TAI KHOAN KHONG DU SO DU DE THUC HIEN', bytes32ToString(_txRef));
            } else if (_amount <= 0) {
                 emit ErrorOccurred('ERR461','AMOUNT PHAI > 0', bytes32ToString(_txRef));
            } else {
                stash.safe_debit(_amount);
                uint timestamp = now;
                walletHistoryCtx.setDebitHistory(_txRef, _nameStash, _amount, timestamp, _timestamp_offchain);
            }
        }    
    }
    
    // TXS - Chuyển điểm giữa 2 ACC
    function transfer(bytes32 _txRef, bytes32 _sender, bytes32 _receiver, int _amount, string memory _note, uint _timestamp_offchain) onlyMember public
    {   
        if (walletHistoryCtx.isTransferHistory(_txRef)) {
            emit ErrorOccurred('ERR471', 'GIAO DICH TRANSFER CO TRACE_ID NAY DA TON TAI', bytes32ToString(_txRef));
        } else if (!walletStorageCtx.isStashRegistry(_sender)) {
            emit ErrorOccurred('ERR404','TAI KHOAN SENDER KHONG TON TAI', bytes32ToString(_txRef));
        } else if (!walletStorageCtx.isStashRegistry(_receiver)) {
            emit ErrorOccurred('ERR404','TAI KHOAN RECEIVER KHONG TON TAI', bytes32ToString(_txRef));
        } else {
            WalletStash senderStash = WalletStash(walletStorageCtx.getStashRegistry(_sender));
            WalletStash receiverStash = WalletStash(walletStorageCtx.getStashRegistry(_receiver));
            if (senderStash.getState() != 1) {
                 emit ErrorOccurred('ERR472','TAI KHOAN SENDER KHONG O TRANG THAI ACTIVE', bytes32ToString(_txRef));
            } else if (receiverStash.getState() != 1) {
                 emit ErrorOccurred('ERR472','TAI KHOAN RECEIVER KHONG O TRANG THAI ACTIVE', bytes32ToString(_txRef));
            } else if (senderStash.getBalance() < _amount) {
                 emit ErrorOccurred('ERR473','TAI KHOAN SENDER KHONG DU SO DU DE THUC HIEN', bytes32ToString(_txRef));
            } else if (_amount <= 0) {
                 emit ErrorOccurred('ERR461','AMOUNT PHAI > 0', bytes32ToString(_txRef));
            } else {
                senderStash.safe_debit(_amount);
                receiverStash.credit(_amount);
                uint timestamp = now;
                walletHistoryCtx.setTransferHistory(_txRef,_sender,_receiver,_amount,_note, timestamp, _timestamp_offchain);
            }
        }
    }
      
    // CALL - Lấy số lượng lịch sử giao dịch DEBIT
    function getDebitHistoryLength() view public returns (uint) {
        return walletHistoryCtx.getDebitHistoryIdxLength();
    }
    
    // CALL - Lấy lịch sử giao dịch DEBIT theo TRACE_NO
    function getDebitHistoryByTxRef(bytes32 _txRef) view public returns (bytes32, bytes32, int, uint, uint){
        return walletHistoryCtx.getDebitHistoryByTxRef(_txRef);
    }

    // CALL - Lấy số lượng lịch sử giao dịch CREDIT
    function getCreditHistoryLength() view public returns (uint) {
        return walletHistoryCtx.getCreditHistoryIdxLength();
    }
    
    // CALL - Lấy lịch sử giao dịch CREDIT theo TRACE_NO
    function getCreditHistoryByTxRef(bytes32 _txRef) view public returns (bytes32, bytes32, int, uint, uint){
        return walletHistoryCtx.getCreditHistoryByTxRef(_txRef);
    }
    
    // CALL - Lấy số lượng lịch sử giao dịch TRANSFER
    function getTransferHistoryLength() view public returns (uint) {
        return walletHistoryCtx.getTransferHistoryIdxLength();
    }

    // CALL - Lấy lịch sử giao dịch TRANSFER theo TRACE_NO  
    function getTransferHistoryByTxRef(bytes32 _txRef) view public returns (bytes32, bytes32, int, string memory, uint, uint){
        return walletHistoryCtx.getTransferHistoryByTxRef(_txRef);
    }
    


// TXS - Đăng ký Acc ETH
    function registerAccETH(address[] memory _listAcc) onlyMember public {
        for (uint i = 0; i < _listAcc.length; i++) {
            if (walletStorageCtx.getMemberApi(_listAcc[i]) != true) {
                walletStorageCtx.setMemberApi(_listAcc[i]);
            }
        }
    }

// TXS - Hủy đăng ký acc ETH
    function unRegisterAccETH(address[] memory _listAcc) onlyMember public {
        for (uint i = 0; i < _listAcc.length; i++) {
            if (walletStorageCtx.getMemberApi(_listAcc[i]) == true) {
                walletStorageCtx.unSetMemberApi(_listAcc[i]);
            }
        }
    }

// CALL - Lấy danh sách acc ETH đã đăng ký
    function getRegisterApiIdx() view public returns (address[] memory){
        return walletStorageCtx.getMemberApiIdx();
    }
    
// CALL - Lấy số lượng acc ETH đã đăng ký
    function getRegisterAccEthLength() view public returns (int16) {
        return walletStorageCtx.getMemberApiLength();
    }

// CALL - Lấy số lượng ACC điểm trên hệ thống
    function getStashNamesLength() view public returns (uint) {
        return walletStorageCtx.getStashNamesLength();

    }
// CALL - Lấy danh sách ACC điểm trên hệ thống
    function getAllStashRegistry() view public returns (bytes32[] memory, address[] memory) {
        return walletStorageCtx.getAllStashRegistry();
    }
    
// TXS - Thay đổi Owned của các Contact Stash, khi deploy Contract Bussiness mới
    function changeOwnerAllStash(address _newOwner) onlyOwner public {
        uint length = walletStorageCtx.getStashNamesLength();
        bytes32[] memory stashName = new bytes32[](length);
        address[] memory stashAddr = new address[](length);
        (stashName, stashAddr) = walletStorageCtx.getAllStashRegistry();
        for (uint i = 0; i < length; i++) {
            WalletStash stash = WalletStash(stashAddr[i]);
            stash.changeOwner(_newOwner);
        }
    }    

// TXS - Thay đổi Owned của các Contact Storage, khi deploy Contract Bussiness mới 
    function changeOwnerStorage(address _newOwner) onlyOwner public {
        walletStorageCtx.changeOwner(_newOwner);
    }

// TXS - Thay đổi Owned của các Contact History, khi deploy Contract Bussiness mới 
    function changeOwnerHistory(address _newOwner) onlyOwner public {
        walletHistoryCtx.changeOwner(_newOwner);
    }

//  CALL - Lấy thông tin của tất cả các tài khoản điểm
    function getAllStashInfo() view public returns (bytes32[] memory, address[] memory, int[] memory, int8[] memory, int8[] memory, int8[] memory) {
        uint length = walletStorageCtx.getStashNamesLength();
        bytes32[] memory stashName = new bytes32[](length);
        address[] memory stashAddr = new address[](length);
        int[] memory balance = new int[](length);
        int8[] memory state = new int8[](length);
        int8[] memory typeStash = new int8[](length);
        int8[] memory level = new int8[](length);
        
        (stashName, stashAddr) = walletStorageCtx.getAllStashRegistry();
        for (uint i = 0; i < length; i++) {
            WalletStash stash = WalletStash(stashAddr[i]);
            balance[i] = stash.getBalance();
            state[i] = stash.getState();
            typeStash[i] = stash.getType();
            level[i] =  stash.getLevel();
        }        
        return (stashName, stashAddr, balance, state, typeStash, level);
    }  
    
    modifier onlyMember() {
        require(walletStorageCtx.getMemberApi(msg.sender), 'CHI CAC ACC ETH DA DANG KY MOI GOI DUOC HAM');
        _;
    }

    function bytes32ToString(bytes32 _bytes32) public pure returns (string memory) {
            uint8 i = 0;
            while(i < 32 && _bytes32[i] != 0) {
                i++;
            }
            bytes memory bytesArray = new bytes(i);
            for (i = 0; i < 32 && _bytes32[i] != 0; i++) {
                bytesArray[i] = _bytes32[i];
            }
            return string(bytesArray);
        }
    
    
}