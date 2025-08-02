package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"simple-project/keys"
	"strings"
)

func base64UrlEncode(data []byte) string {
	enc := base64.URLEncoding.WithPadding(base64.NoPadding)
	return enc.EncodeToString(data)
}

func base64UrlDecode(str string) ([]byte, error) {
	enc := base64.URLEncoding.WithPadding(base64.NoPadding)
	return enc.DecodeString(str)
}

func CreateJwt(userId int, username string) (string, error) {
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}

	headerJson, err := json.Marshal(header)

	if err != nil {
		return "", err
	}

	payload := map[string]interface{}{
		"user_id":  userId,
		"username": username,
	}

	payloadJson, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	secretKey := keys.SecretKey

	encodedHeader := base64UrlEncode(headerJson)
	encodedPayload := base64UrlEncode(payloadJson)

	unsignedToken := encodedHeader + "." + encodedPayload

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(unsignedToken))

	signature := h.Sum(nil)
	encodedSignature := base64UrlEncode(signature)

	token := unsignedToken + "." + encodedSignature

	return token, nil

}
