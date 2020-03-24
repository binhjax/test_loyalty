package config

import (
    "fmt"
    "encoding/json"
    "os"
    "github.com/go-redis/redis"
    "strconv"
    "time"
)
var Configuration *Config
var Redis  *redis.Client


func Init() {
    config_type := os.Getenv("CONFIG_TYPE")
    if len(config_type) == 0 {
      config_type = "file"
    }

    config_file := os.Getenv("CONFIG_FILE")
    if len(config_file) == 0 {
       config_file = "../../config/config.yaml"
    }
    redis_server := os.Getenv("REDIS_SERVER")
    if len(redis_server) == 0 {
       redis_server = "localhost:6379"
    }
    redis_password := os.Getenv("REDIS_PASSWORD")
    if len(redis_password) == 0 {
       redis_password = ""
    }
    redis_db := os.Getenv("REDIS_DB")
    if len(redis_db) == 0 {
       redis_db = "0"
    }
    db, err := strconv.Atoi(redis_db)
  	if err != nil {
  		  db = 0
  	}

    if config_type == "redis" {
        client := redis.NewClient(&redis.Options{
            Addr:     redis_server,
            Password: redis_password, // no password set
            DB:      db,  // use default DB
          })
         _ , err := client.Ping().Result()
        if  err != nil {
          panic(err)
        }
        val, err := client.Get("config").Result()
      	if err != nil {
      		 panic(err)
      	}
        config := &Config{}
        err1 := json.Unmarshal([]byte(val), config)
        if err1 != nil {
            fmt.Println(time.Now()," Cannot parse data ", err)
            panic(err1)
        }
        Configuration = config
        Redis = client
        return
    }
    //Load config file
    fileCfg := LoadConfigFromFile(config_file)

    //Try to connect redis
    client := redis.NewClient(&redis.Options{
      Addr:     fileCfg.Redis.Host,
      Password: fileCfg.Redis.Password, // no password set
      DB:       fileCfg.Redis.Db,  // use default DB
    })
    pong, err := client.Ping().Result()
    if  err == nil {
         fmt.Println("Update config to redis.")
         value, err := json.Marshal(fileCfg)
         if err != nil {
             fmt.Println(err)
         }
         err = client.Set("config",string(value), 0).Err()
         if err != nil {
            fmt.Println(time.Now()," Write config to redis error: ", err)
         }
    } else {
      fmt.Println(pong, err)
    }
    Configuration = fileCfg
    Redis = client
    return
}
