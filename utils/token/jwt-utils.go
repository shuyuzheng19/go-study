package token

import (
	"github.com/dgrijalva/jwt-go"
	"gorm-study/common"
	"gorm-study/entity"
	"time"
)

type myClaims struct {
	Username string
	Role     string
	jwt.StandardClaims
}

func CreateAccessToken(user entity.User) string {

	myClaims := myClaims{
		user.Username,
		"",
		jwt.StandardClaims{
			Id:        user.Username,
			Subject:   user.Username,
			Issuer:    user.NickName,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(common.AccessTokenExpireTime).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)

	result, err := token.SignedString([]byte(common.TokenSigningKey))

	if err != nil {
		panic(err)
	}

	return result

}

func CreateRefreshToken(user entity.User) string {

	myClaims := myClaims{
		user.Username,
		common.RoleAdmin,
		jwt.StandardClaims{
			Id:        user.Username,
			Issuer:    user.NickName + "---->" + user.NickName,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(common.RefreshTokenExpireTime).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)

	result, err := token.SignedString([]byte(common.TokenSigningKey))

	if err != nil {
		panic(err)
	}

	return result

}

func ParseTokenFormToUsername(token string) string {
	claims, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.TokenSigningKey), nil
	})
	if err != nil {
		panic(err)
	}

	valid := claims.Valid

	if valid {
		return claims.Claims.(*myClaims).Subject
	}

	return ""
}
