package generator

import (
	"crypto/rand"
	"fmt"
)

const accept = "1234567890"

func GenerateToken() (string, error) {
	buffer := make([]byte, 5)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(accept)
	token := ""
	for x := 0; x < 5; x++ {
		for i := 0; i < 4; i++ {
			buffer[i] = accept[int(buffer[i])%otpCharsLength]
		}
		token = token + fmt.Sprintf("%s ", string(buffer))
	}

	return token, nil
}
