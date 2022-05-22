package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	salt       = "niganiganiganiga"
	signingKey = "abobaaboba"
	tokenTLL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	SellerId uuid.UUID `json:"seller_id"`
}

type sellerService struct {
}

func NewSellerService() *sellerService {
	return &sellerService{}
}

func (us *sellerService) GenerateToken(sellerId uuid.UUID) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTLL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		sellerId,
	})

	return token.SignedString([]byte(signingKey))
}

func (us *sellerService) ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return uuid.Nil, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.SellerId, nil
}

func (us *sellerService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
