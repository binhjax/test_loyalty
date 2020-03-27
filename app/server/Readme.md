**********************************************************************************************
# **************  API List  ***************

## 0. Authentication
  + Endpoint: /auth/login
  + Method: POST
  + Input:
  {
    "username":"binhnt",
    "password":"123456"
    }
  + Output:
  {
    "rescode": 0,
    "resdecr": "New Jwt token successfully",
    "resdata": {
      "expireAt": "2020-03-25 15:49:24.889251 +0700 +07 m=+1082.684709660",
      "fasthttp_jwt": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlltbHVhRzUwIiwicGFzc3dvcmQiOiJNVEl6TkRVMiIsImV4cCI6MTU4NTEyNjE2NH0.dzBB27zalr7ybyunZaCdTArlh2TwdiAifVrtidyjicTjorBtrkhC5Cd-AdBRv1dOufSlpWZYpsq-WvzAyq0xxA"
    }
  }

## 1. Admin  
  ### 1.1 Network
    + Connect peers
    + Endpoint: /api/admin/network/addpeers
    + Method: GET
    {"rescode":0,"resdecr":"Add peers successfully","resdata":null}

  ### 1.2. Account
    + From store
    + From file
    + Save to file
    + List token
  ### 1.3. Wallet
    + Load
    + List
    + New
    + Save
    + Fillgas
    + Balance
    + Send Eth between account

  ### 1.3. Contract
    + Deploy Contract
    + Initlization Contract
    ##### 1. Register EthAccount
    http://localhost:8080/api/v2/wallet/register
## 2. API v1

### 2.1 Wallet
    ####2.1.1. Create Wallet
    http://localhost:8080/api/v2/wallet/create/vi03/1

    ##### 2.1.3. Set Wallet state
    http://localhost:8080/api/v2/wallet/set_state/vi03/1

    ##### 2.1.4. View wallet balance
    http://localhost:8080/api/v2/wallet/balance/vi03
### 2.2 Transactions  
  ##### 2.2.1. Credit
  http://localhost:8080/api/v2/wallet/deposit/vi03/vi03/1000

  ##### 2.2.2. Debit
  http://localhost:8080/api/v2/wallet/withdraw/vi03/vi03/1000

  ##### 2.2.3. Transfer Money
  http://localhost:8080/api/v2/wallet/transfer/tx01/VI01/vi03/100/Test/1

### 2.3 Report
  http://localhost:8080/api/v2/wallet/summary
