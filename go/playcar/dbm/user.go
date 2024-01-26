package dbm

//用户基础信息

type DbUser struct {
	UserId       int    `gorm:"user_id" json:"user_id"`             // 用户ID从10001自增长
	PhoneNumber  string `gorm:"phone_number" json:"phone_number"`   // 手机号
	Mailbox      string `gorm:"mailbox" json:"mailbox"`             // 邮箱
	HeadPrtraits string `gorm:"head_prtraits" json:"head_prtraits"` // 头像
	NickName     string `gorm:"nick_name" json:"nick_name"`         // 昵称
	//Passwd          string  `gorm:"passwd" json:"passwd"`                   // 密码
	//TradePassWord   string  `gorm:"trade_pass_word" json:"trade_pass_word"` // 操作密码
	//Autograph       string  `gorm:"autograph" json:"autograph"`             // 个性签名
	//Balance         float64 `gorm:"balance" json:"balance"`                 // 我的资产
	//Alipay          string  `gorm:"alipay" json:"alipay"`                   // 支付宝账号
	//Alipayname      string  `gorm:"alipayname" json:"alipayname"`           // 支付宝姓名
	//Realnametype    int     `gorm:"realnametype" json:"realnametype"`       // 0.未实名1.待审核2.已通过
	//CreateTime      string  `gorm:"create_time" json:"create_time"`         // 账号创建时间
	//StatusId        string  `gorm:"status_id" json:"status_id"`             // 状态：0启用，1禁用
	//Address         string  `gorm:"address" json:"address"`                 // 以太坊地址
	//PrivateKey      string  `gorm:"privateKey" json:"privateKey"`           // 以太坊私钥
	//Whitelist       int     `gorm:"whitelist" json:"whitelist"`             // 白名单0.未开启1.已开启
	//Realname        string  `gorm:"realname" json:"realname"`
	//Realno          string  `gorm:"realno" json:"realno"`
	//InvitationId    int     `gorm:"invitation_id" json:"invitation_id"`
	//Invitationcount int     `gorm:"invitationcount" json:"invitationcount"`
	//Szcount         int     `gorm:"szcount" json:"szcount"`
	//Sztime          string  `gorm:"sztime" json:"sztime"`
	//Iscreater       int     `gorm:"iscreater" json:"iscreater"` // 0不是创作者1是创作者
	//Money           float64 `gorm:"money" json:"money"`         // 元气币
	//Email           string  `gorm:"email" json:"email"`
	//InvitationCode  string  `gorm:"invitation_code" json:"invitation_code"`
	//Lotterytimes    int     `gorm:"lotterytimes" json:"lotterytimes"` // 抽奖次数
	//Buycard         int     `gorm:"buycard" json:"buycard"`           // 购物券数量
	//Ustype          int     `gorm:"ustype" json:"ustype"`             // 0.手机号注册1.邮箱注册
	//Tudicard        int     `gorm:"tudicard" json:"tudicard"`         // 元宇宙券数量
	//UpdateTime      string  `gorm:"update_time" json:"update_time"`
	//MapShops        string  `gorm:"map_shops" json:"map_shops"` // 地图商铺数量jsonarray
}

func (*DbUser) TableName() string {
	return "users"
}
