package main

import(
	"fmt"
	"github.com/astaxie/goredis"
)

/*
redis 基本操作：
	1. string
	2. hash   key对应的value可以是任意的对象类型
	3. set
	4. list
	5. zset
 */
type Person struct{
	color string
	height int
}

func main(){
	var client goredis.Client
	// set方法， set key-value name：eponia
	err := client.Set("name",[]byte("eponia"))
	name,err := client.Get("name")
	if err != nil{
		panic(err)
	}

	fmt.Print(string(name))

	// hmset 设置多个hash映射
	p := new(Person)
	p.color = "yellow"
	p.height = 160
	mapping := make(map[string]interface{})
	mapping["age"] = 10
	mapping["gender"] = "female"
	mapping["color"] = p.color
	mapping["height"] = p.height
	fmt.Printf("%v",*p)
	err = client.Hmset("student_info",mapping)
}