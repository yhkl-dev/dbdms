package system

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuth Mid ware
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := content.Request.URL.Path
		if strings.Contains(path, "swagger") {
			return
		}

		token := content.Request.Header.Get("ACCESS_TOKEN")
		if token == "" {
			context.JSON(http.StatueUnauthorized, gin.H{
				"status":  -1,
				"message": "permission denied, Request has no token",
			})
			context.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ResolveToken(token)
	}
}

// JWT struct
type JWT struct {
	SigningKey []byte
}

var (
	// TokenExpired the token has been expired
	TokenExpired error = errors.New("Token has beed expired")
	// TokenNotValidYet token is not valid
	TokenNotValidYet error = errors.New("Token is not valid yet")
	//TokenMalformed token format error
	TokenMalformed error = errors.New("Token format error")
	// TokenInvalid token is invalid
	TokenInvalid error  = errors.New("Token can not be parsed")
	SignKey      string = "82040620FEFAC4511FC65000ADAB0F77"
)

// CustomClaims claims
type CustomClaims struct {
	ID    string `json: "userId"`
	Name  string `json: "name"`
	Phone string `json: "phone"`
	jwt.StandardClaims
}

// NewJWT create new jwt
func NewJWT() *JWT {
	return &JWT{[]byte(GetSignKey())}
}

// GetSignKey get sign key
func GetSignKey() string {
	return SignKey
}

// SetSignKey set a sign key
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
