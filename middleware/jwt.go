package middleware

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"cospend/constant"
	"cospend/models"
	"cospend/pkg/util"

	"github.com/golang-jwt/jwt"
)

// GenerateToken is a JWT New With Claims SigningMethodHS256
func GenerateToken(user models.UserClaims) (string, error) {
	jwtExpired := util.Getenv("JWT_EXPIRED")
	uom := jwtExpired[len(jwtExpired)-1:]
	expiredToken, _ := strconv.Atoi(strings.Replace(jwtExpired, uom, "", -1))

	timeHourMinutes := time.Hour
	if uom == "d" {
		timeHourMinutes = time.Hour * 24
	}

	expiredAt := time.Now().Add(timeHourMinutes * time.Duration(expiredToken)).Unix()
	claims := models.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    "cospend",
			Subject:   "user",
		},
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	var signingKey = []byte(util.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JwtClaim is a JWT Parse With Claims Token
func JwtClaim(token string) (user *models.UserClaims, err error) {
	var signingKey = []byte(util.Getenv("SECRET_KEY"))
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("%s", constant.AUTHORIZATION)
	}

	id := claims["id"].(string)
	name := claims["name"].(string)
	email := claims["email"].(string)
	phoneNumber := claims["phoneNumber"].(string)
	user = &models.UserClaims{
		ID:          id,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
	return
}
