package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	mathRand "math/rand"
	"time"

	"github.com/ttacon/libphonenumber"
)

var allLeters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var errorInvalidPhoneNumber = errors.New("Invalid phone number. Ensure number is in international format (23327xxxxxxx)")

func init() {
	mathRand.Seed(time.Now().UnixNano())
}

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = allLeters[mathRand.Intn(len(allLeters))]
	}
	return string(b)
}

func ParsePhoneNumber(phoneNumber string) (string, error) {
	parsedPhoneNumber, err := libphonenumber.Parse(phoneNumber, "GH")

	if err != nil {
		return "", err
	}

	isOk := libphonenumber.IsValidNumberForRegion(parsedPhoneNumber, "GH")
	if !isOk {
		return "", errorInvalidPhoneNumber
	}

	str := fmt.Sprintf("%s", libphonenumber.Format(parsedPhoneNumber, libphonenumber.E164))
	return str, nil
}

func GenerateSecureKey() string {
	k := make([]byte, 32)
	io.ReadFull(rand.Reader, k)
	return fmt.Sprintf("%x", k)
}
