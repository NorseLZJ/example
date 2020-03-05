package cfg_marshal

type GetConfig struct {
	Code  []string `json:"code"`
	Proxy string   `json:"proxy"`
}

type FileSrc struct {
	ShareDir string `json:"share_dir"`
	Addr     string `json:"addr"`
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
