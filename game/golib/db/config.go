package db

type MySQLConfig struct {
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Net     string `json:"net"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	DbName  string `json:"db_name"`
	MaxConn int    `json:"max_conn"`
	MaxIdle int    `json:"max_idle"`
}

func (mc *MySQLConfig) MySQLDns(names []string) []string {

}
