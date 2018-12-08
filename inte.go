package main

import (
	"fmt"
	"reflect"
)

//定义接口
type Skills interface {
	Running()
	Getname() string
}

type Student struct {
	Name string
	Age  int
}

// 实现接口
func (p Student) Getname() string { //实现Getname方法
	fmt.Println(p.Name)
	return p.Name
}

func (p Student) Running() { // 实现 Running方法
	fmt.Printf("%s running", p.Name)
}
func main() {
	var stu1 Student
	stu1.Name = "wd"
	stu1.Age = 22
	stu1.Running() //调用接口

	var x interface{}

	s := "WD"
	x = s
	y, ok := x.(int)
	z, ok1 := x.(string)
	fmt.Println(y, ok)

	// 在switch中使用 变量名.(type) 查询变量是由哪个类型数据赋值。
	fmt.Println(reflect.TypeOf(z))
	fmt.Println(z, ok1)
}

//wd running
