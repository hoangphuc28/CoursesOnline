package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/pkg/common"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/pkg/utils"
	"strings"
)

func extractToken(token string) (string, error) {

	parts := strings.Split(token, " ")
	if parts[0] != "Bearer" || len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
		return "", utils.ErrInvalidToken
	}
	return parts[1], nil
}
func (m *MiddleareManager) RequireVerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := (c.Param("token"))
		payload, err := utils.ValidateJWT(token, m.cfg)
		if err != nil {
			panic(err)
		}
		c.Set("emailUser", payload.Email)
		c.Set("password", payload.Password)
		c.Set("key", payload.Key)
		c.Set("verified", payload.Verified)

		c.Next()

	}
}

func (m *MiddleareManager) RequiredAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := c.Request.Header.Get("Authorization")
		token, err := extractToken(s)
		if err != nil {
			panic(err)
		}

		payload, err := utils.ValidateJWT(token, m.cfg)
		if err != nil {
			panic(err)
		}

		//User-Service, err := m.userRepo.FindDataWithCondition(map[string]any{"email": payload.Email})
		//if err != nil {
		//	panic(err)
		//}

		if payload.Verified {
			panic(common.NewCustomError(errors.New("This account has not been verified!"), "This account has not been verified!"))
		}
		c.Set("userId", payload.Id)
		c.Set("emailUser", payload.Email)
		c.Next()
	}
}
