package helper

import (
   "testing"
   "fmt"
   "github.com/binhnt-teko/test_loyalty/src/config"
)

func TestConfig(*testing.T) {
   text := "GetMD5Hash"
   hash := GetMD5Hash(text)
   fmt.Print(hash)
   config_file := "../../config/config.yaml"
   cfg := config.LoadConfig(config_file)
   LoadJwtAccounts(cfg)
   LoadKeyStoresToDB(cfg)
}
