package config

type Config struct {
	From     string
	To       string
	Host     string
	Port     int64
	Password string
	Feeds    []string
}
