package main

import(
	"fmt"
	"regexp"
)

/*
func Match(pattern string, b []byte)(matched bool, err error)
func MatchString(pattern string, s string)(matched bool, err error)
以上两个函数区别在于匹配参数的类型，第一个为[]byte, 第二个为string

如果要得到匹配的具体内容，需用下面两个函数
func MustCompile(s string) *Regexp   返回值为一个正则表达式对象类(的指针) ,如果有err会直接抛出
func (re *Regexp) FindAllString(s string, n int) []string   返回n个匹配到的字符串，n为-1时返回所有匹配的字符串

捕获子表达式匹配内容
func (re *Regexp) FindAllStringSubmatch(s string, n int) []string 返回匹配到的结果和所有子表达式匹配结果
 */

func main(){

	// 传入参数需经强制转换成[]byte类型
	isok, _ := regexp.Match("[a-zA-Z]{3}", []byte("zH1"))
	fmt.Printf("%v\n",isok)

	// 直接传入字符串类型
	isok, _ = regexp.MatchString("[a-zA-Z]{3}","zHi")
	fmt.Printf("%v\n", isok)

	reg := regexp.MustCompile("[a-zA-Z]{3}")
	result := reg.FindAllString("zh1iliaoy", -1)
	fmt.Printf("%v\n", result)

	reg1 := regexp.MustCompile("\\d")            //如果正则表达式用双引号括起来，则里面的\符号需转义
	reg2 := regexp.MustCompile(`\d{3}`)			 //如果正则表达式用``括起来，则里面的表达式无需转义
	result = reg1.FindAllString("zh1iliaoy", -1)
	fmt.Printf("%v\n", result)
	result = reg2.FindAllString("zh222iliaoy", -1)
	fmt.Printf("%v\n", result)

	reg3 := regexp.MustCompile(`(.*?)(\d+)(.*?)\d(.*)\d`)
	matchs := reg3.FindAllStringSubmatch("tao123shi5han567",-1)
	fmt.Printf("%v\n", matchs) //会输出匹配到的整个字符串，以及括号中子表达式匹配到的字符串
}