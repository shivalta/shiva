package generator

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GenerateToken() (string, error) {
	token := ""
	for x := 0; x < 5; x++ {
		temp := ""
		for i := 0; i < 4; i++ {
			temp = temp + fmt.Sprintf("%s", strconv.Itoa(rand.Intn(10)))
		}
		token = token + fmt.Sprintf("%s ", temp)
	}

	return token, nil
}
