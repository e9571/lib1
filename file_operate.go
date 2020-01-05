package lib1

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"io"
	"fmt"
	"os/exec"
	"strings"
	"bufio"
)

//文件操作 1.1

//新增 文件写入专业操作类

//文件存在判断
//逐行读取
//逐行写入


//函数列表

//获得当前路径
//重命名文件
//遍历文件夹
//获取指定目录下 所有文件 支持子文件夹 可进行匹配过滤
//拷贝文件


//文件操作

//获得当前路径
func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	CheckErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

//重命名 文件
//源文件 目标文件
func File_Rename(source_file string, target_file string) {
	err := os.Rename(source_file, target_file)

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

//遍历文件夹 获取文件夹列表
//path="E:/Go_source/mqant-master"
//fmt.Println(lib1.GetFilelist(path))
func GetFilelist(path string) []string {
	var path_tmp []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)
		path_tmp = append(path_tmp, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	return path_tmp
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
//path_list,_:=WalkDir(path,".go")
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			//fmt.Println(fi.Name())
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

//拷贝文件
//源文件位置 拷贝位置
func CopyFile(srcName string, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

//读取文件 字节流模式
func ReadAll_Byte(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}


//文件写入专业操作类


//专业文件读写操作

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}


//新增操作函数

//逐行读取文件

func  Create_Source(path string) ([]string,error) {

	var Data_buffer []string

	fi, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return Data_buffer,err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)

	var line int

	line=0


	for {
		a, _, c := br.ReadLine()

		if c == io.EOF {
			break
		}

		//根据数据行缓冲数据
		Data_buffer=append(Data_buffer,string(a))

		line++

	}

	return Data_buffer,err

}

//写入文件 常规写入
func Write_file(str string,path string) (error)  {

	//var wireteString = str


	var filename = "./output1.txt"

	filename=path

	var f *os.File
	var err error

	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if CheckFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		//fmt.Println("文件存在")
	} else {
		f, err = os.Create(filename) //创建文件
		//fmt.Println("文件不存在")
	}

	if err!=nil {
		return  err
	}

	//高级文件读写
	data :=  []byte(str)

	_, err =f.Write(data)



	f.Sync()
	f.Close()

	return  err

}

//数据写入 精确行 支持换行符
func WriteListtoFile(List []string, filePath string) error {

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, v := range List {
		lineStr := fmt.Sprintf("%s", v)
		fmt.Fprintln(w, lineStr)
	}
	return w.Flush()
}

//Map 写入，支持换行符模式
func WriteMaptoFile(m map[string]string, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for k, v := range m {
		lineStr := fmt.Sprintf("%s^%s", k, v)
		fmt.Fprintln(w, lineStr)
	}
	return w.Flush()
}

//2020年1月5日10:00:59
//新建文件 自动判断文件夹是否存在
func Create_New_File(fileName string) (string,error) {

	//根据时间线 生成日志
	//fileName :=static_path+ "log/" +type_str+"_"+ Create_Format_time("flie_time")[0:10] + ".log"

	//如果未定义路径
	/*
	if len(static_path)==0{
		fileName = "log/" +type_str+"_"+ Create_Format_time("flie_time")[0:10] + ".log"
	}
	 */

	//检查文件是否存在
	if Exists(fileName)==true{
		return  fileName,nil
	}

	logFile, err := os.Create(fileName)

	if err != nil {
		//fmt.Println(err)
		return "",err
	}

	logFile.Close()

	return fileName,nil
}




