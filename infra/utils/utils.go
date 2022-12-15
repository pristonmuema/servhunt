package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"servhunt/infra/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationPayloadKey = "authorization_payload"
)

// CORSMiddleware it sets the CORS properties.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			APIResponse(ctx, "", http.StatusUnauthorized, false, err.Error())
			return
		}
		payload, err := tokenMaker.VerifyToken(authorizationHeader)
		if err != nil {
			APIResponse(ctx, "", http.StatusUnauthorized, false, err.Error())
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
