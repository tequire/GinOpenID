package auth

import (
	"errors"
	"fmt"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/tequire/Competentia/pkg/services/tokens"
	"strings"
)

func authorize(c *gin.Context) (*oidc.IDToken, error) {
	header := c.GetHeader("Authorization")
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return nil, errors.New("invalid authorization header")
	}
	return tokens.Verifier().Verify(c, parts[1])
}

// IsAuthorized checks wether a user is authorized.
func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := authorize(c)
		if err != nil {
			fmt.Println(err.Error())
			c.Abort()
			c.JSON(403, gin.H{})
		}
	}
}

// IsAuthorizedWithScopes yolo
func IsAuthorizedWithScopes(scopes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := authorize(c)
		if err != nil {
			c.Abort()
			c.JSON(403, gin.H{})
		}
		fmt.Println(token)
	}
}
