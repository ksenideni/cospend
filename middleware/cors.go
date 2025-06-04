package middleware

import (
	"cospend/constant"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		wHead := c.Writer.Header()
		wHead.Set(constant.ACAO, constant.ACAO_VALUE)
		wHead.Set(constant.ACAM, constant.ACAM_VALUE)
		wHead.Set(constant.ACAH, constant.ACAH_VALUE)
		wHead.Set(constant.ACAC, constant.ACAC_VALUE)
		wHead.Set(constant.HSTS, constant.HSTS_VALUE)
		wHead.Set(constant.CC, constant.CC_VALUE)
		wHead.Set(constant.XCTO, constant.XCTO_VALUE)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
