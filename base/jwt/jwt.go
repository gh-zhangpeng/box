package jwt

import (
	"box/base"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"time"
)

var key = []byte("box-key")

type boxClaims struct {
	UserID int64 `json:"userID"`
	jwt.StandardClaims
}

func ParseToken(token string) (*boxClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &boxClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if tokenClaims.Valid {
		if claims, ok := tokenClaims.Claims.(*boxClaims); ok {
			return claims, nil
		}
		log.Errorf("claims convert fail, err: %s, token: %s", err, token)
		return nil, base.ErrorSystemError
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, base.ErrorInvalidToken
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, base.ErrorExpiredToken
		} else {
			log.Errorf("parse token fail, err: %s, token: %s", err, token)
			return nil, base.ErrorSystemError
		}
	} else {
		log.Errorf("parse token fail, err: %s, token: %s", err, token)
		return nil, base.ErrorSystemError
	}
}

func GenerateToken(userID int64, expiresAt int64) (string, error) {
	claims := boxClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
			//Issuer:    strconv.f,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(key)
	if err != nil {
		log.Errorf("signed string fail, claims: %+v", claims)
		return "", base.ErrorSystemError
	}
	return signedString, nil
}
