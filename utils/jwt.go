package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
	"simple-project/keys"
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
	// creating header of the token
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}

	// marshaling header to json
	headerJson, err := json.Marshal(header)

	if err != nil {
		return "", err
	}

	// creating payload of the token
	expireTime := time.Now().Add(1 * time.Minute).Unix()
	payload := map[string]interface{}{
		"user_id":      userId,
		"username":     username,
		"expired_time": expireTime,
	}

	// marshaling payload
	payloadJson, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	// encoding header and payload to base64url
	encodedHeader := base64UrlEncode(headerJson)
	encodedPayload := base64UrlEncode(payloadJson)

	unsignedToken := encodedHeader + "." + encodedPayload

	// hashing header.payload
	h := hmac.New(sha256.New, []byte(keys.SecretKey))
	h.Write([]byte(unsignedToken))
	signature := h.Sum(nil)

	// encode signature
	encodedSignature := base64UrlEncode(signature)

	// returning token
	token := unsignedToken + "." + encodedSignature

	return token, nil

}

func VerifyJwt(token string) (bool, error) {
	tokenSlice := strings.Split(token, ".")

	encodedHeader := tokenSlice[0]
	encodedPayload := tokenSlice[1]
	tokenSignature := tokenSlice[2]

	payloadJson, err := base64UrlDecode(encodedPayload)

	if err != nil {
		return false, err
	}
	var payloadMap map[string]interface{}

	err = json.Unmarshal(payloadJson, &payloadMap)
	if err != nil {
		return false, err
	}

	// Checking token expiry
	var tokenExpireTime = int64(payloadMap["expired_time"].(float64))
	if tokenExpireTime <= time.Now().Unix() {
		return false, errors.New("token had expired")
	}

	// Checking Signature
	unsignedToken := encodedHeader + "." + encodedPayload

	h := hmac.New(sha256.New, []byte(keys.SecretKey))
	h.Write([]byte(unsignedToken))
	signature := h.Sum(nil)

	// encode signature
	encodedSignature := base64UrlEncode(signature)

	return encodedSignature == tokenSignature, nil
}
