package main

import "net/http"  //go语言内建的http包

//开启一个http服务端的方法：
//func ListenAndServer(addr string, handler http.Handler) error
//该方法第一个参数为http服务器监听的tcp网络地址，如： 127.0.0.1:8080
//第二个参数未服务端处理函数名，通常为空:nil
//则服务端会调用http.DefaultServeMux函数进行处理
//服务端编写的业务逻辑处理函数(类似于web框架中的路由)http.Handle()或http.HandleFunc()默认会注入http.DefaultServeMux中

//还可以开启一个https服务端：
// func ListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) error


//url “/hello”对应的视图函数
func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world"))
}

func main(){
	//指定url的路由
	http.HandleFunc("/hello", hello)

	//运行http server
	http.ListenAndServe("127.0.0.1:8080", nil)
}
