package dbg

import (
	"strings"
	"syscall"
	"unsafe"
)

/**
 * @Author Lockly
 * @Description
 * @Date 2023/9/25
 **/

var blacklistDBG = []string{
	"IDA",
	"OLLY",
	"WINDBG",
	"GHIDRA",
}

const MAX_PATH = 260

func Check() bool {
	handle, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return false
	}

	pe32 := syscall.ProcessEntry32{}
	pe32.Size = uint32(unsafe.Sizeof(pe32))

	err = syscall.Process32First(handle, &pe32)

	for err == nil {
		exeFile := strings.ToUpper(syscall.UTF16ToString(pe32.ExeFile[:MAX_PATH]))
		for _, pn := range blacklistDBG {
			if strings.Contains(exeFile, pn) {
				return true
			}
		}
		err = syscall.Process32Next(handle, &pe32)
	}

	if ret, _, _ := syscall.NewLazyDLL(string([]byte{
		'k', 'e', 'r', 'n', 'e', 'l', '3', '2',
	})).NewProc(string([]byte{
		'I', 's', 'D', 'e', 'b', 'u', 'g', 'g', 'e', 'r', 'P', 'r', 'e', 's', 'e', 'n', 't',
	})).Call(); ret != 0 {
		return true
	}

	return false
}
