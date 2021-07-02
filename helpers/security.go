package helpers

import (
	"crypto/rand"
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"io"
	"math/big"
	"regexp"
)

func init() {
	assertAvailablePRNG()
}

func RetrieveSymbol(value bool) string {
	if value {
		return ":white_check_mark:"
	} else {
		return ":x:"
	}
}

func RetrievePasswordPolicy(passwordLength bool, passwordNum bool, passwordAscii bool, passwordMajAscii bool, passwordSymbol bool) error {
	return emoji.Errorf(
		"password len should be >= 12 characters: %s\n"+
			"password require atleast one number: %s\n"+
			"password require atleast one lower ascii character [a-z]: %s\n"+
			"password require atleast one upper ascii character [A-Z]: %s\n"+
			"password require atleast one symbol: %s", RetrieveSymbol(passwordLength), RetrieveSymbol(passwordNum), RetrieveSymbol(passwordAscii), RetrieveSymbol(passwordMajAscii), RetrieveSymbol(passwordSymbol))
}

func CheckPasswordLever(ps string) error {
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	var regexFunctor = func(regex string, needle string) bool {
		if b, err := regexp.MatchString(regex, needle); !b || err != nil {
			return false
		}
		return true
	}
	hasError := false
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		hasError = true
	} else if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		hasError = true
	} else if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		hasError = true
	} else if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		hasError = true
	} else if len(ps) < 12 {
		hasError = true
	}
	if hasError {
		return RetrievePasswordPolicy(len(ps) > 12, regexFunctor(num, ps), regexFunctor(a_z, ps), regexFunctor(A_Z, ps), regexFunctor(symbol, ps))
	}
	return nil
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-!@#$%*;.?,"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
