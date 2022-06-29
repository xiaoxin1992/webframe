package conf

import (
	"github.com/BurntSushi/toml"
)

type config struct {
	App   *app   `toml:"app"`
	Http  *http  `toml:"http"`
	MySQL *mysql `toml:"mysql"`
	Log   *log   `toml:"log"`
}

type app struct {
	Name string `toml:"name"`
}

var conf = &config{}

func NewConfig(filename string) (err error) {
	_, err = toml.DecodeFile(filename, &conf)
	return
}

func GetConfig() *config {
	return conf
}
