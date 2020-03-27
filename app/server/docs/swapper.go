package docs

import (
  "github.com/binhnt-teko/test_loyalty/app/server/config"
  )

// This return default
// swagger:response commonResponse
type commonResponseWrapper struct {
	// in:body
 	Body config.ApiResponse
}
