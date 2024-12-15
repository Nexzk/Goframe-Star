package users

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"star/internal/dao"
	"star/internal/model/entity"
	"star/utility"
)

type UserClaim struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

func Login(ctx context.Context, username string, password string) (tokenString string, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("username", username).Scan(&user)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}
	if user.Id == 0 {
		return "", errors.New("用户不存在")
	}
	if user.Password != utility.EncryptPassword(password) {
		return "", errors.New("用户名或密码错误")
	}
	UserClaims := &UserClaim{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims)
	return token.SignedString(utility.JwtKey)
}

func Info(ctx context.Context) (user *entity.Users, err error) {
	user = new(entity.Users)
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
	tokenClaims, _ := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return utility.JwtKey, nil
	})
	if claims, ok := tokenClaims.Claims.(*UserClaim); ok && tokenClaims.Valid {
		dao.Users.Ctx(ctx).Where("id", claims.Id).Scan(&user)
		return user, nil
	} else {
		return nil, errors.New("token无效")
	}
}

func GetUid(ctx context.Context) (uint, error) {
	user, err := Info(ctx)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
