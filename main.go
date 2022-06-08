// main은 컴파일만을 위한 것
// 다른 이름으로 사용한다면 컴파일을 하지 않음(라이브러리로 됨)
package main

import (
	"fmt"
	"goTest/something"
	"strings"
)

/* 각 param마다 타입을 지정해줘야한다. 지정하지 않으면 undefined.
 * return에 대한 타입도 지정해야한다.
 */
func otherMultiply(a, b int) int {
	return a * b
}
func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 연속으로된 param을 array 형태로 받는다.
func repeatMe(words ...string) {
	fmt.Println(words)
}

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

	fmt.Println(multiply(2, 2), otherMultiply(2, 2))

	totalLength, upperName := lenAndUpper("yrkim")
	fmt.Println(totalLength, upperName)
	totalOtherLength, _ := lenAndUpper("aaa") // ignore second return value
	fmt.Println(totalOtherLength)

	repeatMe("a", "bb", "ccc", "dddd", "EEEEE")
}
