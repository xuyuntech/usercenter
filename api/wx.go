package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// get		查询
// post		增加
// put		修改
// delete	删除
func (a *Api) wxCallback(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		RespErr(c, fmt.Errorf("need param key"))
		return
	}

	Resp(c, map[string]string{
		"results": "123123",
	})
}
