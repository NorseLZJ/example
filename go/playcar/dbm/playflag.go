package dbm

type PlayFlag struct {
	UserId   int    `gorm:"user_id" json:"user_id"`
	RankType string `gorm:"rank_type" json:"rank_type"`
}

func (*PlayFlag) TableName() string {
	return "play_flag"
}
