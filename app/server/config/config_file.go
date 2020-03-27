package config

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "strings"
  "path/filepath"
  "os"
  "time"
  "github.com/ethereum/go-ethereum/accounts/keystore"
  "encoding/hex"
  "github.com/ethereum/go-ethereum/crypto"
  "encoding/json"
)

func LoadConfigFromFile(file string) *Config {
     var fcfg = &Config{}

     yamlFile, err := ioutil.ReadFile(file)
     if err != nil {
         fmt.Println("yamlFile.Get err: ", err)
         panic(err)
     }

     err = yaml.Unmarshal(yamlFile, fcfg)
     if err != nil {
        fmt.Println("Unmarshal error: ", err)
        panic(err)
     }
     fmt.Println("Load config: ",fcfg.Version, ", Released: ", fcfg.Released)
     return fcfg
}

func LoadKey(root string,addr string) []byte {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
               files = append(files, path)
               return nil
           })
    if err != nil {
         panic(err)
    }
    for _, file := range files {
         fmt.Println("File:", file)
         list := strings.Split(file,"--")
         if len(list) == 3 {
					   account := list[2]
						 if account == strings.TrimPrefix(addr,"0x") {
							 keyjson, err := ioutil.ReadFile(file)
							 if err != nil {
										fmt.Println("Error in read file: ", file )
										return nil
							 }
							 return keyjson
						 }
         }
    }
		return nil
}

func LoadJwtAccountsFromFile(filePath string) error {
    // filePath := cf.Jwt.AccountFile
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return err
    }
    yamlFile, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("Read file: "+ filePath + " error: " + err.Error())
        return err
    }
    accountList := &JwtAccountList{}
    err = yaml.Unmarshal(yamlFile, accountList)
    if err != nil {
        fmt.Println("Unmarshal: " +  err.Error())
        return err
    }
    for _, account :=  range accountList.Accounts {
        err := Redis.Set("jwtaccount."+account.Username, string(account.Password),0).Err()
        if err != nil {
            fmt.Println(time.Now()," Cannot update user: ", account.Username)
            return err
        }
    }
    return nil
}

func LoadAccountsFromKeyStore(root string) error {
	    //root := cf.Keys.KeyStore
	    fmt.Println("Start load accounts from keystores: ",root)
	    var files []string
	    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	               files = append(files, path)
	               return nil
	           })
	    if err != nil {
	        fmt.Println("Cannot find file in folder: ", root)
	         panic(err)
	    }

	    for _, file := range files {
	         fmt.Println("File:", file)
	         list := strings.Split(file,"--")
	         if len(list) != 3 {
	            fmt.Println("File name not correct format:", file)
	            continue
	         }
	         address := list[2]

           keyjson, err := ioutil.ReadFile(file)
           if err != nil {
                fmt.Println("Error in read file: ", file )
                continue
           }
           fmt.Println("Decrypt keystore of address: ", address)
           //Store account private key
           accountKey, err := keystore.DecryptKey( []byte(keyjson), Configuration.Keys.Password)
           if err != nil {
               fmt.Println("Cannot decrypt key file: ", err)
               continue
           }
           privateKey :=  hex.EncodeToString(crypto.FromECDSA(accountKey.PrivateKey))
           account := &TokenAccount{
             Address: address,
             PrivateKey: privateKey,
           }
           value, err := json.Marshal(account)
           if err != nil {
               fmt.Println("Cannot convert to json: ", err)
               continue
           }
           err = Redis.Set("token." + address,string(value), 0).Err()
           if err != nil {
             fmt.Println("Cannot save to redis : ", err)
           }
	    }
	    fmt.Println("End load accounts from keystores: ",root)
      return nil
}

func LoadAccountsFromFile(filePath string) error {
    // filePath := cf.Jwt.AccountFile
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return err
    }
    yamlFile, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("Read file: "+ filePath + " error: " + err.Error())
        return err
    }
    accountList := []TokenAccount{}
    err = yaml.Unmarshal(yamlFile, &accountList)
    if err != nil {
        fmt.Println("Unmarshal: " +  err.Error())
        return err
    }
    for _, account :=  range accountList {
      value, err := json.Marshal(account)
      if err != nil {
          fmt.Println("Cannot convert to json: ", err)
          continue
      }
      err = Redis.Set("token." + account.Address,string(value), 0).Err()
      if err != nil {
        fmt.Println("Cannot save to redis : ", err)
      }
    }
    fmt.Println("End load accounts from file: ",filePath)
    return nil
}
func SaveAccountsToFile(filePath string) error {
      tokens, err  := Redis.Keys("token.*").Result()
      if err != nil {
        // handle error
        fmt.Println(" Cannot get addresses ")
        return err
      }
      accounts := []TokenAccount{}
      for _, token := range tokens {
        val, err := Redis.Get(token).Result()
        if err != nil {
            fmt.Println(time.Now()," Cannot find token: ", token)
            continue
        }
        account := TokenAccount{}
        err = json.Unmarshal([]byte(val), &account)
        if err != nil {
            fmt.Println(time.Now()," Cannot parse account ", err)
            continue
        }
        accounts = append(accounts,account)
      }

      accountlist, err := yaml.Marshal(accounts)
      if err != nil {
          fmt.Println("Json parse accounts error: : ",err)
          return err
      }
      err = ioutil.WriteFile(filePath, accountlist, 0644)
      if err != nil {
        fmt.Println("Cannot save to file ",err)
        return err
      }
      fmt.Println("End save accounts to file: ",filePath)
      return nil
}
