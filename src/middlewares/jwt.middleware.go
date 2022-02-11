package middlewares

import (
	"log"
	"net/http"
	"restapi/src/helpers"
	"restapi/src/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService services.JWTServivce) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := helpers.BuildErrorResponse(helpers.ResponseError{
				Message: "No permission",
				Errors:  "failed to process request",
			})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		// neu token hop le => khong lam gi
		if token.Valid {
			// co the check them cac thuoc tinh Claims o day nua
			claims := token.Claims.(jwt.MapClaims)
			if claims["uid"] == "1" {
				log.Println("------", "Chao chu nhan")
			} else {
				log.Println("Claim[UserID]: ", claims["uid"])
				log.Println("Claim[Issuer] :", claims["iss"])
				log.Println("Claim[Subject] :", claims["sub"])
				log.Println("Claim[Audience] :", claims["aud"])
				log.Println("Claim[ExpiresAt] :", claims["exp"])
				log.Println("Claim[IssuedAt] :", claims["iat"])
				log.Println("Claim[Id] :", claims["jti"])
			}
		} else {
			response := helpers.BuildErrorResponse(helpers.ResponseError{
				Message: "token invalid or expired",
				Errors:  err.Error(),
			})
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
	}
}
