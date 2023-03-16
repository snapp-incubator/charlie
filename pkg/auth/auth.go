package auth

import (
	"errors"
	"github.com/amirhnajafiz/DJaaS/pkg/enum"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	errSigningMethod = errors.New("error in signing method")
	errInvalidToken  = errors.New("token is invalid")
)

type Auth struct {
	key    string
	expire int
}

// New builds a new auth struct for handling JWT tokens.
func New(cfg Config) *Auth {
	return &Auth{
		key:    cfg.PrivateKey,
		expire: cfg.ExpireTime,
	}
}

// GenerateJWT creates a new JWT token.
func (a *Auth) GenerateJWT(c Claims) (string, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// create claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = c.Username
	claims["role"] = c.Role
	claims["state"] = c.State
	claims["expire"] = time.Now().Add(time.Duration(a.expire) * time.Minute)

	// generate token string
	tokenString, err := token.SignedString([]byte(a.key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT gets a token string and extracts the data.
func (a *Auth) ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errSigningMethod
		}

		return []byte(a.key), nil
	})
	if err != nil {
		return nil, err
	}

	// taking out claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c := new(Claims)

		c.Username = claims["username"].(string)
		c.Role = claims["role"].(enum.Role)
		c.State = claims["state"].(enum.State)

		return c, nil
	}

	return nil, errInvalidToken
}
