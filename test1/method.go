package main

import (
	"fmt"
	"time"
	)

type Person struct {
	name string //名字
	sex  string   //性别, 字符类型
	age  int    //年龄
}

func (p Person) SetInfoValue() {
	//p.name =
	fmt.Printf("%v",p)
}

func (p *Person) SetInfoPointer(x string) {
	p.name = x
	fmt.Println("SetInfoPointer")
}

func Test(msg string){
	fmt.Println("This is "+msg)
}

func main(){
	go Test("routine 1")
	go Test("routine 2")
	fmt.Println("This is the main thread"+"GG")
	time.Sleep(1*time.Second)
}

/*
func main(){
	//p := new(Person)
	//p1 := &Person{}
	//p2 := &Person{"owen","male",20}
	//p3 := Person{"eponia","female",21}
	//.Printf("%p",p2)
	//p2.SetInfoValue()
	var x interface{}
	x = 123
	fmt.Println(x)
	x = "yyy"
	fmt.Print(x)

}*/

