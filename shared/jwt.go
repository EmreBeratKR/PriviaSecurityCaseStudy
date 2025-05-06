package shared

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(claims UserClaims, subject string) string {
	secret := getJWTSecret()
	claims.RegisteredClaims = jwt.RegisteredClaims{
		Subject: subject,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)

	if err != nil {
		return ""
	}

	return signedToken
}

func CalculateJWTExpireTime() time.Time {
	return time.Now().Add(24 * time.Hour)
}

func GetUserClaims(jwtToken string) *UserClaims {
	token, err := jwt.ParseWithClaims(jwtToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(getJWTSecret()), nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil
	}

	return claims
}

func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
