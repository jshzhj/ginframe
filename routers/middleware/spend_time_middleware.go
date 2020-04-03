package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jshzhj/ginframe/pkg/logger"
	"time"
)

func SpendTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		fmt.Printf("该请求花费时间:%v", latency)
		logger.Info("zap:http日志",c.Request.RequestURI,200,latency,1,123)
	}
}
