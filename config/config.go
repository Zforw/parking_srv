package config

type MySqlConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	DB   string `json:"db"`
}

type JWTConfig struct {
	SigningKey string `json:"key"`
	ExpireHour int64  `json:"hour"`
}

type ServerConfig struct {
	Name      string      `json:"name"`
	MySqlInfo MySqlConfig `json:"mysql"`
	JWTInfo   JWTConfig   `json:"jwt"`
}

type NacosConfig struct {
	Host        string `mapstructure:"host"`
	Port        int32  `mapstructure:"port"`
	NameSpace   string `mapstructure:"namespace"`
	DataId      string `mapstructure:"dataid"`
	Group       string `mapstructure:"group"`
	LogFilePath string `mapstructure:"log"`
}
