package lib1

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

	logfile.Close()

}



//使用方法
//static_path="D:\\soft\\esb\\sacn_163_radar\\"

//获取最新日志文件名 默认当前目录
//log_flie=lib1.Create_new_log(static_path,"log")

//默认本地模式
//log_flie=lib1.Create_new_log("","log")

//获得最新日志 如果模块需要 返回文件名 使用静态路径
func Create_new_log(static_path string,type_str string) (string,error) {



	var err error
	var fileName string



	//如果未定义路径 使用程序目录生成路径
	if len(static_path)==0{
		//path_exe, _ := Create_path_os()
		//fileName = path_exe+"/"+"log/" +type_str+"_"+ Create_Format_time("flie_time")[0:10] + ".log"

		path_exe, _ := Create_path_os()

		//创建文件夹

		fileName = path_exe+"/"+"log/"
		//检查文件是否存在
		fileName,err=Create_New_File(fileName)

		if err!=nil {
			return "", err
		}

		//创建文件
		fileName =fileName+"/" +type_str+"_"+ Create_Format_time("flie_time") + ".log"

		Write_file("initialize\n",fileName)

	}else {
		//根据时间线 生成日志
		fileName =static_path+ "log/" +type_str+"_"+ Create_Format_time("flie_time")[0:10] + ".log"
		//检查文件是否存在
		fileName,err=Create_New_File(fileName)
	}



	return fileName,err
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

	//logger := log.New(logfile, "", log.Ldate|log.Ltime|log.Llongfile)
	logger := log.New(logfile, "",log.Ldate|log.Ltime)
	logger.Println("data_body:\n",data_write)

	logfile.Close()
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
