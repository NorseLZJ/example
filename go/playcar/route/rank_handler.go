package route

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"playcar/dbm"
	rdb "playcar/rd_m"
	"time"
)

type UploadRankReq struct {
	RankType string  `json:"rank_type"`
	Road     string  `json:"road"` // 赛道 A,B,C,D
	Value    float64 `json:"value"`
	SubMoney float64 `json:"sub_money"`
}

type RankListReq struct {
	RankType string `json:"rank_type"`
	Road     string `json:"road"`  // 赛道
	Index    int    `json:"index"` // 0:day 1:week 2:month
}

type SelectRoadReq struct {
	Road string `json:"road"` // 赛道 A,B,C,D
}

const (
	RankDay   = 0
	RankWeek  = 1
	RankMonth = 2
)

var (
	RankIndexMap = map[int]struct{}{
		RankDay:   {},
		RankWeek:  {},
		RankMonth: {},
	}

	RankTypeMap = map[string]struct{}{
		"single":    {},
		"single_ai": {},
	}
	RankRoad = map[string]struct{}{
		"A": {},
		"B": {},
		"C": {},
		"D": {},
	}
)

func validate(rankType, road string, index int) bool {
	// 第三个没有这个参数就找个默认值填一下,其它的没有参数填空字符串
	if _, ok := RankRoad[road]; road != "" && !ok {
		return false
	}
	if _, ok := RankTypeMap[rankType]; rankType != "" && !ok {
		return false
	}
	if _, ok := RankIndexMap[index]; !ok {
		return false
	}
	return true
}

func uploadRank(c *gin.Context) {
	req := &UploadRankReq{}
	user := getJwtUser(c)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "data err"})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "data err"})
		return
	}
	if !validate(req.RankType, req.Road, 0) {
		c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("req err:%#v", req)})
		return
	}
	keys := GetRankKeys(fmt.Sprintf("rank:%s:%s:", req.RankType, req.Road))
	for _, key := range keys {
		rdb.RankBoard.AddScoreSmall(key, user.UserId, req.Value)
	}
	count := int64(0)
	if err := dbm.GoDb().Table("play_flag").Where("user_id=? and rank_type=?", user.UserId, req.RankType).Count(&count).Error; err != nil {
		log.Println(err)
	}
	if count == 0 {
		_ = dbm.GoDb().Create(&dbm.PlayFlag{UserId: int(user.UserId), RankType: req.RankType}).Error
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": "ok"})
}

func rankList(c *gin.Context) {
	req := &RankListReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "参数错误"})
		return
	}
	if !validate(req.RankType, req.Road, req.Index) {
		c.JSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("req err:%#v", req)})
		return
	}

	keys := GetRankKeys(fmt.Sprintf("rank:%s:%s:", req.RankType, req.Road))
	if len(keys) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "rank list err"})
		return
	}
	data := map[string]interface{}{}
	switch req.Index {
	case RankDay:
		data["day"] = rdb.RankBoard.GetTopN(keys[req.Index], 100)
	case RankWeek:
		data["week"] = rdb.RankBoard.GetTopN(keys[req.Index], 100)
	case RankMonth:
		data["month"] = rdb.RankBoard.GetTopN(keys[req.Index], 500)
	}
	c.JSON(http.StatusBadRequest, data)
}

func selectRoad(c *gin.Context) {
	req := &SelectRoadReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "参数错误"})
		return
	}
	if !validate("", req.Road, RankDay) {
		c.JSON(http.StatusBadRequest, gin.H{"err": "参数错误"})
		return
	}
	user := getJwtUser(c)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "data err"})
		return
	}
	value := fmt.Sprintf("%s_%d_%s_%d", req.Road, user.UserId, user.UserName, time.Now().UnixMilli())
	upKey := md5.Sum([]byte(value))
	value = fmt.Sprintf("%x", upKey)
	// TODO  扣积分
	c.JSON(http.StatusBadRequest, gin.H{"msg": "ok"})
}
