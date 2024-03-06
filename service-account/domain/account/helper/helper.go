package helper

import (
	"fmt"
	"service-account/domain/account/model"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateNoRekening(request model.CreateAccountRequest) (noRekening int64) {
	var time = time.Now()
	noRek := fmt.Sprintf("%v%v%v%v%v", time.Minute(), len(request.Nama), time.Day(), time.Second(), time.Hour())
	intNoRek, _ := strconv.Atoi(noRek)
	noRekening = int64(intNoRek)
	return
}

func StrToInt64(s string) (i int64) {
	str, _ := strconv.Atoi(s)
	return int64(str)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
