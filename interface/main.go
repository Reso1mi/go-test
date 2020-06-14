package main

import (
	"fmt"
	"time"
)

type Retiever interface {
	Get(url string) string
}

type MyRetiever struct {
	UserAgent string
}

type SBRetiever struct {
	Name string
}

func (r SBRetiever) Get(s string) string {
	return s
}

func (r MyRetiever) Get(s string) string {
	return s
}

func main() {
	var r Retiever
	r = &MyRetiever{"resolmi"}
	//fmt.Println(r.(type))
	switch v := r.(type) {
	case MyRetiever:
		fmt.Println(v.UserAgent)
	case SBRetiever:
		fmt.Println(v.Name)
	}
	//Type assertion
	fmt.Println(float64(3) / 2)

	a := time.Now()
	b := a.Add(time.Second * 10)
	fmt.Println(a.Unix(), b.Unix())
}
