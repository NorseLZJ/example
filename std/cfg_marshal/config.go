package cfg_marshal

type GetConfig struct {
	Code  []string `json:"code"`
	Proxy string   `json:"proxy"`
}

type FileSrc struct {
	ShareDir  string `json:"share_dir"`
	UpLoadDir string `json:"up_load_dir"`
	Port      string `json:"port"`
}

type Active struct {
	Min      int    `json:"Min"`
	Max      int    `json:"Max"`
	BaseAddr string `json:"BaseAddr"`
}

type MPing struct {
	Frequency int    `json:"Frequency"`
	Sleep     int    `json:"Sleep"`
	Size      string `json:"Size"`
	Addr      string `json:"Addr"`
}

type Reptile struct {
	// sql
	SqlConfig struct {
		Driver   string `json:"driver"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Db       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
		Charset  string `json:"charset"`
	} `json:"sqlConfig"`
	// url
	Url struct {
		SoldUrl string `json:"sold_url"`
		SellUrl string `json:"sell_url"`
	} `json:"url"`
	// userAgent
	UserAgent string `json:"user_agent"`
	// 区
	District []string `json:"district"`
}
