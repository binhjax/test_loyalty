package config

import (
  "fmt"
  "strings"
  "time"
  "encoding/json"
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

type SpecialDate struct {
    time.Time
}

type Config struct {
    Version string `yaml:"version" json:"version"`
    Released SpecialDate `yaml:"released" json:"released"`
    Update bool `yaml:"update" json:"update"`
    Jwt struct {
        Enable bool `yaml:"enable" json:"enable"`
        Signkey string `yaml:"signkey" json:"signkey"`
        ExpiredAt int `yaml:"expiredAt" json:"expiredAt"`
        LoadAccount bool `yaml:"loadAccount" json:"loadAccount"`
        AccountFile string `yaml:"accountFile" json:"accountFile"`
    } `yaml:"jwt" json:"jwt"`
    Channel struct {
        TransferQueue int `yaml:"transferqueue" json:"transferqueue"`
        LogQueue int `yaml:"logqueue" json:"logqueue"`
    } `yaml:"channel" json:"channel"`
    Webserver struct {
			  Port string `yaml:"port" json:"port"`
        Tls bool `yaml:"tls" json:"tls"`
        CertificateFile string `yaml:"certificateFile" json:"certificateFile"`
        KeyFile string `yaml:"keyFile" json:"keyFile"`
        MaxRpcConnection int `yaml:"maxrpc" json:"maxrpc"`
        MaxListenRpcConnection int `yaml:"maxlistenrpc" json:"maxlistenrpc"`
        RoutingMode int `yaml:"routingMode" json:"routingMode"`
        NonceMode int `yaml:"nonceMode" json:"nonceMode"`
		} `yaml:"webserver" json:"webserver"`
    Networks []struct {
        Name string `yaml:"name" json:"name"`
        Http string `yaml:"http" json:"http"`
        WebSocket string `yaml:"websocket" json:"websocket"`
        LocalAddr string `yaml:"local" json:"local"`
    } `yaml:"networks" json:"networks"`
		Redis struct {
        MaxConn int  `yaml:"maxconn" json:"maxconn"`
			  Host string `yaml:"host" json:"host"`
			  Password string `yaml:"password" json:"password"`
			  Db int `yaml:"db" json:"db"`
		} `yaml:"redis" json:"redis"`
}

func (sd *SpecialDate) UnmarshalYAML(unmarshal func(interface{}) error) error {
    fmt.Println("SpecialDate.UnmarshalYAML: parse date")
    var input string
    if err := unmarshal(&input); err != nil {
        fmt.Println("SpecialDate.UnmarshalYAML: Cannot parse string")
        return err
    }
    strInput := strings.Trim(input, `"`)
    newTime, err := time.Parse("2006-01-01T00:00:00Z", strInput)
    if err != nil {
      fmt.Println("SpecialDate.UnmarshalYAML: parse time error: ",err)
        return err
    }
    sd.Time = newTime
    return nil
}

func (cf *Config) toJson() string {
    fmt.Println("Convert Config struct to json: ")
    b, err := json.Marshal(cf)
    if err != nil {
        fmt.Println(err)
        return "{}"
    }
    return string(b)
}
