package controllers

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
  "github.com/binhnt-teko/test_loyalty/app/server/middleware"
  "github.com/binhnt-teko/test_loyalty/app/server/helper"
	"github.com/binhnt-teko/test_loyalty/app/server/config"
	"time"
	"encoding/json"
)
type ApiAuth struct {}

// JSON Authentication
func (api ApiAuth)  Login(c  *routing.Context) error {
	fmt.Println("API Access Login. ")

	var jwtaccount = &config.JwtAccount{}
	val := c.PostBody()
	err := json.Unmarshal([]byte(val), jwtaccount)
	if err != nil {
			fmt.Println(time.Now()," Cannot parse account ", err)
			return err
	}

	username := jwtaccount.Username
	password := jwtaccount.Password

	fmt.Println("Username:  " + username  + ", Password: " + password)

	if helper.JwtVerifyUsername(username,password) {
			fmt.Println("Start JwtVerifyUsername")
			qUser := []byte(username)
			qPasswd := []byte(password)

			fasthttpJwtCookie := c.Request.Header.Cookie("fasthttp_jwt")

			// for example, server receive token string in request header.
			if len(fasthttpJwtCookie) > 0 {
					fmt.Println("Get Cookie:  "+ string(fasthttpJwtCookie))
					token, _, _ := middleware.JWTValidate(string(fasthttpJwtCookie))
					if token.Valid {
						return Response(c,98,"Existed JWT token, no need recreate",nil)
					} else {
						fmt.Println("Session key is not valid:  "+ string(fasthttpJwtCookie))
					}
			}
			fmt.Println("Start creating token ")

			tokenString, expireAt := middleware.CreateToken(qUser, qPasswd)
			//fmt.Println("Get cookied ")
			// Set cookie for domain
			cookie := fasthttp.AcquireCookie()
			//fmt.Println("End get cookie ")

			cookie.SetKey("fasthttp_jwt")

			cookie.SetValue(tokenString)
			cookie.SetExpire(expireAt)
			c.Response.Header.SetCookie(cookie)
			fmt.Println("End creating token ")
			data := make(map[string]string)
			data["fasthttp_jwt"] = tokenString
			data["expireAt"] = expireAt.String()
			// jsonString, _ := json.Marshal(datas)
			return Response(c,0,"New Jwt token successfully", data)
	}
	return Response(c,99,"Username or password failed",nil)
}

func (api ApiAuth) LoadJwtAccount(c  *routing.Context) error {
		var jwtfile = &FilePath{}
		err := json.Unmarshal([]byte(c.PostBody()), jwtfile)
		if err != nil {
				fmt.Println(time.Now()," Cannot parse filepath ", err)
				Response(c,99,"Cannot parse filepath",err)
				return err
		}

		if err := config.LoadJwtAccountsFromFile(jwtfile.Filepath); err != nil {
				fmt.Println(time.Now()," Cannot load jwt account ", err)
				Response(c,99,"Load file failed",err)
		}
		return Response(c,0,"Load file finished",nil)
}
