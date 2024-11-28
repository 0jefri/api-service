package config

// Server Config
type Server struct {
	Port int `mapstructure:"port"`
}

// Database Config
type Database struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
	Driver   string `mapstructure:"driver"`
}

// Configuration
type Config struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
}
