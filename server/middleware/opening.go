package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/common/response"
	"github.com/springbear2020/self-hub/server/utils"
)

type OpenTokenInfo struct {
	UID      uint   `json:"uid"`
	Username string `json:"username"`
}

const openTokenKey = "x-open-token"

func OpenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(openTokenKey)
		if token == "" {
			response.NoAuth("非法访问，令牌缺失", c)
			c.Abort()
			return
		}

		tokenJSONBytes, err := utils.AesGCMDecrypt(token, []byte(global.GVA_CONFIG.JWT.SigningKey))
		if err != nil {
			response.NoAuth("非法访问，解密失败", c)
			c.Abort()
			return
		}

		var tokenInfo OpenTokenInfo
		err = json.Unmarshal(tokenJSONBytes, &tokenInfo)
		if err != nil {
			response.NoAuth("非法访问，解析失败", c)
			c.Abort()
			return
		}

		if tokenInfo.UID <= 0 || tokenInfo.Username == "" {
			response.NoAuth("非法访问，字段非法", c)
			c.Abort()
			return
		}

		c.Set("uid", tokenInfo.UID)
		c.Set("username", tokenInfo.Username)

		c.Next()
	}
}
