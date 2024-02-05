package api

import (
	"errors"
	"net/http"

	"github/filswan/boost-service/service"

	"github.com/gin-gonic/gin"
)

type BoostApi struct {
	BaseApi
}

func (api *BoostApi) RPCMethod(c *gin.Context) {
	var req JSONRPCReq
	if err := api.ParseReq(c, &req); err != nil {
		return
	}
	if req.Method != "Boost.ProviderStorageAsk" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	id := req.ID
	if len(req.Params) < 1 {
		api.ErrResponse(c, -1, errors.New("params required"), id)
		return
	}
	provider, ok := req.Params[0].(string)
	if !ok {
		api.ErrResponse(c, -1, errors.New("invalid para type"), id)
		return
	}
	askInfo, err := service.BoostService.ProviderStorageAsk(provider)
	if err != nil {
		api.ErrResponse(c, -1, err, id)
		return
	}
	api.Response(c, askInfo, id)
}
