package conf

type mysql struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}
