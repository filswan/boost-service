package main

import (
	"fmt"
	"github/filswan/boost-service/api"
	"github/filswan/boost-service/config"
	"github/filswan/boost-service/log"
	"github/filswan/boost-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	conf := config.Conf()
	log.Init(conf.Log.Level)
	if err := service.Init(conf.Boost.FullAPI, conf.Boost.Repo); err != nil {
		panic(err)
	}
	Router := gin.Default()
	boostApi := new(api.BoostApi)
	Router.POST("/rpc", boostApi.RPCMethod)
	Router.Run(fmt.Sprintf(":%d", config.Conf().Port))
}
