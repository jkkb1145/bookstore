package initjwt

import (
	"context"
	"demo02/global"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// jwt密钥
var jwtSecret = []byte("jkkb114514")

const (
	//一个用户登录网站时，给他的token的过期时间
	//过期需要重新登录
	AccessTokenExpire = 2 * time.Hour
	//刷新token过期时间
	RefreshTokenExpire = 7 * 24 * time.Hour
)

type Claims struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// GenerateTokenPair 生成访问token和刷新token
func GenerateTokenPair(userID uint, username string) (*TokenResponse, error) {
	// 生成访问token
	accessClaims := Claims{
		UserID:    userID,
		Username:  username,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// 生成刷新token
	refreshClaims := Claims{
		UserID:    userID,
		Username:  username,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// 将token存储到Redis
	if err := StoreTokenInRedis(userID, accessTokenString, refreshTokenString); err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    int64(AccessTokenExpire.Seconds()),
	}, nil
}

// StoreTokenInRedis 将token存储到Redis
func StoreTokenInRedis(userID uint, accessToken, refreshToken string) error {
	ctx := context.Background()
	userKey := fmt.Sprintf("user_tokens:%d", userID)

	// 使用hash存储用户的token信息
	err := global.RedisClient.HMSet(ctx, userKey,
		"access_token", accessToken,
		"refresh_token", refreshToken,
		"created_at", time.Now().Unix(),
	).Err()
	if err != nil {
		return err
	}

	// 设置过期时间为刷新token的过期时间
	return global.RedisClient.Expire(ctx, userKey, RefreshTokenExpire).Err()
}

// 检查token在redistribution中是否有效
func IsTokenValidInRedis(userID uint, token string, tokenType string) bool {
	ctx := context.Background()
	userKey := fmt.Sprintf("user_tokens:%d", userID)

	var redisToken string
	var err error

	if tokenType == "access" {
		redisToken, err = global.RedisClient.HGet(ctx, userKey, "access_token").Result()
	} else {
		redisToken, err = global.RedisClient.HGet(ctx, userKey, "refresh_token").Result()
	}

	if err != nil {
		return false
	}
	return redisToken == token
}
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 检查token是否在Redis中被撤销
		if !IsTokenValidInRedis(claims.UserID, tokenString, claims.TokenType) {
			return nil, errors.New("token已被撤销")
		}
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 删除Token实现退出登录
func RevokeToken(userID uint) error {
	ctx := context.Background()
	userKey := fmt.Sprintf("user_tokens:%d", userID)
	return global.RedisClient.Del(ctx, userKey).Err()
}
