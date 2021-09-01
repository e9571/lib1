package lib1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//版本 1.1

//<<<<<<< Updated upstream
//修复 时间->时间戳 误差问题 使用高精度 时区模式

var Path_exe string

func init() {

   //文件补丁 程序运行目录
   	Path_exe, _ = Create_path_os()

}

//2016年4月7日13:12:03

//=======
//>>>>>>> Stashed changes
//Go 语言时间专用格式化 函数 返回 常规时间 返回文件名时间 返回 unix 时间
//time
//flie_time
//unix
//获取常规时间
//lib1.Create_Format_time("time") 
//获取帮助文件
///lib1.Create_Format_time("help") 
/*
数值样式参考
help:time,flie_time,msec,micro,nano,unix,unix_micro,unix_msec,unix_nano,time_str,msec_str,micro_str,nano_str
2021-9-1 11:55:08.822
2021-9-1 11:55:08.823065
2021-9-1 11:55:08.823065300
20210901115508
202191115508823
202191115508823065
202191115508823065300
1630468508
1630468508823
1630468508823065
1630468508824063000

*/
//更新 2021年9月1日11:35:16
func Create_Format_time(type_str string) string {

	tNow := time.Now()
	timestamp := tNow.Unix()
	
	//时间格式化参数 官方 文档
	/*
	const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
    )
    */
	

	str_time := ""

	switch type_str {
	case "time":
		str_time = time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	//新增参数 毫秒
    case "msec":
		str_time = tNow.Format("2006-1-2 15:04:05.000")
	//新增参数 微秒
    case "micro":
		str_time = tNow.Format("2006-1-2 15:04:05.000000")
	//新增参数 纳秒
   case "nano":
		str_time = tNow.Format("2006-1-2 15:04:05.000000000")
   //Unix 模式
	case "unix":
	   str_time = fmt.Sprintf("%d", timestamp)
	case "unix_micro":
	   str_time =fmt.Sprintf("%d", time.Now().UnixNano() /  1000000)
	case "unix_msec":
		str_time = fmt.Sprintf("%v", time.Now().UnixNano() /1000)
	case "unix_nano":
		str_time = fmt.Sprintf("%v", time.Now().UnixNano())
  //Str模式
	case "time_str":
		str_time = time.Unix(timestamp, 0).Format("20060102150405")
    case "msec_str":
		str_time = tNow.Format("200612150405.000")
		str_time = strings.Replace(str_time, ".", "", -1)
    case "micro_str":
		str_time = tNow.Format("200612150405.000000")
		str_time = strings.Replace(str_time, ".", "", -1)
   case "nano_str":
		str_time = tNow.Format("200612150405.000000000")
		str_time = strings.Replace(str_time, ".", "", -1)
	default:
		str_time="help:time,msec,micro,nano,unix,unix_micro,unix_msec,unix_nano,time_str,msec_str,micro_str,nano_str"

	}

	return str_time

}

//Go 语言 专用格式化 指定时间格式 转为  unix 时间
//记得拷贝高精度时间转换包到文件夹下 2020年2月1日18:29:18
func Get_appoint_number(time_str string) string {

  if len(Path_exe)==0{
      return "-1"
  }

	os.Setenv("ZONEINFO", Path_exe+"/conf/data.zip")
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
