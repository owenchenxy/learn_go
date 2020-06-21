package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"strings"
)

func main(){
	url := "http://www.baidu.com"
	content_type := "application/x-www-form-urlencoded"
	//resp, err := http.Post(url, content_type, strings.NewReader("id=1,name=2"))

	// GET 请求
	resp, err := http.Get(url)

	// POST 请求
	//content_type := "application/x-www-form-urlencoded"
	//resp, err := http.Post(url, content_type, strings.NewReader("id=1,name=2"))

	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()               //主程序结束前关闭
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}