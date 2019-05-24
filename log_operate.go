// mysql_lib
package lib1

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"strings"
)

//版本 1.1

//新增静态日志文件定位 支持服务器环境

//2017年12月23日11:31:37

//日志快速写入
//Create_log
//动态生成最新日志
//Create_new_log
//遍历文件夹 同时检查数据文件是否超时
//GetFileInfo_list_log
//遍历文件夹 获取文件列表
//GetFilelist
//删除指定时间范围文件名
//Delete_time_area_file

//日志快速写入 内容 文件路径
func Create_log(data_write string) {

	log_path := "log/write_event.log"

	logfile, err := os.OpenFile(log_path,  os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		//os.Exit(-1)
	}
	//写入日志信息
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

	logger.Println(data_write)

}

/*

//Golang 原始文件返回结构参考
const (
	Ldate         = 1 << iota     //日期示例： 2009/01/23
	Ltime                         //时间示例: 01:23:23
	Lmicroseconds                 //毫秒示例: 01:23:23.123123.
	Llongfile                     //绝对路径和行号: /a/b/c/d.go:23
	Lshortfile                    //文件和行号: d.go:23.
	LUTC                          //日期时间转为0时区的
	LstdFlags     = Ldate | Ltime //Go提供的标准抬头信息
)
*/

//增量写入模式

// 需要先获得创建文件目录
// log_flie=lib1.Create_new_log("log")

func Create_log_add(data_write string,file_name string) {

	log_path := "log/write_event.log"

	//如果文件名有效
	if len(file_name)>0 {

		log_path=file_name
	}

	logfile, err :=  os.OpenFile(log_path,  os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
	}

	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

	logger.Println(data_write)

}



//使用方法
//static_path="D:\\soft\\esb\\sacn_163_radar\\"

//获取最新日志文件名 默认当前目录
//log_flie=lib1.Create_new_log(static_path,"log")

//默认本地模式
//log_flie=lib1.Create_new_log("","log")

//获得最新日志 如果模块需要 返回文件名 使用静态路径
func Create_new_log(static_path string,type_str string) string {

	//根据时间线 生成日志
	fileName :=static_path+ "log/" +type_str+"_"+ Create_Format_time("flie_time")[0:10] + ".log"

	//如果未定义路径
	if len(static_path)==0{
		fileName = "log/" +type_str+"_"+ Create_Format_time("flie_time")[0:10] + ".log"
	}


	//检查文件是否存在
	if Exists(fileName)==true{
		return  fileName
	}

	logFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()

	return fileName
}

//如果使用静态文件位置 2018年7月12日14:20:20

//使用静态文件路径
func Create_log_add_static(static_path string,data_write string,file_name string) {
	log_path := static_path+"log/write_event.log"
	//如果文件名有效
	if len(file_name)>0 {
		log_path=file_name
	}
	logfile, err :=  os.OpenFile(log_path,  os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	logger := log.New(logfile, "", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(data_write)
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Log_operate_check() {

	//获取当前路径
	path := GetCurrentPath()
	//当前路径操作
	root := path + "/log"
	fmt.Println("Log_operate:" + root)
	//获得路径列表
	path_list := GetFilelist(root)
	fmt.Println(root)
	//fmt.Println(path_list)
	//进行遍历操作
	day_area := 7
	GetFileInfo_list_log(path_list, path, day_area)

}

//遍历 获取时间 同时进行数据检查 Log 日志专用
func GetFileInfo_list_log(data_source []string, path_source string, day_area int) {
	for i := 0; i < len(data_source); i++ {

		_, err := os.Stat(data_source[i])

		if err != nil {
			fmt.Println(err)
		}

		//如果大于指定时间 进行文件删除操作
		Delete_time_area_file(data_source[i], path_source, day_area)

	}
}

//删除指定时间范围的 文件 重命名模式

func Delete_time_area_file(path_str string, path_source string, day_area int) {

	//打开文件
	fileinfo, err := os.Stat(path_str)
	
	if err != nil {

		fmt.Println(err)
	}

	file_name := fileinfo.Name()

	t := time.Now()

	// 直接数值类型 格式化后 转int


	//获取时间差 进行重命名
	if TIMESTAMPDIFF(t.Unix(), fileinfo.ModTime().Unix(), "DAY") > day_area {

		//根据规则重命名文件
		newPath := path_source + "/log/" + "delete_" + file_name

		err := os.Rename(path_str, newPath)

		if err != nil {
			//如果重命名文件失败,则输出错误 file rename Error!
			fmt.Println("file rename Error!")
			//打印错误详细信息
			fmt.Printf("%s", err)

		} else {
			//如果文件重命名成功,则输出 file rename OK!
			fmt.Println("file rename OK!")
		}
	}

}



//获取文件路径 日志使用
func GetCurrentPath_log() string {
	file, _ := exec.LookPath(os.Args[0])
	//fmt.Println("file:", file)
	path, _ := filepath.Abs(file)
	//fmt.Println("path:", path)
	splitstring := strings.Split(path, "\\")
	size := len(splitstring)
	splitstring = strings.Split(path, splitstring[size-1])
	ret := strings.Replace(splitstring[0], "\\", "/", size-1)
	return ret
}
