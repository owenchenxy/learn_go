package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

func Add(x int, ch chan int){
	fmt.Println(x + x)
	ch <- 1
}

type Student struct{
	Name string
	age int
}

func main(){
	/* slice内存共享
	a := make([]int,3,5) //长度是3， 容量是5

	a1 := a              //将a1也声明为slice，此时a1和a指向相同的内存地址（注意只是相同的内存地址，变量属性仍然是自己的，比如len）
	fmt.Println("a1=a")
	fmt.Printf("%p  %v\n",a,a)
	fmt.Printf("%p  %v\n",a1,a1)


	a1 = append(a1,10)   //向a1中添加元素10，此时a1的len变为4（a的len仍为3），此时a1值为[0,0,0,10], a的值仍为[0,0,0]
	fmt.Println("a1 append 10")
	fmt.Printf("%p  %v\n",a,a)
	fmt.Printf("%p  %v\n",a1,a1)

	a1[0] = 5            //由于目前仍指向同一个内存地址，a的相应值也会改变，a=[5,0,0],a1=[5,0,0,10]
	fmt.Println("a1[0] = 5")
	fmt.Printf("%p  %v\n",a,a)
	fmt.Printf("%p  %v\n",a1,a1)

	a1 = append(a1,20)   //此时a1=[5,0,0,10,20], a=[5,0,0], len(a1)=5, 已经等于cap(a1),如果再添加元素则会面临重新开辟内存
	fmt.Println("a1 append 20")
	fmt.Printf("%p  %v\n",a,a)
	fmt.Printf("%p  %v\n",a1,a1)

	a1 = append(a1,30)   //此时len(a1)为6，已超出原来的容量，需要重新开辟内存，大小为当前len的两倍（12）,此时a和a1不再共享内存地址
	fmt.Println("a1 append 30")
	fmt.Printf("%p  %v\n",a,a)
	fmt.Printf("%p  %v\n",a1,a1)

	a1[0] = 100          //此时a1=[100,0,0,10,20,30],a=[5,0,0]
	fmt.Println("a1[0]=100")
	fmt.Printf("%p  %v\n",a,a)
	fmt.Printf("%p  %v\n",a1,a1)


	chs := make([]chan int, 10)
	for i:=0;i<10;i++ {
		chs[i] = make(chan int)
		go Add(i, chs[i])
	}

	for _,ch := range(chs){
		<- ch
	}

	//对数组类型进行json编码
	a := [5]int{1,2,3,4,5}
	res,err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n",string(res))
    //对map类型进行json编码
	m := make(map[string]string)
	m["owen"] = "chen"
	res1,err1 := json.Marshal(m)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("%v",string(res1))
	//对结构体进行json编码
	student := Student{"owen", 25}
	res2,err2 := json.Marshal(student)
	if err2 != nil{
		panic(err2)
	}
	fmt.Printf("%v\n",string(res2))
	*/
	m := md5.New()                          //创建一个md5对象
	m.Write([]byte("text to encrypt"))       //将需加密的内容写入该对象
	cipher := m.Sum(nil)                    //将m的checksum和与nil的相加，得到的即是m的checksum
	cipherStr := hex.EncodeToString(cipher)
	fmt.Print(cipherStr)
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

