package auth

/*
token认证中间件
*/

import (
	"context"
	"errors"
	"fmt"

	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/baseerror"
	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/utils/middlewareutil"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"

	"github.com/dgrijalva/jwt-go"
)

var (
	// AccessTokenValidErr Token验证失败
	AccessTokenValidErr = baseerror.NewBaseError("AccessToken 验证失败")
	// AccessTokenValidationErrorExpiredErr  AccessToken过期
	AccessTokenValidationErrorExpiredErr = baseerror.NewBaseError("AccessToken过期")
	// AccessTokenValidationErrorMalformedErr  AccessToken格式错误
	AccessTokenValidationErrorMalformedErr = baseerror.NewBaseError("AccessToken格式错误")
)

// ValidJWTMiddleware 验证JWT TOKEN
func ValidJWTMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			token := fmt.Sprint(ctx.Value(middlewareutil.JWTConTextKey))
			if token == "" {
				err = errors.New("请登录")
				logger.Error("缺少token", zap.Error(err))
				return "", err
			}
			user, err := middlewareutil.ParseToken(token)
			if err != nil {
				if err, ok := err.(*jwt.ValidationError); ok {
					if err.Errors&jwt.ValidationErrorMalformed != 0 {
						return "", AccessTokenValidationErrorMalformedErr
					}
					if err.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
						return "", AccessTokenValidationErrorExpiredErr
					}
				}
				return "", AccessTokenValidErr
			}
			if user != nil {
				ctx = context.WithValue(ctx, "username", user.Username)
			}
			return next(ctx, request)
		}
	}
}
