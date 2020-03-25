package api

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
  "github.com/binhnt-teko/test_loyalty/app/server/middleware"
  "github.com/binhnt-teko/test_loyalty/app/server/config"
)
type ApiAdmin struct {}

func (api ApiAdmin) LoadJwtAccounts(c *routing.Context) error {

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

func (api ApiAdmin) LoadAccounts(c  *routing.Context) error {

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
}
