package lib1

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
    "strconv"
	"regexp"
	"sort"
	"strings"
	"time"
	"path/filepath"
)

//专业函数工具类

// 1  最小化加载
// 2  新增多项应用函数
// 3  跨平台

//2019年1月22日10:35:46

//函数列表

/*

专业数据工具类 用于快速处理数据

//用于数据分词
Word_Split
//用于变量传递过滤，防止 SQL注入
Sql_filtrate
//用于检测数组中是否有该项数据 //PHP移植
In_array
//用于获取高速随机字符串
High_rand
//用于获取高速随机字符串，可加前缀
High_Rand_User
//正则表达式 提取指定特征
Get_data_preg_str
//提取数字
Get_data_preg_int
//提取浮点数
Get_data_preg_float
//提取比特币地址
Get_data_preg_address
//生成 MD5
DefaultEncodeMD5
//生成特定 CRC 校验值
Create_CRC

新增函数

//HTTP访问函数
Get_HTTP
//快速 int 转换函数
Parse_int
//快速 int8 转换函数
Parse_int8
//快速 int64 转换函数
Parse_int64
//快速 float 转换函数
Parse_float
//正则表达式 List 模式 稳定返回
Get_data_preg_list
//正则表达式 动态匹配兼容 整数 浮点数
Create_number
//格式化错误
SprintfErr
//获取本机IP
Create_ip

*/


//高速随机数

func High_rand(table string) string {

	//var key_id string
	t := time.Now()

	key_id := table + "_" + fmt.Sprintf("%d", t.Unix()) + "_" + fmt.Sprintf("%d", rand.Intn(1000000))

	return key_id

}

//用于生成指定范围内的随机数
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

//U
func High_Rand_User(prefix string) string {

	//var key_id string
	t := time.Now()

	key_id :=  prefix+fmt.Sprintf("%d", t.Unix()) + "_" + fmt.Sprintf("%d", RandInt64(1000000,9999999))

	return key_id

}

//超精简 HTTP Get 访问
func Get_HTTP(url_str string) (string,error) {

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Get(url_str)

	//对数据进行延迟关闭
	defer resp.Body.Close()

	if err!=nil {
		return "",err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "",err
	}

	return string(body),err
}


//数值快速转换

//通用 String 转 int 自动处理异常
func Parse_int( value string) int {

	int,err:=strconv.Atoi(value)

	if err != nil {
		return -1
	}

	return int
}

//转换成 int8 类型第一种方法
func Parse_int8( value string) int8 {

	value_tmp,err:=strconv.ParseInt(value, 10, 8)

	if err != nil {
		return -1
	}

	return int8(value_tmp)
}


//通用 String 转 int64 自动处理异常
func Parse_int64( value string) int64 {

	int64, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return -1
	}

	return int64
}

//通用 String 转 float 自动处理异常
func Parse_float( value string) float64 {

	float,err := strconv.ParseFloat(value,32/64)

	if err != nil {
		return -1
	}

	return float
}



//正则表达式
//应用参考
//regexp_str = `{[.\s\S]*?}`
//regexp_str = `[+-]?(\d+)`
func Get_data_preg(parameter string, source string) map[int]string {

	reg := regexp.MustCompile(parameter)

	line := reg.FindAllString(source, -1)

	result := make(map[int]string)

	for i := 0; i < len(line); i++ {

		result[i] = line[i]
	}

	return result
}

//正则表达式 List 版本 顺序排列
func Get_data_preg_list(parameter string, source string) []string {

	reg := regexp.MustCompile(parameter)


	line := reg.FindAllString(source, -1)

	//使用动态数组
	 var result []string

	for i := 0; i < len(line); i++ {

		result = append(result, line[i])
	}

	return result
}

//增强输出 返回包含指定特征的数据块 和正则表达式配合 只返回第一个符合的
//应用参考：
//height=lib1.Get_data_preg_number(lib1.Get_data_preg_str(result,"ga-type"))
func Get_data_preg_str(parameter map[int]string, sign string) string {

	data:=""

	for _, value := range parameter {

		if strings.Contains(value, sign)==true{

			data=value
			break

		}

	}

	return data
}

//返回字符串中的数字 数据类型只匹配第一个 只能是整数
func Get_data_preg_int(parameter string) string {

	data:=""

	regexp_str := `[+-]?(\d+)`
	result:=Get_data_preg(regexp_str,parameter)

	data=result[0]

	return data
}

//浮点数匹配
func Get_data_preg_float(parameter string) string {

	data:=""

	//^[\\+\\-]?[\\d]+(\\.[\\d]+)?$
	regexp_str := `[0-9]+.[0-9]+`

	//测试新版数据过滤
	regexp_str = `[\\+\\-]?[0-9]+.[0-9]+`
	result:=Get_data_preg(regexp_str,parameter)

	data=result[0]

	return data
}

//动态匹配兼容 整数 浮点数
func Create_number(source string) string{

	//先匹配整数
	regexp_str := `[+-]?(\d+)`

	result1:=Get_data_preg(regexp_str,source)

	if len(result1)==0 {
		return ""
	}

	//检测到多个整数
	if  len(result1)>1{

		//启用浮点数匹配

		regexp_str := `[\+\-]?[0-9]+.[0-9]+`

		result1=Get_data_preg(regexp_str,source)

		return result1[0]

	}

	return result1[0]
}

