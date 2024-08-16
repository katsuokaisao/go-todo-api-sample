package domain

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTGenerator interface {
	Generate(userUUID string, scopes []Scope) (string, error)
}

type jwtGenerator struct {
	secret string
}

func NewJWTGenerator(secret string) JWTGenerator {
	return &jwtGenerator{
		secret: secret,
	}
}

func (g *jwtGenerator) Generate(userUUID string, scopes []Scope) (string, error) {
	ss := make([]string, len(scopes))
	for i, scope := range scopes {
		ss[i] = string(scope)
	}
	scopesStr := ""
	if len(ss) > 0 {
		scopesStr = strings.Join(ss, ",")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userUUID,
		"scopes": scopesStr,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(g.secret))
}

type JWTParser interface {
	Parse(tokenString string) (*JWTClaims, error)
	HasScope(claims *JWTClaims, requiredScope Scope) bool
}

type jwtParser struct {
	secret string
}

func NewJWTParser(cfg *APIEnv) JWTParser {
	return &jwtParser{
		secret: cfg.JwtSecret,
	}
}

func (p *jwtParser) Parse(tokenString string) (*JWTClaims, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(p.secret), nil
	})
	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	userID, ok := (*claims)["userID"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}
	scopesStr, ok := (*claims)["scopes"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}
	scopesArr := strings.Split(scopesStr, ",")
	if !ValidateScopes(scopesArr) {
		return nil, ErrInvalidToken
	}
	scopes := make([]Scope, len(scopesArr))
	for i, scope := range scopesArr {
		scopes[i] = Scope(scope)
	}

	return &JWTClaims{
		UserID: userID,
		Scopes: scopes,
	}, nil
}

func (p *jwtParser) HasScope(claims *JWTClaims, requiredScope Scope) bool {
	for _, scope := range claims.Scopes {
		if scope.HasScope(requiredScope) {
			return true
		}
	}
	return false
}

type JWTClaims struct {
	UserID string  `json:"userID"`
	Scopes []Scope `json:"scopes"`
}

type Scope string

const (
	ScopeAdmin         Scope = "admin"
	ScopeTodoRead      Scope = "todo:r"
	ScopeTodoReadWrite Scope = "todo:rw"
)

func ToScope(s string) (Scope, bool) {
	switch s {
	case "admin":
		return ScopeAdmin, true
	case "todo:r":
		return ScopeTodoRead, true
	case "todo:rw":
		return ScopeTodoReadWrite, true
	default:
		return "", false
	}
}

func (s Scope) HasScope(requiredScope Scope) bool {
	if s == ScopeAdmin || s == requiredScope {
		return true
	}

	sa := strings.Split(string(s), ":")
	rsa := strings.Split(string(requiredScope), ":")
	if sa[0] == rsa[0] {
		// todo:rw * todo:r
		// todo:rw * todo:rw
		// todo:r * todo:r
		// todo:r * todo:rw
		if sa[1] == "rw" {
			return true
		}
		if sa[1] == "r" && rsa[1] == "r" {
			return true
		}
	}

	return false
}

func ValidateScopes(scopes []string) bool {
	for _, scope := range scopes {
		if !ValidateScope(scope) {
			return false
		}
	}
	return true
}

func ValidateScope(scope string) bool {
	switch Scope(scope) {
	case ScopeAdmin, ScopeTodoRead, ScopeTodoReadWrite:
		return true
	default:
		return false
	}
}
