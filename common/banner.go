package common

import "fmt"

/**
 * @Author Lockly
 * @Description
 * @Date 2023/9/25
 **/

var b = `
          _                          
         ( )                         
 __  ___ | |_  _ _  ___  ___  __  __ 
/o )( o )( o \( V )( o \( o )(_' (_' 
\__\ \_/ /___/ ) / / __//_^_\/__)/__)
 _|/          /_/  |_|               
`

func S() {
	fmt.Printf("%v\n", b)
	fmt.Printf("%18s -- author lockly\n\n", "")
}