//提取所有英文字符 比特币地址专用
func Get_data_preg_address(parameter string) string {

	data:=""
	regexp_str := `[A-Za-z\d]+`
	result:=Get_data_preg(regexp_str,parameter)

	data=result[0]

	return data
}


//超级定位 先根据特征搜索指定字符串 之后截取指定字符串后面指定长度的字符

//适用于 特殊字符中无法使用正则表达式定位的时候

//使用方法  使用前  需要先确认过获取到的源码中有目标字符串 2018年5月7日09:25:53

//获取指定偏移后的指定长度字符串
//regexp_str = `Confirmations[.\s\S]*?</dd>` //整体区域
//height=lib1.Get_data_preg_search_str(html_str,regexp_str,"</dd>",30)//字符串结束位置 数据长度

func Get_data_preg_search_str(parameter string,regexp_str string,sign string,number int) string {

	//截取指定字符串

	//使用正则表达式 多维配合

	result:=Get_data_preg(regexp_str,parameter)
	value:=Get_data_preg_str(result,sign)

	parameter=value[0:number]

	return parameter

}


//打印指定颜色 输出


func Printf(str string) {

	//设置界面颜色
	//加载颜色包 go get github.com/fatih/color
	//项目主页 https://github.com/fatih/color

	//使用亮绿色
	color.HiGreen(str)

}


//排序 Map 防止不不稳定的情况出现

func Taxis_map(m map[string]string) []string {

	sorted_keys := make([]string, 0)
	for k, _ := range m {
		sorted_keys = append(sorted_keys, k)
	}

	// sort 'string' key in increasing order
	sort.Strings(sorted_keys)

	return sorted_keys

}

//通用错误检查
func CheckErr(err error) {
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(-1)
	}
}

//手动格式化错误 用于远程传递
func SprintfErr(err error)  string{

	if err==nil {
		return ""
	}

	return fmt.Sprintf("%s", err.Error())
}

//手动格式化 向 error 传输数据
func SprintfErrWrite(str string) error {

	var err error = errors.New(str)
	return  err
}

//手动格式化错误 用于远程传递
func SprintfErrMSG(err error)  string{

	if err==nil {
		return ""
	}

	return fmt.Sprintf("%s", err.Error())
}


// MD5 数据计算
func DefaultEncodeMD5(data_source string) string {

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data_source))
	cipherStr := md5Ctx.Sum(nil)

	return strings.ToUpper(hex.EncodeToString(cipherStr))

}

//应用数据分词

func Word_Split(data_source string, data_sign string) []string {

	return strings.Split(data_source, data_sign)

}


//最新编码处理库
type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)


//生成 CRC
//crc_tmp :=[3]string{"", "", ""}
//生成 CRC
//crc_tmp[0]=sql_str_value["stock_id"]
//crc_tmp[1]=sql_str_value["type_name"] //使用板块名称
//crc_tmp[2]=date_str
func Create_CRC(data_source [3]string) string {

	var str_tmp string

	for i := 0; i < len(data_source); i++ {
		str_tmp+=data_source[i]
	}

	str_tmp=DefaultEncodeMD5(str_tmp)

	return  str_tmp
}

//SQL专业过滤 来源 变形金刚框架 该框架的数据防御被阿里安全检测
func Sql_filtrate(str string) string{

	str = strings.Replace(str, ")", "", -1)//替换括号
	str = strings.Replace(str, "\"", "", -1)//替换双引号
	str = strings.Replace(str, "'", "", -1)//替换单引号

	return str
}

//遍历数组 判断重复 PHP 函数移植

/*
应用参考
if len(value1)>60 &&in_array(intlist_tmp,value1)==false{
}
*/

func In_array(array_list []string,value string) bool{

	for i := 0; i < len(array_list); i++ {

		if array_list[i]==value{
			return true
		}
	}

	return false
}

//通用数字转换
//原始数据  josn 偏移获取 num类型 转换长度
//conversion_Num(gjson.Get(data_source, "data."+strconv.Itoa(i) +".open").Num,10)
func  Conversion_Num(number float64,number_place int)  string{

	return 	strconv.FormatFloat(number, 'f', number_place, 32)

}


//获取本机IP
func Create_ip() string {

	ip:="0"

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ip
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				ip=ipnet.IP.String()
				break
			}
		}
	}

	return ip

}

//获取启动参数
func Create_Args()  []string{

	var Args_list []string

	//获取启动参数
	for idx, args := range os.Args {
		fmt.Println("启动参数" + strconv.Itoa(idx) + ":", args)

		Args_list=append(Args_list,args)
	}

	return Args_list

}

//获取程序文件名
func Create_exe_name() string {

	//自动获取程序 名称和路径
	exePath := os.Args[0]
	base := filepath.Base(exePath)
	suffix := filepath.Ext(exePath)
	return strings.TrimSuffix(base, suffix)
}

//使用内核调用 可执行文件路径 2019年2月22日9:55:02
func Create_path_os()  (string,error){

	ex, err := os.Executable()

	if err != nil {
		//进行异常提示
		fmt.Println(err)
		fmt.Scanln()
		return "",err
	}

	exPath := filepath.Dir(ex)

	return exPath+"",err
}