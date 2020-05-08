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
