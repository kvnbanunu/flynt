package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int    `json:"user_id"`
	Timezone string `json:"timezone"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func NewToken(id int, timezone string) (string, error) {
	role := "user"
	idStr := strconv.Itoa(id)
	if idStr == os.Getenv("ADMIN_ID") {
		role = "admin"
	}

	claims := Claims{
		id,
		timezone,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(CFG.Secret))
}
