package config

type DatabaseConfig struct {
	User    string `json:"user"`
	License string `json:"license"`
	Order   string `json:"order"`
	Spot    string `json:"spot"`
}

type MySqlConfig struct {
	Host string         `json:"host"`
	Port int32          `json:"port"`
	User string         `json:"user"`
	Pass string         `json:"pass"`
	DB   DatabaseConfig `json:"db"`
}

type ServerConfig struct {
	Name      string      `json:"name"`
	MySqlInfo MySqlConfig `json:"mysql"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int32  `mapstructure:"port"`
	NameSpace string `mapstructure:"namespace"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
