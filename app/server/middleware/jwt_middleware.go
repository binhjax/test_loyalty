package middleware

import (
	"errors"
	"fmt"
	// "github.com/valyala/fasthttp"
	"github.com/qiangxue/fasthttp-routing"
  "github.com/binhnt-teko/test_loyalty/app/server/config"
)

func CheckTokenMiddleware(c *routing.Context) error {
	if config.Configuration.Jwt.Enable {
		fasthttpJwtCookie := c.Request.Header.Cookie("fasthttp_jwt")

		if len(fasthttpJwtCookie) == 0 {
			return errors.New("login required")
		}

		token, _, err := JWTValidate(string(fasthttpJwtCookie))

		if !token.Valid {
			return errors.New("your session is expired, login again please")
		}
		return err
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
		for _, middleware_item := range MiddlewareList {
			if err := middleware_item(c); err != nil {
				res := &config.ApiResponse{
						Rescode: 99,
						Resdecr: "Please login",
						Resdata: err.Error(),
				}
				fmt.Fprintf(c,res.ToJson())
				return nil
			}
		}
		c.Next()
		return nil
	}
}
