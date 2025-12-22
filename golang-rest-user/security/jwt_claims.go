package security

import (
	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	UserID     uint   `json:"user_id"`
	Username   string `json:"username"`
	TenantCode string `json:"tenant_code"`
	jwt.RegisteredClaims
}
