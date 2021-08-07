package config

const (
	APP_PORT   = ":8080"
	JWT_SECRET = "jwt-secret"
)

type Redis struct {
	Host     string
	Protocol string
}
