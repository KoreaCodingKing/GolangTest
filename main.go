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

	// 변수/상수 이름 type = value
	const name string = "yrkim"

	var otherName string = "kim"
	otherOtherName := "kim"  // 이 문법은 func 내부에서만 가능, 변수에서만 적용 가능
	otherOtherName = "false" // value값이 다른 타입 false등 다른 타입이면 오류. 위에서 string 으로 선언되어있음.
	otherName = "aaa"
	fmt.Println(name, otherName, otherOtherName)
}
