package delay

import "time"

/**
 * @Author Lockly
 * @Description
 * @Date 2023/9/25
 **/

func Check() bool {
	s := time.Now()
	time.Sleep(10 * time.Second)
	e := time.Now()
	sleepTime := e.Sub(s)
	if sleepTime >= 10*time.Second {
		return true
	}
	return false
}
