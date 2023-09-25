package resources

import (
	"github.com/klauspost/cpuid"
	"runtime"
	"syscall"
	"unsafe"
)

/**
 * @Author Lockly
 * @Description
 * @Date 2023/9/25
 **/

type memoryStatusEx struct {
	dwLength                uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func Check() bool {
	if cpuid.CPU.VM() {
		return true
	}

	memStatus := memoryStatusEx{}
	memStatus.dwLength = (uint32)(unsafe.Sizeof(memStatus))

	if ret, _, _ := syscall.NewLazyDLL(string([]byte{
		'k', 'e', 'r', 'n', 'e', 'l', '3', '2',
	})).NewProc(string([]byte{
		'G', 'l', 'o', 'b', 'a', 'l', 'M', 'e', 'm', 'o', 'r', 'y', 'S', 't', 'a', 't', 'u', 's', 'E', 'x',
	})).Call((uintptr)(unsafe.Pointer(&memStatus))); ret == 0 {
		return false
	}

	if runtime.NumCPU() < 2 || memStatus.ullTotalPhys < 1<<31 {
		return true
	}

	return false
}
