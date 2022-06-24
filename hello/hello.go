package main

import (
	"connor/greetings"
	"log"
)

func main() {
	// 给Log预定义一些属性，比如输出前缀和禁止打印时间、源文件、行号的标识
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// 调用 Hello 方法，返回 message, err
	message, err := greetings.Hello("")
	// 如果err不为nul，则表示有错误发生，将错误信息打印出来，并且退出程序
	if err != nil {
		log.Fatal(err)
	}
	// 打印 message 信息
	print(message)
}
