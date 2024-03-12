package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

var Encrypt *Encryption

// EncryptionAES 对称加密
type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

// PadPwd 填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// 去除填充
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}
	//去掉的长度
	unpadNum := int(dst[len(dst)-1])

	op := []byte("error")
	if len(dst) < unpadNum {
		return op, nil
	}
	str := dst[:len(dst)-unpadNum]
	return str, nil
}

// 加密
func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		return src
	}
	//密码填充
	NewSrcByte := PadPwd(srcByte, block.BlockSize())
	dst := make([]byte, len(NewSrcByte))
	//加密
	block.Encrypt(dst, NewSrcByte)
	//base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

// 解密
func (k *Encryption) AesDecoding(pwd string) string {
	//pwdByte := []byte(pwd)
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd
	}
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		return pwd
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	//去除填充
	dst, err = UnPadPwd(dst)
	if err != nil {
		return "0"
	}
	return string(dst)
}

func (k *Encryption) SetKey(key string) {
	k.key = key
}
