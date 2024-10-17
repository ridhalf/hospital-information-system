package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hospital-information-system/auth"
	"hospital-information-system/model/api"
	"hospital-information-system/model/web"
	"hospital-information-system/service"
	"net/http"
	"strings"
)

func AuthMiddleware(authJwt auth.Jwt, userService service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			authFailedMiddleware(ctx)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authJwt.ValidateToken(tokenString)
		if err != nil {
			authFailedMiddleware(ctx)
			return
		}
		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			authFailedMiddleware(ctx)
		}
		userID := int(payload["user_id"].(float64))
		request := web.UserFindByIdRequest{Id: userID}
		user, err := userService.FindById(request)
		if err != nil {
			authFailedMiddleware(ctx)
			return
		}
		ctx.Set("user", user)
	}
}
func authFailedMiddleware(ctx *gin.Context) {
	response := api.APIResponse("access Denied: Please log in to continue.", 401, "Unauthorized", nil)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
