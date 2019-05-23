package lib1

import (

	"fmt"
	"strconv"
	"time"
)

//版本 1.1

//修复 时间->时间戳 误差问题 使用高精度 时区模式

//2019年4月7日13:12:03

//Go 语言时间专用格式化 函数 返回 常规时间 返回文件名时间 返回 unix 时间
//time
//flie_time
//unix
//lib1.Create_Format_time("time")
func Create_Format_time(type_str string) string {

	tNow := time.Now()
	timestamp := tNow.Unix()

	str_time := ""

	switch type_str {
	case "time":
		//只能是 2006-01-02 15:04:05 根据官方文档
		// http://blog.csdn.net/juxuny/article/details/43409983
		str_time = time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	case "flie_time":
		str_time = time.Unix(timestamp, 0).Format("2006_01_02_15_04_05")
	case "unix":
		str_time = fmt.Sprintf("%d", timestamp)

	}

	return str_time

}

//Go 语言 专用格式化 指定时间格式 转为  unix 时间
func Get_appoint_number(time_str string) string {

	loc, _ := time.LoadLocation("Asia/Shanghai")        //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", time_str, loc)

	return fmt.Sprintf("%d", tt.Unix())
	
}

//把 unix 时间按转为 指定时间
func Convert_appoint_number(unix_str string) string {

	int_value, _ := strconv.ParseInt(unix_str, 10, 64)

	return time.Unix(int_value, 0).Format("2006-01-02 15:04:05")

}

//时间类专用计算

//获取两个时间的时间差 传入 unix 数据 秒数
//SECOND
//MINUTE
//HOUR
//DAY 最多只到天，因为跨月的话 ，会涉及到的月度日期的不一致问题
func TIMESTAMPDIFF(datetime_expr_start int64, datetime_expr_complete int64, type_str string) int {

	//进行数据转换
	//v1, err := strconv.ParseFloat(datetime_expr1, 64)
	//v2, err := strconv.ParseFloat(datetime_expr2, 64)
	//value := math.Abs(v1 - v2)
	value := datetime_expr_start - datetime_expr_complete

	//进行数据转换 如果需要整除
	value_str := strconv.FormatInt(value, 10)
	value_float, _ := strconv.ParseFloat(value_str, 64)

	value_return := 1001

	switch type_str {
	case "SECOND":
		//string := strconv.FormatInt(int64, value_tmp)
		value_tmp := strconv.FormatInt(value, 16)
		//value, _ = strconv.Atoi(value_tmp)
		value_return, _ = strconv.Atoi(value_tmp)
		return value_return
	case "MINUTE":
		value_float = value_float / 60
	case "HOUR":
		value_float = value_float / 60 / 60
	case "DAY":
		value_float = value_float / 60 / 60 / 24
	}

	//对数据进行格式化过滤
	value_tmp := fmt.Sprintf("%0.0f", value_float)

	//重新对数据进行赋值
	value_return, _ = strconv.Atoi(value_tmp)

	return value_return

}
