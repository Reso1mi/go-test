package depend

import "fmt"

func init() {
	fmt.Println("1. 被依赖的包init函数会先执行")
}

func GetName() string {
	return "resolmi"
}
