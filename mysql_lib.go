
package lib1

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*

MySQL 快速连接模块 Go语言版本

该版本与 C#，PHP 的 mysql_lib 用法一致

非ORM模式下的源生连接


by 9571 china  xi'an

95714623@qq.com

2018年2月28日16:40:59


主要功能

1 查询数据，列表输出
Mysql_connect_query_config_api

2 插入数据，获得返回ID
Mysql_connect_query_insert_api

3 插入数据，不需要返回ID
Mysql_connect_query_insert_api_Not_ID

4 查询数据，之返回单条结果 比如 Select count(id) as number from log
Mysql_connect_query_update_api

5 生成标准 SQL 插入语句，传入一个MAP和Table，key->value 模式，传入后会返回一个 标准 SQL，插入语句，可直接执行
Assemble_insert

6 生成标准的 MySQL 防止重复插入语句，传入一个MAP，之后传入 用于判断是否重复的MAP，之后传入 Table，函数会生成一个标准的防止重复插入 SQL语言，可直接执行
Assemble_insert_exists

7 生成标准 的 Assemble_update 语句 传入value MAP 之后传入 Where MAP 之后传入 Table 函数会生成一个 Update 语言，可直接执行

8 获取制定数据表结构
Table_field_list

附录：MAP 专业排序 防止 SQL数据生成时不稳定的情况

//排序 Map 防止不不稳定的情况出现


func Taxis_map(m map[string]string) []string {

	sorted_keys := make([]string, 0)
	for k, _ := range m {
		sorted_keys = append(sorted_keys, k)
	}
	
	sort.Strings(sorted_keys)

	return sorted_keys

}


*/

//1 连接配置模块 参考
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
func Mysql_connect_query_config_api(sql_str string, database_config string) map[int]map[string]string {

	db, err := sql.Open("mysql", database_config)
	if err != nil {
		fmt.Println("message")
		fmt.Println(err)
	}

	rows2, _ := db.Query(sql_str)
    cols, _ := rows2.Columns()
    vals := make([][]byte, len(cols))

	scans := make([]interface{}, len(cols))
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		rows2.Scan(scans...)
		row := make(map[string]string)
		for k, v := range vals {
			key := cols[k]
			row[key] = string(v)
		}
		result[i] = row
		i++
	}

	// close connect  
	defer db.Close()

	return result
}


//原版模式 获得插入返回ID
func Mysql_connect_query_insert_api(sql_str string, database_config string) int64 {

	db, err := sql.Open("mysql", database_config)

	var ins_id int64

	if err != nil {
		fmt.Println("message")
		fmt.Println(err)
	}

	ret, _ := db.Exec(sql_str)
	ins_id=0
	ins_id, _ = ret.LastInsertId()

	defer db.Close()

	return ins_id

}

//新版模式 不需要返回ID 
func Mysql_connect_query_insert_api_Not_ID(sql_str string, database_config string)  {

	db, err := sql.Open("mysql", database_config)

	if err != nil {
		fmt.Println("message")
		fmt.Println(err)
	}
	db.Exec(sql_str)

	defer db.Close()

}

//无返回查询 更新
func Mysql_connect_query_update_api(sql_str string, database_config string) int64 {

	db, err := sql.Open("mysql", database_config)
	if err != nil {
		fmt.Println("message")
		fmt.Println(err)
	}
	ret, _ := db.Exec(sql_str)

	ins_id, _ := ret.RowsAffected()

	defer db.Close()

	return ins_id

}

//返回单条数据 number 标准封装
func Mysql_query_data(sql_str string, database_config string) string {

	db, err := sql.Open("mysql", database_config)
	if err != nil {
		fmt.Println("message")
		fmt.Println(err)
	}

	number := ""
	rows3 := db.QueryRow(sql_str)
	rows3.Scan(&number)

	defer db.Close()

	return number

}

//3 SQL语句工具类

//1 生成插入语句

func Assemble_insert(data map[string]string, table string) string {

	var sql_str = " insert into " + table + " ("

	//对数据进行排序
	sorted_keys1 := Taxis_map(data)

	for _, k := range sorted_keys1 {
		sql_str += k + ","
	}

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

	//填充字段
	for key, value := range data {
		sql_str += " " + key + " = '" + value + "',"
	}

	sql_str = sql_str[0 : len(sql_str)-1]

	//强制 where
	sql_str += " where "


	for key1, value1 := range where {
		sql_str += " " + key1 + "='" + value1 + "'" + " and "
	}

	sql_str = sql_str[0 : len(sql_str)-4]

	//sql_str += " ) "

	return sql_str

}

//3 生成防止重复插入语句
func Assemble_insert_exists(data map[string]string, where map[string]string, table string) string {

	var sql_str = " insert into " + table + " ("



	sorted_keys1 := Taxis_map(data)

	//填充字段 生成 Key 部分


	for _, k := range sorted_keys1 {
		// fmt.Printf("k=%v, v=%v\n", k, m[k])
		sql_str += k + ","
	}

	sql_str = sql_str[0 : len(sql_str)-1]


	sql_str += " ) select "



	for _, k := range sorted_keys1 {
		sql_str += "'" + data[k] + "',"
	}

	sql_str = sql_str[0 : len(sql_str)-1]

	sql_str += " from " + table + " where not exists ( select "

	for key, _ := range data {
		sql_str += key + ","
	}

	sql_str = sql_str[0 : len(sql_str)-1]

	sql_str += " from " + table + " where  "

	for key, value := range where {
		sql_str += key + "='" + value + "' and "
	}

	sql_str = sql_str[0 : len(sql_str)-4]

	sql_str += " ) LIMIT 1 "

	return sql_str

}


//MySQl Tool mode


//使用方法
//Table_field_list("room_info","mj_database")
func Table_field_list(table_str string,database_name string)[]string{


	sql_str:="SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.Columns WHERE table_name='"+table_str+"' and table_schema='"+database_name+"'"

	result:=Mysql_connect_query_config_api(sql_str,Database_str_global(database_name))

	var table_list []string

	for _, value := range result {
		table_list = append(table_list, value["COLUMN_NAME"])
	}

	return  table_list

}

//映射对应的数据库 单值模式
//sql_str="SELECT * FROM `room_info` LIMIT 0, 1"
//fmt.Println(Table_field_map(sql_str,"room_info","mj_database"))
func Table_field_map(sql_str string,table_str string,database_name string)map[string]string{

	//get database map process odd number mode
	result:=Mysql_connect_query_config_api(sql_str,Database_str_global(database_name) )

	data_source := make(map[string]string)

	for _, value := range result {

		for key1, value1 := range value {
			data_source[key1] = value1
		}

		break
	}
	return  data_source
}