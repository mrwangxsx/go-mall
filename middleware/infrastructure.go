package middleware

import (
	"github.com/gin-gonic/gin"
	"go-mall/util"
)

func StartTrace() gin.HandlerFunc {

	return func(c *gin.Context) {
		traceId := c.Request.Header.Get("traceid")
		pSpanId := c.Request.Header.Get("pSpanId")
		spanId := util.GenerateSpanID(c.Request.RemoteAddr)
		if traceId == "" {
			traceId = pSpanId
		}
		c.Set("traceid", traceId)
		c.Set("pSpanId", pSpanId)
		c.Set("spanid", spanId)
		c.Next()
	}
}
