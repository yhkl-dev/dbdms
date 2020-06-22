package midware

import (
	helper "dbdms/helpers"
	"dbdms/helpers/regex"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// JWTAuth Mid ware
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		if strings.Contains(path, "swagger") {
			return
		}

		token := context.Request.Header.Get("ACCESS_TOKEN")
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  -1,
				"message": "permission denied, Request has no token",
			})
			context.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ResolveToken(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  -1,
				"message": err.Error(),
			})
			context.Abort()
			return
		}
		method := context.Request.Method
		matched, _ := regex.IsRuleMatch(path)
		if matched {
			reg := regexp.MustCompile(`\/(\d+)`)
			path = reg.ReplaceAllString(path, `/:id`)
		}
		permPath := fmt.Sprintf("%v:%v", method, path)
		s := `
			SELECT
				count(p.code_name)
			FROM
				USER u,
				user_role_mapping ur,
				role r,
				role_permission_mapping rp,
				permission p 
			WHERE
				ur.user_id = u.id 
				AND ur.role_id = r.id 
				AND rp.role_id = r.id 
				AND rp.permission_id = p.id 
				AND u.user_name = "%v"
				AND p.code_name = "%v"
		`
		sql := fmt.Sprintf(s, claims.Name, permPath)
		var count int
		helper.SQL.Raw(sql).Row().Scan(&count)
		if count == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  -1,
				"message": "permission denied, No permission",
			})
			context.Abort()
			return
		}
		context.Set("claims", claims)
	}
}

// JWT struct
type JWT struct {
	SigningKey []byte
}

var (
	db *gorm.DB
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

func (j *JWT) ResolveToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}

		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func (j *JWT) RefreshTokne(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", nil
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}