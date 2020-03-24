package config

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "encoding/json"
)
var cfg *Config

type Setting struct {
  gorm.Model
  Version string
  Config string  `sql:"type:text"`
  Description string  `sql:"type:text"`
  Active bool `gorm:"default:false"`
}

func (s *Setting) GetConfig() *Config {
  cf := &Config{}
  err := json.Unmarshal([]byte(s.Config), cf)
  if err != nil {
      fmt.Println("Cannot parse string to Config", err.Error())
      return nil
  }
  return cf
}

func LoadConfig(file string) *Config {
    fileCfg := LoadConfigFromFile(file)

    //fmt.Println("Config: ",fileCfg.toJson())
    //Create mysql connection
    db, err := gorm.Open("mysql", fileCfg.MysqlConnectionUrl())
    if fileCfg.Mysql.Debug {
       db.LogMode(true)
    }

    if err != nil {
      panic("failed to connect database: " + err.Error())
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&Setting{})

    setting := &Setting{}
    // find setting with version
    if err := db.Where("version = ?", fileCfg.Version).First(setting).Error; err != nil {
        fmt.Println("version: ",fileCfg.Version," not existed. Update new version")
        setting := Setting{
          Version: fileCfg.Version,
          Config: string(fileCfg.toJson()),
          Active: false,
        }
        //fmt.Println("Create new record")
        db.Create(&setting)
    } else {
       if fileCfg.Update {
           //Update current record
          setting.Config = string(fileCfg.toJson())
          db.Save(&setting)
          fmt.Println("version: ",fileCfg.Version," update successfully")
       }
    }
    //Load First active setting
    fmt.Println("Load first active record")
    activeSetting := &Setting{}
    if err := db.Where("active = ?", true).First(activeSetting).Error; err != nil {
        fmt.Println("Cannot find active config. Load first config:")
        db.First(activeSetting)
    }

    if activeSetting != nil {
      fmt.Println("Parse json to Config struct")
      //Load config
      cfg = activeSetting.GetConfig()
      return cfg
    }

    panic("Cannot load config ")
    return nil
}
