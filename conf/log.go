package conf

type log struct {
	Level     string `toml:"level"`
	Path      string `toml:"path"`
	Format    string `toml:"format"`
	MaxSize   int    `toml:"maxSize"`
	MaxBackup int    `toml:"maxBackup"`
	LocalTime bool   `toml:"localTime"`
	Compress  bool   `toml:"compress"`
	Console   bool   `toml:"console"`
}
