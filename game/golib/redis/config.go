package redis

type RedConfig struct {
	Net             string `json:"net"`
	Addr            string `json:"addr"`
	Pass            string `json:"pass"`
	MaxIdle         int    `json:"max_idle"`
	MaxActive       int    `json:"max_active"`
	IdleTimeout     int    `json:"idle_timeout"`
	Wait            bool   `json:"wait"`
	MaxConnLifetime int    `json:"max_conn_lifetime"`
}
