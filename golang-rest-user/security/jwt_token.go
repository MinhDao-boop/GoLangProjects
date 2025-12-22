package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	jwtConfig *JWTConfig
}

func NewManager(jwtConfig *JWTConfig) *Manager {
	return &Manager{jwtConfig: jwtConfig}
}

func (m *Manager) GenerateAccessToken(userID uint, username, tenantCode string) (string, error) {
	claims := &AccessTokenClaims{
		UserID:     userID,
		Username:   username,
		TenantCode: tenantCode,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.jwtConfig.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.jwtConfig.AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.jwtConfig.SecrteKey)
}

func (m *Manager) ParseAccessToken(tokenStr string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return m.jwtConfig.SecrteKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
