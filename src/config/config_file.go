package config

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "strings"
  "time"
  "encoding/json"
)

type SpecialDate struct {
    time.Time
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
    Mysql struct {
        Host string `yaml:"host" json:"host"`
        Port string `yaml:"port" json:"port"`
        Username string `yaml:"username" json:"username"`
        Password string `yaml:"password" json:"password"`
        Database string `yaml:"database" json:"database"`
        Debug bool `yaml:"debug" json:"debug"`
    } `yaml:"mysql" json:"mysql"`
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
		Keys  struct {
        LoadKey  bool `yaml:"loadKey" json:"loadKey"`
			  KeyStore string `yaml:"keystore" json:"keystore"`
				Password string `yaml:"password" json:"password"`
		} `yaml:"keys" json:"keys"`
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
		Contract struct {
        GasPrice string `yaml:"gasprice" json:"gasprice"`
        GasLimit uint64 `yaml:"gaslimit" json:"gaslimit"`
				Owner string `yaml:"owner" json:"owner"`
				InitialToken int64 `yaml:"initialToken" json:"initialToken"`
				MasterKey1 string `yaml:"masterkey1" json:"masterkey1"`
				MasterKey2 string `yaml:"masterkey2" json:"masterkey2"`
				Address string `yaml:"address" json:"address"`
		} `yaml:"contract" json:"contract"`
    F5Contract struct {
        GasPrice string `yaml:"gasprice" json:"gasprice"`
        GasLimitDefault uint64 `yaml:"gaslimitdefault" json:"gaslimitdefault"`
        GasLimit map[string]uint64 `yaml:"gaslimit" json:"gaslimit"`
        EthBudget string `yaml:"ethBudget" json:"ethBudget"`
        Owner string `yaml:"owner" json:"owner"`
      	Address string `yaml:"address" json:"address"`
		} `yaml:"f5contract" json:"f5contract"`
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
func (cf *Config) MysqlConnectionUrl() string {
    return cf.Mysql.Username + ":" + cf.Mysql.Password + "@tcp(" +  cf.Mysql.Host + ":" + cf.Mysql.Port + ")/" +   cf.Mysql.Database + "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True"
}

func LoadConfigFromFile(file string) *Config {
     var fcfg = &Config{}

     yamlFile, err := ioutil.ReadFile(file)
     if err != nil {
         fmt.Println("yamlFile.Get err: ", err)
     }

     err = yaml.Unmarshal(yamlFile, fcfg)
     if err != nil {
         fmt.Println("Unmarshal error: ", err)
     }
     fmt.Println("Load config: ",fcfg.Version, ", Released: ", fcfg.Released)
     return fcfg
}
