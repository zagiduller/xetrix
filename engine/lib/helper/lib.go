package helper

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CurrentTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func ToFormatedJson(data interface{}) ([]byte, error) {
	format, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("ToFormatedJson: %s ", err)
	}
	return format, nil
}

func GetTokenStringFromReq(r *http.Request) string {
	var tokenString string
	headerFields := strings.Fields(r.Header.Get("Authorization"))
	if len(headerFields) < 2 {
		var coockieName string
		coockieName = "auth._token.local"

		c, err := r.Cookie(coockieName)
		if err != nil {
			log.Printf("ParIdFromReqMiddleware: %s", err)
		}

		if c != nil && len(c.Value) > 10 {
			tokenString = c.Value[9:] //Bearer%20
		}
	} else {
		tokenString = headerFields[1]
	}
	return tokenString
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
