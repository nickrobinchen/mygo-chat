package logic

import (
	"fmt"
	"mygochat/config"
	"runtime"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Logic struct {
	ServerId string
}

func New() *Logic {
	return new(Logic)
}

func (logic *Logic) Run() {
	//read config
	logicConfig := config.Conf.Logic

	runtime.GOMAXPROCS(logicConfig.LogicBase.CpuNum)
	logic.ServerId = fmt.Sprintf("logic-%s", uuid.New().String())
	//init publish redis
	if err := logic.InitPublishRedisClient(); err != nil {
		logrus.Panicf("logic init publishRedisClient fail,err:%s", err.Error())
	}

	//init rpc server
	if err := logic.InitRpcServer(); err != nil {
		logrus.Panicf("logic init rpc server fail")
	}
}
