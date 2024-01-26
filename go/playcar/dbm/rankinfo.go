package dbm

type RankInfo struct {
	Id         int     `gorm:"id" json:"id"`
	UserId     int     `gorm:"user_id" json:"user_id"`
	Score      float64 `gorm:"score" json:"score"`
	RankType   string  `gorm:"rank_type" json:"rank_type"`
	CreateTime string  `gorm:"create_time" json:"create_time"`
}

func (*RankInfo) TableName() string {
	return "rank_info"
}
