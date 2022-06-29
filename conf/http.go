package conf

type http struct {
	Host  string `toml:"host"`
	Port  int    `toml:"port"`
	Level string `toml:"level"`
}
