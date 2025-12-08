package util

import (
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(9000) + 1000)
}

func HashOTP(otp string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(otp), 10)
	return string(b), err
}

func CheckOTP(hashedOTP, inputOTP string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedOTP), []byte(inputOTP)) == nil
}
