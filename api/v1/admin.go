package v1

import (
	"danci-api/global"
	"danci-api/middleware"
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func AdminLogin(context *gin.Context) {
	var loginParam request.Login
	_ = context.ShouldBind(&loginParam)
	if err, user := services.Login(loginParam); err != nil {
		response.FailWithMessage("用户名不存在或者密码错误", context)
	} else {
		tokenNext(context, user)
	}
}

func RegisterAdmin(context *gin.Context) {
	var adminParam request.AdminParam
	_ = context.ShouldBind(&adminParam)
	user, err := services.RegisterAdmin(model.Admin{
		UserName: adminParam.Username,
		Password: adminParam.Password,
		NickName: adminParam.Nickname,
	})
	if err != nil {
		response.FailWithMessage("管理员创建失败！", context)
	} else {
		tokenNext(context, user)
	}
}

func tokenNext(context *gin.Context, user *model.Admin) {
	j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		ID:         user.ID,
		NickName:   user.NickName,
		Username:   user.UserName,
		BufferTime: global.GVA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                              // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", context)
		return
	}
	response.OkWithDetailed(response.LoginResponse{
		User:      *user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", context)
	return
}
