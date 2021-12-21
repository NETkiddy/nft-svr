package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// DecodeWeAppUserInfo 解密微信小程序用户信息
func DecodeWeAppUserInfo(encryptedData string, sessionKey string, iv string) (string, error) {
	cipher, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		err = fmt.Errorf("encryptedData: %v decode error:%v", encryptedData, err.Error())
		return "", err
	}

	key, keyErr := base64.StdEncoding.DecodeString(sessionKey)
	if keyErr != nil {
		err = fmt.Errorf("sessionKey: %v decode error:%v", sessionKey, keyErr.Error())
		return "", keyErr
	}

	theIV, ivErr := base64.StdEncoding.DecodeString(iv)
	if ivErr != nil {
		err = fmt.Errorf("iv: %v decode error:%v", iv, ivErr.Error())
		return "", ivErr
	}

	result, resultErr := AESDecrypt(cipher, key, theIV)
	if resultErr != nil {
		return "", resultErr
	}
	return string(result), nil
}

// AESDecrypt AES解密
func AESDecrypt(cipherBytes, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, iv)
	dst := make([]byte, len(cipherBytes))
	blockModel.CryptBlocks(dst, cipherBytes)
	dst = PKCS7UnPadding(dst, block.BlockSize())
	return dst, nil
}

// PKCS7UnPadding pkcs7填充方式
func PKCS7UnPadding(dst []byte, blockSize int) []byte {
	length := len(dst)
	unpadding := int(dst[length-1])
	return dst[:(length - unpadding)]
}

func GetSignature(secret, method, endpoint, content, gmtDate, contentType string) string {
	if contentType == "" {
		contentType = "application/json"
	}

	contentMD5 := ""
	if len(content) != 0 {
		contentMD5 = base64.StdEncoding.EncodeToString([]byte(Md5(content)))
	}
	msg := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", method, endpoint, contentMD5, contentType, gmtDate)
	signature := base64.StdEncoding.EncodeToString(HmacSha1(secret, msg))

	return signature
}

func Md5(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}

func HmacSha1(key, data string) []byte {
	hmac := hmac.New(sha1.New, []byte(key))
	hmac.Write([]byte(data))
	return hmac.Sum([]byte(""))
}
