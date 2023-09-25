package utils

import (
	"fmt"
	"io"
	"os"
)

/**
 * @Author Lockly
 * @Description
 * @Date 2023/9/24
 **/

func W(name, content string) {
	os.Mkdir("output", 0777)
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Print("Create folder failed")
	}

	defer f.Close()
	io.WriteString(f, content)
	f.Close()
}

func R(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return content
}
