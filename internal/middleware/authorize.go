package middleware

import (
	"context"
	"errors"
	"fmt"
	"go-training/internal/common"
	"go-training/internal/component/tokenprovider"
	"go-training/internal/modules/user/entity"
	"strings"

	userrepo "go-training/internal/modules/user/repository"

	"github.com/gin-gonic/gin"
)

type AuthenStore interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*entity.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extracTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequireAuth(tokenProvider tokenprovider.Provider, userRepo userrepo.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extracTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := userRepo.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
