package jwtauth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// JWTAuth Mid ware

// JWT struct
type JWT struct {
	SigningKey []byte
}

var (
	db *gorm.DB
	// TokenExpired the token has been expired
	errTokenExpired error = errors.New("Token has beed expired")
	// TokenNotValidYet token is not valid
	errTokenNotValidYet error = errors.New("Token is not valid yet")
	//TokenMalformed token format error
	errTokenMalformed error = errors.New("Token format error")
	// TokenInvalid token is invalid
	errTokenInvalid error  = errors.New("Token can not be parsed")
	signKey         string = "82040620FEFAC4511FC65000ADAB0F77"
)

// CustomClaims claims
type CustomClaims struct {
	UserID    string `json:"user_Id"`
	UserName  string `json:"user_name"`
	UserPhone string `json:"user_phone"`
	jwt.StandardClaims
}

// NewJWT create new jwt
func NewJWT() *JWT {
	return &JWT{[]byte(GetSignKey())}
}

// GetSignKey get sign key
func GetSignKey() string {
	return signKey
}

// SetSignKey set a sign key
func SetSignKey(key string) string {
	signKey = key
	return signKey
}

// CreateToken create jwt token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ResolveToken resolve token
func (j *JWT) ResolveToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errTokenNotValidYet
			} else {
				return nil, errTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errTokenInvalid
}

// RefreshToken refresh token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
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
	return "", errTokenInvalid
}
