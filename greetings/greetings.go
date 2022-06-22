package greetings

import "fmt"

func Hello(name string) string {
	// 将字符串赋值给变量 message
	message := fmt.Sprintf("Hi, %v. Welcome !", name)
	return message
}
