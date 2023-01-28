package tokens_services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/A-Siam/bracker/auth/src/model"
)

type Claims struct {
	Groups []string `json:"groups"`
	jwt.RegisteredClaims
}

func ConvertUserToClaims(user model.User, expirationInMins int64) Claims {
	groups := make([]string, 0)
	for _, group := range user.Groups {
		groups = append(groups, group.Name)
	}
	return Claims{
		Groups: groups,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth",
			Subject:   user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationInMins) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

func GetAccessToken(claims Claims) (string, error) {
	key, exist := os.LookupEnv("ACCESS_TOKEN_KEY")
	if !exist {
		return "", errors.New("token key is not set")
	}
	return getToken(claims, key)
}

func GetRefreshToken(claims Claims) (string, error) {
	key, exist := os.LookupEnv("REFRESH_TOKEN_KEY")
	if !exist {
		return "", errors.New("token key is not set")
	}
	return getToken(claims, key)
}

func getToken(claims Claims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
