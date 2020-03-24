package config

import (
   "testing"
   "fmt"
 	 "github.com/joho/godotenv"
   // "os"
)

func TestConfig(t *testing.T) {
   config_file := "../../config/config.yaml"
   cfg = LoadConfigFromFile(config_file)
   // fmt.Print(cfg.toJson())
   if len(cfg.Networks) < 2 {
      t.Errorf("Config need more nodes ")
   }
}
func TestConfigRedis(t *testing.T) {
  // cwd, err := os.Getwd()
  // t.Log(err)
  // t.Log(cwd)
  godotenv.Load()
  cfg = LoadConfig()
  fmt.Print(cfg.toJson())
}
