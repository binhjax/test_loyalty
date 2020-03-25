package config

import (
   "testing"
   "fmt"
 	 "github.com/joho/godotenv"
   // "os"
)

func TestConfig(t *testing.T) {
   config_file := "../../config/config.yaml"
   cfg := LoadConfigFromFile(config_file)
   // fmt.Print(cfg.toJson())
   if len(cfg.Networks) < 2 {
      t.Errorf("Config need more nodes ")
   }
}
func TestConfigRedis(t *testing.T) {
  godotenv.Load()
  Init()
  fmt.Print(Configuration)
}
func TestLoadJwtAccountsFromFile(t *testing.T) {
  godotenv.Load()
  Init()
  err := LoadJwtAccountsFromFile("../../config/jwt_account.yml")
  if err != nil {
      t.Errorf("Cannot load jwt accounts ")
  }
  // fmt.Print(Configuration)
}
func TestLoadAccountsFromFile(t *testing.T) {
  godotenv.Load()
  Init()
  err := LoadAccountsFromFile("../../config/keystore")
  if err != nil {
      t.Errorf("Cannot load wallet accounts ")
  }
  // fmt.Print(Configuration)
}
