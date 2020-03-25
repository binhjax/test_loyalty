package helper

import (
   "testing"
   "fmt"
   "github.com/binhnt-teko/test_loyalty/src/config"
)

func TestGetMD5Hash(*testing.T) {
   text := "GetMD5Hash"
   hash := GetMD5Hash(text)
   fmt.Print(hash)
}
func TestJwtVerifyUsername(*testing.T) {
   config.Init()
   if JwtVerifyUsername("binhnt","123456") {
     fmt.Println("Existed user")
   }
   else {
      fmt.Println("Not find existed user")
   }
}
