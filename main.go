// main은 컴파일만을 위한 것
// 다른 이름으로 사용한다면 컴파일을 하지 않음(라이브러리로 됨)
package main

import (
	"fmt"
	"goTest/something"
)

/* 만약 function을 다른 모듈로 export하고 싶가면 Println 처럼 대문자 사용한다.
 */
func main() {
	fmt.Println("Helloword")
	something.SayHello()

	// Not working
	// sayBye not exported by package something
	// something.sayBye()
}
