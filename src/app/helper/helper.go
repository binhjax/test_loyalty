package helper

import (
	"fmt"
  "os"
	"strings"
  "github.com/binhnt-teko/test_loyalty/src/config"
  "crypto/md5"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v2"
  "io/ioutil"
	wa "github.com/binhnt-teko/test_loyalty/src/app/account"
)

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func JwtVerifyUsername(username string,password string) bool {
	val, err := config.Redis.Get("jwtaccount." + username).Result()
	if err != nil {
			fmt.Println(time.Now()," Cannot find jwtaccount: ", username)
			return false
	}
	if val != password {
		fmt.Println(time.Now()," Password is not correct: ", username)
		return false
	}
	return true
}
