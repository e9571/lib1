// mysql_lib
package lib1

import (
	"database/sql"
	"fmt"
	"strconv"

	//驱动安装  go get github.com/mattn/go-sqlite3
	//https://github.com/mattn/go-sqlite3
	//https://studygolang.com/articles/5456
	_ "github.com/mattn/go-sqlite3"
)

// sqlite3 敏捷操作版本
// 版本 1.1


//版本 1.2

//修正文件访问问题

//2019年2月20日10:28:44




//升级应用 直接返回底层错误
func SQLite_connect_query_config_api(sql_str string, dbPath string) (map[int]map[string]string,error) {

	if CheckFileIsExist(dbPath)==false{
		return nil,SprintfErrWrite("File not find......")
	}

	db, err := sql.Open("sqlite3", dbPath)

	result := make(map[int]map[string]string)

	if err != nil {
		if db!=nil {
			db.Close()
		}
		return result,err
	}

	rows2, _ := db.Query(sql_str)

	//查询数据，取所有字段
	//rows2, _ := db.Query(sql_str)

	//返回所有列
	cols, err := rows2.Columns()

	if  err!=nil{
		return result,err
	}

	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))

	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))

	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	//fmt.Println(scans)

	//使用 Key value 模式获取数据
	i := 0

	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}

	//fmt.Println(result)

	// 延迟关闭 连接
	//defer db.Close()
	db.Close()

	return result,err
}

//无返回查询 插入
func SQLite_connect_query_insert_api(sql_str string, dbPath string) (int64,error) {

	//修复 SQLite 官方 BUG ，在查询文件的时候先检查文件是否存在
	if CheckFileIsExist(dbPath)==false{
		return -1,SprintfErrWrite("File not find......")
	}

	db, err := sql.Open("sqlite3", dbPath)
	var ins_id int64

	if err != nil {

		if db!=nil {
			db.Close()
		}

		return ins_id,err
	}
	//db.Exec(sql_str)
	ret, err := db.Exec(sql_str)

	if err != nil {
		db.Close()
		return ins_id,err
	}

	ins_id, err = ret.LastInsertId()

	if err != nil {
		db.Close()
		return ins_id,err
	}

	//ins_id := 0

	//fmt.Println(ins_id)

	//defer db.Close()
	db.Close()

	return ins_id,err

}

//无返回查询 更新
func SQLite_connect_query_update_api(sql_str string, dbPath string) (int64,error) {

	//数据补丁 检查文件是否存在
	if CheckFileIsExist(dbPath)==false{
		return -1,SprintfErrWrite("File not find......")
	}

	db, err := sql.Open("sqlite3", dbPath)

	var ins_id int64

	if err != nil {

		if db!=nil {
			db.Close()
		}

		return ins_id,err
	}

	ret, err := db.Exec(sql_str)

	if err != nil {
		db.Close()
		return ins_id,err
	}

	ins_id, err = ret.RowsAffected()

	if err != nil {
		db.Close()
		return ins_id,err
	}

	//defer db.Close()
	db.Close()

	return ins_id,err

}

//返回单条数据 number 标准封装
func SQLite_query_data(sql_str string, dbPath string) (string,error) {

	//dbPath="D:/db.db"

	//数据补丁 检查文件是否存在
	if CheckFileIsExist(dbPath)==false{
		return "",SprintfErrWrite(dbPath+" File not find......")
	}

	db, err := sql.Open("sqlite3", dbPath)

	number := ""

	if err != nil {
		if db!=nil {
			db.Close()
		}
		return number,err
	}

	//fmt.Println(sql_str)

	rows3 := db.QueryRow(sql_str)
	rows3.Scan(&number)

	//fmt.Println(number)
	//defer db.Close()
	db.Close()

	return number,err

}

//缓存插入执行 超高速缓存模式
func SQLite_connect_query_insert_api_buffer(sql_str []string, dbPath string) (string,bool) {

	if CheckFileIsExist(dbPath)==false{
		return ("File not find......"),false
	}

	db, err := sql.Open("sqlite3", dbPath)
	var ins_id int64
    var  msg string

	if err != nil {

		if db!=nil {
			db.Close()
		}

		return SprintfErrMSG(err),false
	}
	//db.Exec(sql_str)
	msg=""

	number:=len(sql_str)

	var number_write int

	number_write=0

	for i := 0; i < len(sql_str); i++ {
		ret, err := db.Exec(sql_str[i])

		if err != nil {
			msg+=SprintfErrMSG(err)
			msg+=","
			fmt.Println("insert_buffer_err:"+SprintfErrMSG(err))
			continue
		}

		number_write++
		ins_id, err = ret.LastInsertId()
		fmt.Println("write:"+strconv.FormatInt(ins_id,10))
	}


	//ins_id := 0

	//fmt.Println(ins_id)

	//defer db.Close()
	db.Close()

	if number==number_write {

		fmt.Println("数据写入 完成 总数:"+strconv.Itoa(number)+" 实际写入数:"+strconv.Itoa(number_write))
		return "",true
	}

	fmt.Println("数据写入 CRC 异常 总数:"+strconv.Itoa(number)+" 实际写入数:"+strconv.Itoa(number_write))
	return msg,false

}
