package config

import (
   "testing"
   "fmt"
)

func TestConfig(*testing.T) {
   config_file := "./config.yaml"
   cfg = LoadConfigFromFile(config_file)
   fmt.Print(cfg.toJson())
   fmt.Print(cfg.MysqlConnectionUrl())
   // if total != 10 {
   //    t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
   // }
   cfg = LoadConfig(config_file)
}
