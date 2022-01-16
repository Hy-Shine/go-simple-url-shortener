package initiailze

type RedisConfig struct {
	DB int `json:"db"`
	Addr string `json:"ip"`
	Password string `json:"password"`
	Expire int
}