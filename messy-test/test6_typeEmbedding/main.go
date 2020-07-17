package main

import "fmt"

/*
嵌入类型，或者嵌套类型，这是一种可以把已有的类型声明在新的类型里的一种方式，这种功能对代码复用非常重要。
在其他语言中，有继承可以做同样的事情，但是在Go语言中，没有继承的概念，
Go提倡的代码复用的方式是组合，所以这也是嵌入类型的意义所在，组合而不是继承，所以Go才会更灵活。
标准库中 ReadWriter就是一个很好的例子
type ReadWriter interface {
	Reader
	Writer
}
*/
type GoodBye interface {
	goodBye()
}

func sayGoodBye(gd GoodBye) {
	fmt.Printf("%T say goodBye\n", gd)
	gd.goodBye()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

func (u user) sayHello() {
	fmt.Println("Hello，i am a user")
}

//user实现了GoodBye接口
func (u user) goodBye() {
	fmt.Println("GoodBye，user implement GoodBye")
}

//覆盖embed type(内嵌类型) user的sayHello方法（覆盖可能不太准确，因为还是可以通过ad.user.sayHello调用）
//不管我们如何同名覆盖，都不会影响内部类型，我们还可以通过访问内部类型来访问它的方法、属性字段等。
func (a admin) sayHello() {
	fmt.Println("Hello，i am a admin")
}

func main() {
	ad := admin{user{"张三", "imglw@gmail.com"}, "管理员"}
	//如果admin也有name字段，那么这里打印的就是admin的name字段，会覆盖内嵌类型的name
	fmt.Println("可以直接调用,名字为：", ad.name)           //张三
	fmt.Println("也可以通过内部类型调用,名字为：", ad.user.name) //张三
	fmt.Println("但是新增加的属性只能直接调用，级别为：", ad.level)  //管理员
	//如果admin没有覆盖user的sayHello方法，这里调用的就是user的sayHello
	ad.sayHello()      //Hello，i am a admin
	ad.user.sayHello() //Hello，i am a user
	//如果内部类型实现了某个接口，那么外部类型也被认为实现了这个接口
	sayGoodBye(ad.user)
	sayGoodBye(ad)

}
