package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"playcar/dbm"
	rdb "playcar/rd_m"
)

type LoginReq struct {
	UserId uint64 `json:"user_id"`
}

// 登录处理函数
func login(c *gin.Context) {
	req := &LoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser := &dbm.DbUser{}
	if err := dbm.JavaDb().Where("user_id=?", req.UserId).Find(&dbUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
	}
	// 模拟用户验证，实际情况下应与数据库或其他身份验证系统交互
	token, err := generateToken(dbUser.NickName, dbUser.PhoneNumber, req.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	playFlag := make([]*dbm.PlayFlag, 0)
	_ = dbm.GoDb().Where("user_id=?", req.UserId).Find(&playFlag).Error
	ret := make([]string, 0, len(playFlag))
	for _, v := range playFlag {
		ret = append(ret, v.RankType)
	}
	rdb.RankBoard.AddOnline(req.UserId, token)
	c.JSON(http.StatusOK, gin.H{"token": token, "play_list": ret})
}

func info(c *gin.Context) {
	val, _ := c.Get("user")
	user := val.(*JwtUser)
	c.JSON(http.StatusOK, user)
}
