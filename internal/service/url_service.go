package service

import (
	"crypto/rand"
	"encoding/binary"
	"strings"

	"github.com/Rokli/URL-Shortener/internal/domain"
)

type UrlService struct {
}

func (u *UrlService) GenerateShortCode() string {

	var buf [8]byte
	rand.Read(buf[:])

	num := binary.BigEndian.Uint64(buf[:])

	return encodeBase62(num)
}

func encodeBase62(n uint64) string {
	var base62 = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	if n == 0 {
		return "0"
	}

	var builder strings.Builder
	for n > 0 {
		builder.WriteByte(base62[n%62])
		n /= 62
	}

	runes := []rune(builder.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func (u *UrlService) ValidateURL(url string) bool {
	return true
}

func (u *UrlService) CreateShortURL(originalURL string) (*domain.Url, error) {
	return domain.CreateNewUrl(), nil
}

func (u *UrlService) GetOriginalURL(shortCode string) (string, error) {
	return "google.com", nil
}
