package main

func test(arr []int) (int,bool){
	sum := 0
	for _,v := range(arr){
		sum += v
	}
	return sum, true
}

/*
type Student struct{
	Human
	score int
}
type Human struct{
	name string
	age int
}


func (x Human) Walk(steps int,place string)(int,bool){
	fmt.Printf("%s is %d years old\n",x.name,x.age)
	fmt.Printf("he walked %d steps in %s\n",steps, place)
	return steps,true
}


 */
//func main (){
	//arr := []int{1,2,3}
	//f := func(x,y int)int{ return x+y}
	//a := f(2,3)
	/*
	x:=[]int{1,2,3}
	fmt.Print(test(x))

	h1 := Human{"eponia",80}
	//s1 := Student{Human{"owen",25},100}
	a,b := h1.Walk(100,"park")
	fmt.Printf("%v%v",a,b)
	*/

//}
