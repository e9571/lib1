// mysql_lib
package lib1

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//1 连接配置模块 使用全局模式
/*
func Database_str(name string) string {

	config := make(map[string]string)

	//填充 MySQL 应用设置
	config["volume_base_analyse_log"] = "admin:123@tcp(192.168.1.105:3306)/volume_base_analyse_log?charset=utf8"
	//config["volume_base_analyse_log"] = "root:@tcp(127.0.0.1:3306)/volume_base_analyse_log?charset=utf8"

	return config[name]

}
*/




//2 SQL连接输出模块 动态泛型
// 2017年10月3日10:27:01
// 修改细节 即使返回 捕获所有异常
func Mysql_connect_query_config_api(sql_str string, database_config string) (map[int]map[string]string,error) {

	db, err := sql.Open("mysql", database_config)

	result := make(map[int]map[string]string)

	//捕获所有异常
	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		db.Close()
		return result,err
	}

	//查询数据，取所有字段
	rows2, err := db.Query(sql_str)

	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		rows2.Close()
		db.Close()
		return result,err
	}

	//返回所有列
	cols, err := rows2.Columns()

	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		rows2.Close()
		db.Close()
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

	// 关闭 连接

	  db.Close()

	return result,err
}


//原版模式 获得插入返回ID

// 2017年10月3日10:27:01
// 修改细节 即使返回 捕获所有异常
func Mysql_connect_query_insert_api(sql_str string, database_config string) (int64,error) {

	db, err := sql.Open("mysql", database_config)

	var ins_id int64

	ins_id=-1

	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		db.Close()
		return ins_id,err
	}
	//db.Exec(sql_str)
	ret, err := db.Exec(sql_str)

	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		db.Close()
		return  ins_id,err

	}


	ins_id, err = ret.LastInsertId()

	//ins_id := 0

	//fmt.Println(ins_id)

	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		db.Close()
		return  ins_id,err

	}


	db.Close()

	return ins_id,err

}

//新版模式 不需要返回ID 只要不报错就可以

// 2017年10月3日10:27:01
// 修改细节 即使返回 捕获所有异常 返回成功 失败
func Mysql_connect_query_insert_api_3(sql_str string, database_config string)  bool{

	db, err := sql.Open("mysql", database_config)


	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		db.Close()
		return  false
	}

	_, err = db.Exec(sql_str)

	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		db.Close()
		return  false

	}

	 db.Close()

	return true
}

//无返回查询 更新

// 2017年10月3日10:27:01
// 修改细节 即使返回 捕获所有异常 获得影响行数
func Mysql_connect_query_update_api(sql_str string, database_config string) (int64,error) {

	db, err := sql.Open("mysql", database_config)

	var  ins_id  int64

	ins_id=-1

	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		db.Close()
		return  ins_id,err
	}
	ret, err := db.Exec(sql_str)

	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		db.Close()
		return  ins_id,err
	}

	ins_id, err = ret.RowsAffected()

	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		db.Close()
		return  ins_id,err
	}

	db.Close()

	return ins_id,err
}

//返回单条数据 number 标准封装

func Mysql_query_data(sql_str string, database_config string) (string,error) {

	db, err := sql.Open("mysql", database_config)
	number := ""
	if err != nil {
		fmt.Println("message:"+sql_str)
		fmt.Println(err)
		db.Close()
		return number,err
	}

	//fmt.Println(sql_str)

	rows3 := db.QueryRow(sql_str)
	rows3.Scan(&number)

	//fmt.Println(number)
	 db.Close()

	return number,err

}

//Update 专用函数( golang 版本 ) 更新指定字段 基于 Base_id
//lib1.Update_Field(base_id,field,value,,table_name,database)
func Update_Field (base_id string,field string,value string,table_name string,database string)  {

	sql_str:="UPDATE `"+table_name+"` SET  `"+field+"` = '"+value+"' WHERE `id` = '"+base_id+"'"

	Mysql_connect_query_insert_api_3(sql_str,database)

	Printf("Update:"+base_id+" field:"+field+" value:"+value)
	
}


//3 SQL语句工具类

//1 生成插入语句

