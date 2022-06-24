package greetings

import "fmt"
import "errors"
import "math/rand"
import "time"

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name is empty!")
	}

	// 将字符串赋值给变量 message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init 在初始化的时候就会被调用
func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("greetings initialize")
}

// randomFormat方法随机返回formats中的一条数据，类似于集合
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see u, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
