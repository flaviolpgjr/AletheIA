package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func GeneratePromiseHash(text string) string {
	normalized := strings.ToLower(strings.TrimSpace(text))

	sum := sha256.Sum256([]byte(normalized))

	return hex.EncodeToString(sum[:])
}