func Assemble_insert(data map[string]string, table string) string {

	var sql_str = " insert into " + table + " ("
	//var sql_str1 string

	//对数据进行排序

	sorted_keys1 := Taxis_map(data)

	//填充字段

	for _, k := range sorted_keys1 {
		// fmt.Printf("k=%v, v=%v\n", k, m[k])
		sql_str += k + ","
	}

	//获取指定位数
	sql_str = sql_str[0 : len(sql_str)-1]

	sql_str += " ) values ( "


	//使用排序模式
	for _, k := range sorted_keys1 {
		sql_str += "'" + data[k] + "',"
	}

	sql_str = sql_str[0 : len(sql_str)-1]

	sql_str += " ) "

	return sql_str

}

//2 生成更新语句

func Assemble_update(data map[string]string, where map[string]string, table string) string {

	var sql_str = " update " + table + " set "
	//var sql_str1 string

	//填充字段
	for key, value := range data {
		sql_str += " " + key + " = '" + value + "',"
	}

	//获取指定位数
	sql_str = sql_str[0 : len(sql_str)-1]

	//强制 where
	sql_str += " where "

	//填充 Value
	for key1, value1 := range where {
		sql_str += " " + key1 + "='" + value1 + "'" + " and "
	}

	sql_str = sql_str[0 : len(sql_str)-4]

	return sql_str

}

//3 生成防止重复插入语句
func Assemble_insert_exists(data map[string]string, where map[string]string, table string) string {

	var sql_str = " insert into " + table + " ("

	//对 map 进行排序 确保每次 Key => Value 对应

	sorted_keys1 := Taxis_map(data)


	for _, k := range sorted_keys1 {
		// fmt.Printf("k=%v, v=%v\n", k, m[k])
		sql_str += k + ","
	}

	sql_str = sql_str[0 : len(sql_str)-1]

	//合成 Select 部分

	sql_str += " ) select "

	//生成 Value 部分

	for _, k := range sorted_keys1 {
		// fmt.Printf("k=%v, v=%v\n", k, m[k])
		sql_str += "'" + data[k] + "',"
	}

	sql_str = sql_str[0 : len(sql_str)-1]

	//生成判定选择部分

	sql_str += " from " + table + " where not exists ( select "

	//再次生成 Key 部分

	for key, _ := range data {
		sql_str += key + ","
	}

	sql_str = sql_str[0 : len(sql_str)-1]

	//生成判定选择部分

	sql_str += " from " + table + " where  "

	for key, value := range where {
		sql_str += key + "='" + value + "' and "
	}

	sql_str = sql_str[0 : len(sql_str)-4]

	sql_str += " ) LIMIT 1 "

	return sql_str

}


//MySQltool mode




//create table field list
//Table_field_list("room_info","mj_database")
func Table_field_list(table_str string,database_name string,database_str string)([]string,error){

	//sql_str:="SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.Columns WHERE table_name='room_info' AND table_schema='mj_database'"
	sql_str:="SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.Columns WHERE table_name='"+table_str+"' and table_schema='"+database_name+"'"

	var table_list []string

	result,err:=Mysql_connect_query_config_api(sql_str,database_str)

	//fmt.Println(data_source)

	//create table list

	if err!=nil {
		return table_list,err
	}



	for _, value := range result {
		table_list = append(table_list, value["COLUMN_NAME"])
	}

	return  table_list,err

}

//映射对应的数据库 单值模式
//sql_str="SELECT * FROM `room_info` LIMIT 0, 1"
//fmt.Println(Table_field_map(sql_str,"room_info","mj_database"))
func Table_field_map(sql_str string,table_str string,database_name string)map[string]string{

	//进行进行数数据库查询
	//获取指定字段

	result,_:=Mysql_connect_query_config_api(sql_str,database_name)

	// data must make
	data_source := make(map[string]string)

	//get data
	for _, value := range result {
		//fmt.Printf("%s->%s", key, value["id"])
		for key1, value1 := range value {
			//fmt.Printf("%s-->%s \n", key1, value1)
			data_source[key1] = value1
		}

		break
	}

	//fmt.Println(data_source)

	return  data_source

}