package route

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"net/http"
	rdb "playcar/rd_m"
	"time"
)

type route struct {
	SecretKey string
	*gin.Engine
}

var (
	GRoute *route
)

type JwtUser struct {
	UserName, PhoneNumber string
	UserId                uint64
}

func Init(secretKey, host string) {
	GRoute = &route{
		SecretKey: secretKey,
		Engine:    gin.Default(),
	}

	// 设置路由
	init := GRoute.Group("")
	{
		// 不需要登陆的
		init.POST("/login", login)
	}

	api := GRoute.Group("/api").Use(Secured)
	{
		api.GET("info", info)
		api.POST("upload_rank", uploadRank)
		api.GET("rank_list", rankList)
		api.POST("select_road", selectRoad)
	}

	// 启动服务器
	_ = GRoute.Run(host)
}

func Secured(c *gin.Context) {
	abort := func(errMsg string) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
		c.Abort()
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		abort("Missing Authorization header")
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(GRoute.SecretKey), nil
	})

	if err != nil || !token.Valid {
		abort("Invalid or expired token")
		return
	}

	exp := int64(token.Claims.(jwt.MapClaims)["exp"].(float64))
	if time.Now().Unix() < exp {
		abort("Invalid or expired token")
		return
	}

	// 在这里处理需要身份验证的逻辑，例如返回受保护资源的数据
	username := token.Claims.(jwt.MapClaims)["username"].(string)
	phone := token.Claims.(jwt.MapClaims)["phone"].(string)
	userId := uint64(token.Claims.(jwt.MapClaims)["uid"].(float64))
	if !rdb.RankBoard.IsOnline(userId, tokenString) {
		abort("Invalid or expired token")
		return
	}

	c.Set("user", &JwtUser{
		UserName:    username,
		UserId:      userId,
		PhoneNumber: phone,
	})
}

func generateToken(username, phone string, uid uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"phone":    phone,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"uid":      uid,
	})
	tokenString, err := token.SignedString([]byte(GRoute.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getJwtUser(c *gin.Context) *JwtUser {
	uu, exist := c.Get("user")
	if !exist {
		return nil
	}
	user, ok := uu.(*JwtUser)
	if !ok {
		return nil
	}
	return user
}
