package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yon-module/yon-framework/server/response"
)

func Signature() gin.HandlerFunc {
	clientIDEnv := os.Getenv("CLIENT_ID")
	clientSecretEnv := os.Getenv("CLIENT_SECRET")
	return func(c *gin.Context) {
		signature := c.GetHeader("X-Signature")
		if signature == "" {
			response.ErrorResponse(response.Unauthorized, "Missing X-Signature header", nil).Json(c)
			c.Abort()
			return
		}

		timeStamp := c.GetHeader("X-Timestamp")
		if timeStamp == "" {
			response.ErrorResponse(response.Unauthorized, "Missing X-Timestamp header", nil).Json(c)
			c.Abort()
			return
		}

		clientId := c.GetHeader("X-Client-Id")
		if clientId == "" {
			response.ErrorResponse(response.Unauthorized, "Missing X-Client-Id header", nil).Json(c)
			c.Abort()
			return
		}

		if clientId != clientIDEnv {
			response.ErrorResponse(response.Unauthorized, "Client ID is not registered", nil).Json(c)
			c.Abort()
			return
		}

		bodyEncrypted := ""
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			b, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				response.ErrorResponse(response.Unauthorized, "Missing request body", nil).Json(c)
				c.Abort()
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewReader(b))

			bodyEncrypted = minifySha256Body(b)
		}

		auth := c.GetHeader("Authorization")
		token := ""
		if auth != "" {
			parts := strings.Fields(auth)
			if len(parts) == 2 {
				token = parts[1]
			} else {
				token = auth
			}
		}

		signData := strings.Join([]string{
			c.Request.Method,
			c.Request.URL.Path,
			token,
			bodyEncrypted,
			timeStamp,
		}, ":")

		mac := hmac.New(sha256.New, []byte(clientSecretEnv))
		mac.Write([]byte(signData))
		expectedSig := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(expectedSig), []byte(signature)) {
			response.ErrorResponse(response.Unauthorized, "Invalid signature", nil).Json(c)
			c.Abort()
			return
		}

		c.Next()
	}
}

func minifySha256Body(bodyBytes []byte) string {
	minified := bodyBytes
	trimmed := bytes.TrimSpace(bodyBytes)
	if len(trimmed) > 0 {
		var js interface{}
		if json.Unmarshal(trimmed, &js) == nil {
			if b, err := json.Marshal(js); err == nil {
				minified = b
			} else {
				minified = trimmed
			}
		} else {
			minified = trimmed
		}
	}
	sum := sha256.Sum256(minified)
	return hex.EncodeToString(sum[:])
}
