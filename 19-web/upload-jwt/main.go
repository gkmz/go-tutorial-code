package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Claims 是应用使用的 JWT 声明。
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken 使用环境变量中的密钥生成短时访问令牌。
func GenerateToken(userID int, secret []byte) (string, error) {
	if len(secret) < 32 {
		return "", errors.New("JWT_SECRET must contain at least 32 bytes")
	}
	claims := Claims{UserID: userID, RegisteredClaims: jwt.RegisteredClaims{Issuer: "go-tutorial", ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute))}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// ParseToken 校验令牌签名算法、签名和声明有效期。
func ParseToken(value string, secret []byte) (*Claims, error) {
	parsed, err := jwt.ParseWithClaims(value, &Claims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil || !parsed.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := parsed.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}

func randomName() (string, error) {
	var bytes [16]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes[:]), nil
}

func main() {
	secret := []byte(os.Getenv("JWT_SECRET"))
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.POST("/upload", func(c *gin.Context) {
		const maxSize = 8 << 20
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file"})
			return
		}
		if file.Size > maxSize {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file too large"})
			return
		}
		name, err := randomName()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		if err := os.MkdirAll("uploads", 0o750); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		dst := filepath.Join("uploads", name+strings.ToLower(filepath.Ext(filepath.Base(file.Filename))))
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"name": filepath.Base(dst)})
	})
	r.GET("/token", func(c *gin.Context) {
		token, err := GenerateToken(1, secret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.String(http.StatusOK, fmt.Sprint(token))
	})
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
