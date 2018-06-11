package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xuyuntech/usercenter/manager"
)

type Api struct {
	Listen  string
	Manager manager.Manager
}

func (a *Api) Run() error {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "X-Access-Token"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	// 微信认证回调接口，用作微信开发者绑定
	r.GET("/wx_callback", a.wxCallback)

	// /auth 开头的接口需要 token 认证，也就是下列接口必须是用户登录之后调用的
	authorized := r.Group("/auth")
	// token 认证中间件
	authorized.Use(func(c *gin.Context) {})
	{
		// 绑定用户手机号
		authorized.POST("/bind_phone", a.bindPhone)
		// 获取用户基础信息
		authorized.GET("/userinfo", a.userinfo)
	}

	// TODO RBAC

	// 上传初始用户资料，上线之前需要将门店已有用户数据导入
	//    上传方式为 Excel 导入
	r.POST("/upload_users", a.uploadUsers)
	return r.Run(a.Listen)
}

func RespErr(c *gin.Context, err error, msg ...string) {
	results := map[string]interface{}{
		"status": 1,
		"err":    err.Error(),
	}
	if len(msg) >= 1 {
		results["msg"] = msg[0]
	}
	c.JSON(200, results)
}

func Resp(c *gin.Context, results interface{}) {
	c.JSON(200, map[string]interface{}{
		"status": 0,
		"data":   results,
	})
}
