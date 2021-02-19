package middlewareauth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	controller "webcam/controllers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := controller.Token

		log.Println("user.Tokens", token)
		// var token Userdata

		if " " == token {
			log.Println("user.Tokens", token)

			c.Abort()

			c.Writer.Header().Set("Pragma", "no-cache")

			c.Redirect(301, "/")

		} else {
			atClaims := jwt.MapClaims{}

			fmt.Println("before username in token :", atClaims["user_id"], atClaims["user_name"])

			var jwkey = []byte(os.Getenv("JWT_Key"))

			tkn, err := jwt.ParseWithClaims(token, atClaims, func(token *jwt.Token) (interface{}, error) {

				return jwkey, nil
			})
			if err != nil {

				if err == jwt.ErrSignatureInvalid {

					c.Writer.WriteHeader(http.StatusUnauthorized)

					return
				}

				c.Writer.WriteHeader(http.StatusBadRequest)

				return
			}
			if !tkn.Valid {

				c.Writer.WriteHeader(http.StatusUnauthorized)

				return
			}
			fmt.Println("after username in token :", atClaims["user_id"], atClaims["user_name"])

			c.Set("username", atClaims["user_name"])

			c.Set("userid", atClaims["user_id"])

			c.Next()
		}
	}
}
