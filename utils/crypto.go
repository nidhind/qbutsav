package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// Generate access token
func GenerateAccessToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// Generate SHA1 hash
func GenerateHash(v string) []byte {
	hash := sha1.Sum([]byte(v))
	return hash[:]
}

// Generate simple JWT token
func GenerateJWTToken(claims map[string]interface{}) string {
	c := make(jwt.MapClaims)
	for k, v := range claims {
		c[k] = v
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// TODO: validate 'JWT_SECRET' on server startup
	s, _ := t.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return s
}

// Parse simple JWT token
func ParseJWTToken(t string) (map[string]interface{}, string) {
	j, err := jwt.Parse(t, keyFunc())
	if err == nil && j.Valid {
		// Token is valid
		jwtClaims, _ := j.Claims.(jwt.MapClaims)
		claims := make(map[string]interface{})
		for k, v := range jwtClaims {
			claims[k] = v
		}
		return claims, ""
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			// The token is not JWT or is malformed
			return map[string]interface{}{}, "MALFORMED_TOKEN"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			// Yet claims are accessable
			jwtClaims, _ := j.Claims.(jwt.MapClaims)
			claims := make(map[string]interface{})
			for k, v := range jwtClaims {
				claims[k] = v
			}
			return claims, "EXPIRED_TOKEN"
		} else {
			return map[string]interface{}{}, err.Error()
		}
	} else {
		return map[string]interface{}{}, err.Error()
	}
}

// Retrievs secret key for jwt.Parse function
func keyFunc() jwt.Keyfunc {
	return func(_ *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}
}
