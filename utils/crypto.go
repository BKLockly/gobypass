package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"github.com/eknkc/basex"
)

/**
 * @Author Lockly
 * @Description
 * @Date 2023/9/24
 **/

var PwdKey = []byte("n8jXvNbusTQywJ9P")

func pd(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func ae(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	encryptBytes := pd(data, blockSize)
	crypted := make([]byte, len(encryptBytes))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

func E(data []byte) (string, error) {
	res, err := ae(data, PwdKey)
	if err != nil {
		return "", err
	}
	base85, _ := basex.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~")
	return base85.Encode(res), nil
}

func X(sBytes []byte) []byte {
	for i, _ := range sBytes {
		sBytes[i] = sBytes[i] ^ 24
	}
	return sBytes
}
