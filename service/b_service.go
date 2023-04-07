package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/isyscore/isc-gobase/config"
	baseHttp "github.com/isyscore/isc-gobase/http"
	"github.com/isyscore/isc-gobase/logger"
	"github.com/isyscore/isc-gobase/redis"
	"github.com/isyscore/isc-gobase/server/rsp"
	baseTime "github.com/isyscore/isc-gobase/time"
	"net/http"
	"time"
)

func FrontBdef(c *gin.Context) {
	//logger.Info("接口 FrontCfOkOk ")
	// x-request-id
	// x-b3-traceid
	// x-b3-spanid
	// x-b3-parentspanid
	// x-b3-sampled
	// x-b3-flags
	// x-ot-span-context

	headers := http.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}

	haveRedis := c.Param("haveRedis")
	haveMysql := c.Param("haveMysql")

	data, err := baseHttp.GetOfStandard(config.GetValueString("biz.url.d")+"/api/d/bd/"+haveMysql, headers, map[string]string{})
	if err != nil {
		logger.Error("接口 FrontCfOkOk 报错", err)
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}

	if haveRedis == "true" {
		rdb, _ := redis.GetClient()
		ctx := context.Background()
		rdb.Set(ctx, "FrontBdef", baseTime.TimeToStringYmdHms(time.Now()), time.Minute*1)
	}

	rsp.SuccessOfStandard(c, data)
}

func FrontBcf(c *gin.Context) {
	//logger.Info("接口 FrontCfOkStop ")

	headers := http.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}

	haveRedis := c.Param("haveRedis")
	haveMysql := c.Param("haveMysql")
	data, err := baseHttp.GetOfStandard(config.GetValueString("biz.url.c")+"/api/c/bc/"+haveRedis+"/"+haveMysql, headers, map[string]string{})
	if err != nil {
		logger.Error("接口 FrontCfOkStop 报错", err)
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}

	if haveRedis == "true" {
		rdb, _ := redis.GetClient()
		ctx := context.Background()
		rdb.Set(ctx, "FrontBcf", baseTime.TimeToStringYmdHms(time.Now()), time.Minute*1)
	}

	rsp.SuccessOfStandard(c, data)
}

func FrontBef(c *gin.Context) {
	//logger.Info("接口 FrontCfStopOk ")
	headers := http.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}

	haveRedis := c.Param("haveRedis")
	haveMysql := c.Param("haveMysql")
	data, err := baseHttp.GetOfStandard(config.GetValueString("biz.url.e")+"/api/e/be/"+haveMysql, headers, map[string]string{})
	if err != nil {
		logger.Error("接口 FrontCfStopOk 报错", err)
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}

	if haveRedis == "true" {
		rdb, _ := redis.GetClient()
		ctx := context.Background()
		rdb.Set(ctx, "FrontBcf", baseTime.TimeToStringYmdHms(time.Now()), time.Minute*1)
	}
	rsp.SuccessOfStandard(c, data)
}

func FrontBadef(c *gin.Context) {
	//logger.Info("接口 FrontCfStopOk ")
	headers := http.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}

	haveRedis := c.Param("haveRedis")
	haveMysql := c.Param("haveMysql")
	data, err := baseHttp.GetOfStandard(config.GetValueString("biz.url.a")+"/api/a/ba/"+haveRedis+"/"+haveMysql, headers, map[string]string{})
	if err != nil {
		logger.Error("接口 FrontCfStopOk 报错", err)
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}

	if haveRedis == "true" {
		rdb, _ := redis.GetClient()
		ctx := context.Background()
		rdb.Set(ctx, "FrontBadef", baseTime.TimeToStringYmdHms(time.Now()), time.Minute*1)
	}
	rsp.SuccessOfStandard(c, data)
}
