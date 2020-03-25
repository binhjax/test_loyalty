package api

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
  "github.com/binhnt-teko/test_loyalty/app/server/middleware"
  "github.com/binhnt-teko/test_loyalty/app/server/helper"
)
type ApiAuth struct {}

func (api ApiAuth) Login_get(c  *routing.Context) error {
		fmt.Println("GET: API Access Login")
		return Response(c,99,"Please use POST method",nil)
}

// JSON Authentication
func (api ApiAuth)  Login(c  *routing.Context) error {
	fmt.Println("API Access Login. ")

	username := string(c.FormValue("username"))
	password := string(c.FormValue("password"))

	fmt.Println("Username:  " + username  + ", " + password)

	if helper.JwtVerifyUsername(username,password) {
			fmt.Println("Start JwtVerifyUsername")
			qUser := []byte(username)
			qPasswd := []byte(password)

			fasthttpJwtCookie := c.Request.Header.Cookie("fasthttp_jwt")

			// for example, server receive token string in request header.
			if len(fasthttpJwtCookie) == 0 {
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
				return Response(c,0,"New Jwt token successfully",nil)
			}
			return Response(c,98,"Existed JWT token, no need recreate",nil)
	}
	return Response(c,99,"Username or password failed",nil)
}
