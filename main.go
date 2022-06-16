// main은 컴파일만을 위한 것
// 다른 이름으로 사용한다면 컴파일을 하지 않음(라이브러리로 됨)
package main

import (
	"fmt"
	"strings"

	"goTest/something"
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

/* naked return
 * return 할 variable을 명시 하지 않아도 된다.
 * return 타입에 return할 변수명을 적어야한다.
 * return을 명시하는게 좋지 않을까?
 */
func nakedLenAndUpper(name string) (nakedTotalLength int, nakedUpperName string) {
	// defer 함수가 return 후 실행합니다.
	// 아래 defer부터 역순으로 실행한다.
	defer fmt.Println("done1")
	defer fmt.Println("done2")
	defer fmt.Println("done3")
	defer fmt.Println("done4")
	nakedTotalLength = len(name)
	nakedUpperName = strings.ToUpper(name)

	return
}

// 연속으로된 param을 array 형태로 받는다.
func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	// array에 loop를 해줌
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	// index 생략
	total := 0
	for _, number := range numbers {
		total += number
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return total
}

func canIDrink(age int) bool {
	/* warning: should use return age >= 18
	 * 아마도 더 간결하게 쓰라고 경고주는 것 같다.
	 * 아무리 warning에 대한 걸 찾아봐도 안나와서 포기.
	 */
	// if age < 18 {
	// 	return false
	// }
	// return true

	// return age >= 18

	// 변수 koreanAge 생성, 세미콜론 오른쪽에서 변수를 만들고 있다.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
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
	fmt.Println("variants", name, otherName, otherOtherName)

	fmt.Println("func", multiply(2, 2), otherMultiply(2, 2))

	totalLength, upperName := lenAndUpper("yrkim")
	fmt.Println(totalLength, upperName)
	totalOtherLength, _ := lenAndUpper("aaa") // ignore second return value
	fmt.Println(totalOtherLength)

	repeatMe("a", "bb", "ccc", "dddd", "EEEEE")

	nakedTotalLength, nakedUpperName := nakedLenAndUpper("yrkim")
	fmt.Println("naked func", nakedTotalLength, nakedUpperName)

	result := superAdd(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("result", result)

	fmt.Println(canIDrink(16))
}
