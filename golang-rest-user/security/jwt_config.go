package security

import (
	"os"
	"time"
)

type JWTConfig struct {
	SecrteKey      []byte
	AccessTokenTTL time.Duration
	Issuer         string
}

func LoadJWTConfig() *JWTConfig {
	return &JWTConfig{
		SecrteKey:      []byte(os.Getenv("JWT_SECRET_KEY")),
		AccessTokenTTL: time.Hour * 24,
		Issuer:         "golang-rest-user",
	}
}
