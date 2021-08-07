package config

const (
	APP_PORT = ":8080"
)

type Redis struct {
	Host     string
	Protocol string
}