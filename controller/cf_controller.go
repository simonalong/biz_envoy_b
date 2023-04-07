package controller

import (
	"biz-b-service/service"
	"github.com/isyscore/isc-gobase/server"
)

func CfController() {
	// 测试入口和出口的拦截情况
	server.Get("front/bdef/:haveRedis/:haveMysql", service.FrontBdef)
	server.Get("front/bcf/:haveRedis/:haveMysql", service.FrontBcf)
	server.Get("front/bef/:haveRedis/:haveMysql", service.FrontBef)
	server.Get("front/badef/:haveRedis/:haveMysql", service.FrontBadef)
}
