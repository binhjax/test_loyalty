package config

import (
	// "fmt"
  // "strings"
  // "github.com/go-redis/redis"
  "encoding/json"
)

type ApiResponse struct  {
	  Rescode int  `json:"rescode"`
		Resdecr string  `json:"resdecr"`
		Resdata interface{}  `json:"resdata"`
}

func (res *ApiResponse) ToJson() string {
	  b, err := json.Marshal(res)
    if err != nil {
        //fmt.Println("Json parse error: : ",err)
        return ""
    }
		//fmt.Println("Json parse ok: : ",string(b))
    return string(b)
}
