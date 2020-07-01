package main

import (
	"database/sql" //这是一个sql接口包，对所有的数据库都可以用这个接口操作，但是需要相应的数据库驱动包
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
//需要安装数据库相应的驱动 go get github.com/go-sql-driver/mysql
func main(){
	//连接数据库的格式如下： mysql是驱动，root是用户名，123是密码， tcp是连接方式，test是db名，？后面是其他参数，此处是字符集
	db, err := sql.Open("mysql","root:emc123123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	//要提前建好数据库并且设置好字段
/*
	// 增
	stmt, err := db.Prepare("INSERT user_info SET username=?, department=?, create_time=?")
	_, err = stmt.Exec("eponia","tech","2020-10-01 00:00:00")
	if err != nil{
		panic(err)
	}

	//删
	stmt, err = db.Prepare("DELETE FROM user_info WHERE id=?")
	res, err := stmt.Exec("1")
	rows, err := res.RowsAffected()          //修改的记录条数， 可用于判断是否执行成功
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	// 改
	stmt, err = db.Prepare("UPDATE user_info SET department=? WHERE username=?")
	res, err := stmt.Exec("finance","eponia")
	rows, err := res.RowsAffected()          //修改的记录条数， 可用于判断是否执行成功
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)


 */
	// 查
	rows, err := db.Query("select * from user_info")
	if err != nil {
		panic(err)
	}
	for rows.Next(){
		var id int
		var username string
		var department string
		var create_time string
		err = rows.Scan(&id, &username, &department, &create_time) //用&符号，表示对这几个内存地址进行操作，此处是把扫描到的值存到相应的地址，从而对这几个变量fuz
		fmt.Println(id,username,department,create_time)
	}

}