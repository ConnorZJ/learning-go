package main

import (
	"connor/greetings"
)

func main() {
	// 调用 Hello 方法，返回 message
	message := greetings.Hello("Connor")
	// 打印 message 信息
	print(message)
}
