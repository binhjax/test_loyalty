package middleware

import (
	"errors"
	"fmt"
	// "github.com/valyala/fasthttp"
	"github.com/qiangxue/fasthttp-routing"
  "github.com/binhnt-teko/test_loyalty/app/server/config"
	"time"
)

func CheckTokenMiddleware(c *routing.Context) error {
	if config.Configuration.Jwt.Enable {
		fasthttpJwtCookie := c.Request.Header.Cookie("fasthttp_jwt")

		if len(fasthttpJwtCookie) == 0 {
			return errors.New("login required")
		}
		fmt.Println(string(fasthttpJwtCookie))
		token, _, err := JWTValidate(string(fasthttpJwtCookie))

		if !token.Valid {
			return errors.New("Token is invalid, login again please")
		}

		if claims, ok := token.Claims.(*UserCredential) ; !ok  {
			return  errors.New("Token is invalid, cannot get user info")
		} else {
			if time.Now().Unix() > claims.ExpiresAt {
				return errors.New("Token is expired, login again please")
			}
			return err
		}

	}
	return nil
}

type MiddlewareType func(c *routing.Context) error

var MiddlewareList = []MiddlewareType{
	CheckTokenMiddleware,
}

// BasicAuth is the basic auth handler
func JWTMiddleware() routing.Handler {
	return func(c *routing.Context) error {
		// for _, middleware_item := range MiddlewareList {
		// 	if err := middleware_item(c); err != nil {
		// 		res := &config.ApiResponse{
		// 				Rescode: 99,
		// 				Resdecr: "Please login",
		// 				Resdata: err.Error(),
		// 		}
		// 		fmt.Fprintf(c,res.ToJson())
		// 		c.Abort()
		// 		return nil
		// 	}
		// }
		c.Next()
		return nil
	}
}
