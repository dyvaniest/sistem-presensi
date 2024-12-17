package middlewares

import (
	"net/http"
	"sistem-presensi/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		sessionToken, err := ctx.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-Type") == "application/json" {
				ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
				ctx.Abort()
			}

			ctx.Redirect(http.StatusSeeOther, "/login")
			return
		}

		claims := &models.Claims{}
		tkn, err := jwt.ParseWithClaims(sessionToken, claims, func(token *jwt.Token) (interface{}, error) {
			return models.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Unauthorized"})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Bad Request"})
			ctx.Abort()
			return
		}

		if !tkn.Valid {
			ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("username", claims.Username)
		ctx.Next()
	})
}
