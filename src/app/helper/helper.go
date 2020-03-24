package helper

import (
	"fmt"
  "os"
	"strings"
  "github.com/binhnt-teko/test_loyalty/src/config"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/ethereum/go-ethereum/accounts/keystore"
	 "path/filepath"
	"crypto/md5"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v2"
  "io/ioutil"
	wa "github.com/binhnt-teko/test_loyalty/src/account"
)
type JwtAccount struct {
  gorm.Model
  Username string `json:"username"`
	Password string `json:"password"`
  Active bool `gorm:"default:false"`
}

type JwtAccountList struct {
  Accounts []struct {
    Username string `yaml:"username" json:"username"`
    Password string `yaml:"password" json:"password"`
  } `yaml:"accounts" json:"accounts"`
}


func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func JwtVerifyUsername(cfg *config.Config, username string,password string) bool {
    db, err := gorm.Open("mysql", cfg.MysqlConnectionUrl())
    if cfg.Mysql.Debug {
       db.LogMode(true)
    }

    if err != nil {
      panic("failed to connect database: " + err.Error())
    }
    defer db.Close()
    acc := &JwtAccount{}
    fmt.Println("Start search username: ")
    if err := db.Where("username = ?", username).Where("password = ?", password).Where("active = ?", true).First(acc).Error; err != nil {
      fmt.Println("username: ",username," not existed or password failed or not activate")
      return false
  }
   return true
}
func LoadJwtAccounts(cf *config.Config)  {
    filePath := cf.Jwt.AccountFile
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        panic("File: "+ filePath + " not existed")
    }
    yamlFile, err := ioutil.ReadFile(filePath)
    if err != nil {
           panic("Read file: "+ filePath + " error: " + err.Error())
    }
    accountList := &JwtAccountList{}
    err = yaml.Unmarshal(yamlFile, accountList)
    if err != nil {
        panic("Unmarshal: " +  err.Error())
    }

    //Create connection to mysql
    db, err := gorm.Open("mysql", cf.MysqlConnectionUrl())
    if cf.Mysql.Debug {
       db.LogMode(true)
    }

    if err != nil {
      panic("failed to connect database: " + err.Error())
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&JwtAccount{})

    for _, account :=  range accountList.Accounts {
        acc := JwtAccount{}
       //Check db if not exist with username
       if err := db.Where("username = ?", account.Username).First(&acc).Error; err != nil {
           fmt.Println("Jwt Account: ",account.Username," not existed. Update new Jwt account: " + err.Error())
           user := JwtAccount{
             Username: account.Username,
             Password: GetMD5Hash(account.Password),
             Active: false,
           }
           //fmt.Println("Create new record")
           db.Create(&user)
       }
    }
}

func LoadKeyStoresToDB(cf *config.Config){

    root := cf.Keys.KeyStore
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

    fmt.Println("Connect to mysql: ",cf.MysqlConnectionUrl())
    //Create connect to db
    db, err := gorm.Open("mysql", cf.MysqlConnectionUrl())
    if cf.Mysql.Debug {
       db.LogMode(true)
    }

    if err != nil {
      panic("failed to connect database: " + err.Error())
    }
    defer db.Close()

    // Migrate the schema
    db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")
    db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&wa.TokenAccount{})


    for _, file := range files {
         fmt.Println("File:", file)
         list := strings.Split(file,"--")
         if len(list) != 3 {
            fmt.Println("File name not correct format:", file)
            continue
         }
         account := list[2]

         token := &wa.TokenAccount{}
        // find token with address
         if err := db.Where("address = ?", account).First(token).Error; err != nil {
             fmt.Println("Not find account with address: ", account," Start read file")

             //Load account in db
             keyjson, err := ioutil.ReadFile(file)
             if err != nil {
                  fmt.Println("Error in read file: ", file )
                  continue
             }
             fmt.Println("Decrypt keystore of account: ", account)
             //Store account private key
             accountKey, err := keystore.DecryptKey( []byte(keyjson), cf.Keys.Password)
             if err != nil {
                 fmt.Println("Cannot decrypt key file: ", err)
                 continue
             }
             privateKey :=  hex.EncodeToString(crypto.FromECDSA(accountKey.PrivateKey))
             new_account := &wa.TokenAccount{
               Address: account,
               PrivateKey: privateKey,
               Active: true,
             }
             //fmt.Println("Create new record")
             db.Create(new_account)
         }
    }
    fmt.Println("End load accounts from keystores: ",root)
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
