package greetings

import "fmt"
import "errors"

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name is empty!")
	}

	// 将字符串赋值给变量 message
	message := fmt.Sprintf("Hi, %v. Welcome !", name)
	return message, nil
}
