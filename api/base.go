package api

import (
	"net/http"
	"net/http/httputil"

	"github/filswan/boost-service/log"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (api *BaseApi) ParseReq(c *gin.Context, receiverPointer any) error {
	body, _ := httputil.DumpRequest(c.Request, true)
	log.Info(string(body))
	if err := c.ShouldBind(receiverPointer); err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return err
	}
	return nil
}

func (api *BaseApi) Response(c *gin.Context, data any, reqID ...int) {
	id := 1
	if len(reqID) > 0 {
		id = reqID[0]
	}
	api.response(c, JSONRPCResult{
		Jsonrpc: "2.0",
		ID:      id,
		Result:  data,
	})
}

func (api *BaseApi) response(c *gin.Context, data any, statusCode ...int) {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	c.JSON(code, data)
}

func (api *BaseApi) ErrResponse(c *gin.Context, code int, err error, reqID ...int) {
	id := 1
	if len(reqID) > 0 {
		id = reqID[0]
	}
	api.response(c, JSONRPCResult{
		Jsonrpc: "2.0",
		ID:      id,
		Error: &JSONResultError{
			Code:    code,
			Message: err.Error(),
		},
	})
}

type JSONRPCReq struct {
	Jsonrpc string `json:"jsonrpc"   binding:"required"`
	Method  string `json:"method"    binding:"required"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"        binding:"required"`
}

type JSONRPCResult struct {
	Jsonrpc string           `json:"jsonrpc"`
	ID      int              `json:"id"`
	Result  any              `json:"result,omitempty"`
	Error   *JSONResultError `json:"error,omitempty"`
}

type JSONResultError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
