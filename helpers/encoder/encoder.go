package encoder

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func EncodeEmailVerify(str string) string {
	encode := base64.StdEncoding.EncodeToString([]byte(str))
	encodeEncrypt := base64.StdEncoding.EncodeToString([]byte(str + viper.GetString(`encrypt.additional`)))
	return fmt.Sprintf("u=%s&v=%s", encode, AesEncrypt(encodeEncrypt, viper.GetString(`encrypt.keystring`)))
}

func DecodeEmailVerify(email string, verify string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(email)
	if err != nil {
		return "", err
	}
	decrypt := AesDecrypt(verify, viper.GetString(`encrypt.keystring`))
	decodeDecrypt, err := base64.StdEncoding.DecodeString(decrypt)
	if err != nil {
		return "", err
	}
	decodeStr := string(decodeDecrypt)
	decodeStr = strings.ReplaceAll(decodeStr, viper.GetString(`encrypt.additional`), "")
	if decodeStr == string(decode) {
		return decodeStr, nil
	}
	return "", nil
}
