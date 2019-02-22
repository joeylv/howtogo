package main

import (
	"fmt"
	"reflect"
)

type rSkills interface {
	reading()
	running()
	setName()
}

type rStudent struct {
	Name string
	Age  int
}

func (self rStudent) setName(name string) {
	self.Name = name
	fmt.Printf("%s is running\n", self.Name)
}

func (self rStudent) runing() {
	fmt.Printf("%s is running\n", self.Name)
}
func (self rStudent) reading() {
	fmt.Printf("%s is reading\n", self.Name)
}
func main() {
	stu1 := rStudent{Name: "wd", Age: 22}
	stu1.setName("joey")
	inf := new(rSkills)
	stu_type := reflect.TypeOf(stu1)
	fmt.Println(reflect.TypeOf(stu1))
	inf_type := reflect.TypeOf(inf).Elem()                   // 特别说明，引用类型需要用Elem()获取指针所指的对象类型
	fmt.Println(stu_type.String())                           //main.Student
	fmt.Println(stu_type.Name())                             //Student
	fmt.Println(stu_type.PkgPath())                          //main
	fmt.Println(stu_type.Kind())                             //struct
	fmt.Println(stu_type.Size())                             //24
	fmt.Println(inf_type.NumMethod())                        //2
	fmt.Println(inf_type.Method(0), inf_type.Method(0).Name) // {reading main func() <invalid Value> 0} reading
	fmt.Println(inf_type.MethodByName("reading"))            //{reading main func() <invalid Value> 0} true

	fmt.Println("=======================")
	stu2 := reflect.ValueOf(&stu1)                    //获取Value类型
	stu2.Elem().FieldByName("Name").SetString("josh") //设置值

	stu2_type := reflect.TypeOf(stu2)
	fmt.Println(stu2_type.String())  //main.Student
	fmt.Println(stu2_type.Name())    //Student
	fmt.Println(stu2_type.PkgPath()) //main
	fmt.Println(stu2_type.Kind())    //struct
	fmt.Println(stu2_type.Size())
	//fmt.Println(str2.Elem(),age) //jack 11

}
