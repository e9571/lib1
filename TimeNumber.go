package lib1

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"sort"
	//1.6 版本以上才可用
	//_ "time/tzdata" // Embed timezone data for portability, especially on Windows
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
//dir 新增文件目录模式
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
	//文件目录模式
	case "dir":
		str_time = time.Unix(timestamp, 0).Format("20060102")
	default:
		str_time="help:time,msec,micro,nano,unix,unix_micro,unix_msec,unix_nano,time_str,msec_str,micro_str,nano_str"

	}

	return str_time

}

//Go 语言 专用格式化 指定时间格式 转为  unix 时间
//记得拷贝高精度时间转换包到文件夹下 2020年2月1日18:29:18
func Unix_number(time_str string) string {

  if len(Path_exe)==0{
      return "-1"
  }

  //转换补丁
	time_str = strings.Replace(time_str, "/", "-", -1)

	os.Setenv("ZONEINFO", Path_exe+"/conf/data.zip")
	loc, _ := time.LoadLocation("Asia/Shanghai")        //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", time_str, loc)

	return fmt.Sprintf("%d", tt.Unix())
	
}

//把 unix 时间按转为 指定时间
func Unix_time(unix_str string) string {

	if len(unix_str)!=10 {
		return unix_str
	}

	int_value, _ := strconv.ParseInt(unix_str, 10, 64)

	return time.Unix(int_value, 0).Format("2006-01-02 15:04:05")

}

//获取指定时间与当前时间间隔的秒数

func Uinx_interval(timeStr string) float64 {

	location := "Asia/Shanghai"

	timestamp_last, _ := ConvertToUnixTimestamp(timeStr, location)

	timestamp_real, _ := ConvertToUnixTimestamp(Create_Format_time("time"), location)

	//获取绝对值
	return math.Abs(float64(timestamp_real - timestamp_last))
}

// ConvertToUnixTimestamp 将时间字符串转换为 Unix 时间戳（秒）
// 输入格式必须为 "2006-01-02 15:04:05" (YYYY-MM-DD HH:MM:SS)
// location 为时区，例如 "Asia/Shanghai" 或 "UTC"
func ConvertToUnixTimestamp(timeStr, location string) (int64, error) {
	// 加载时区
	loc, err := time.LoadLocation(location)
	if err != nil {
		return 0, err
	}

	// 使用内置的时间格式解析
	const layout = "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(layout, timeStr, loc)
	if err != nil {
		return 0, err
	}

	// 返回 Unix 时间戳（秒）
	return t.Unix(), nil
}


//新增时间快捷计算函数

//获取币安标准时间 返回 unix 还是 time
func Time_standard(unix string,Type string) string {

	if len(unix)!=13 {
		return ""
	}

	unix = unix[0 : len(unix)-3]

	if Type=="unix" {
		return unix
	}
	return  Unix_time(unix)
}

//获取指定日期是周几 亚洲时间
//fmt.Println(models.Week(lib1.Create_Format_time("time")))
func Week_day(timeStr string)string  {

	layout := "2006-01-02 15:04:05"
	utcTime, _ := time.Parse(layout, timeStr)

	Week:=int(utcTime.Local().Weekday())-1

	//日期补丁
	if Week==-1 {
		Week=6
	}

	return strconv.Itoa(Week)
}



//生成高精度 group_id 秒级
//天级数据 Group_id_sec("")[0:8]
//小时级数据 Group_id_sec("")[0 :10]
//分钟级数据 lib1.Group_id_sec("")[ :12 ]
func Group_id_sec(symbol string) string{
	group_id:=Create_Format_time("time")
	group_id=strings.Replace(group_id, " ", "", -1)
	group_id=strings.Replace(group_id, ":", "", -1)
	group_id=strings.Replace(group_id, "-", "", -1)

	if len(symbol)>0 {
		return symbol+"_"+group_id
	}

	return group_id
}

//格式化指定时间，生成数据简写与时序标志位(返回 标准时序 天 小时 分钟)
func Create_time_id(timeStr string) (string,string,string,string){

	layout := "2006-01-02 15:04:05" //常规时间标志位
	layout_day:= "20060102" //天标志 年+月+日
	layout_hour:= "0215" //小时标志 天+小时
	layout_minute:= "1504" //分钟标志 小时+分钟

	utcTime, _ := time.Parse(layout, timeStr)

	time_sign := utcTime.Format("20060102150405")

	date_day := utcTime.Format(layout_day)
	date_hour := utcTime.Format(layout_hour)
	date_minute := utcTime.Format(layout_minute)

	return time_sign,date_day,date_hour,date_minute
}


//Grok3
// ConvertTimeFormat 将时间字符串从 "20060102150405" 格式转换为 "2006-01-02 15:04:05"
//时间解码
func Data_sign_decode(input string) (string) {
    // 解析输入时间字符串
    t, err := time.Parse("20060102150405", input)
    if err != nil {
    	fmt.Println("invalid time format, expected YYYYMMDDHHMMSS")
        return ""
    }
    
    // 转换为目标格式
    return t.Format("2006-01-02 15:04:05")
}

/*
	// 示例输入时间字符串
	inputs := []string{
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"2006.01.02 15:04:05",
		"2006-01-02T15:04:05",
		"Mon, 02 Jan 2006 15:04:05 MST",
	}

	// 循环调用 ConvertToCompactTimeFormat 函数
	for _, input := range inputs {
		result, err := ConvertToCompactTimeFormat(input)
		if err != nil {
			fmt.Printf("Error converting '%s': %v\n", input, err)
			continue
		}
		fmt.Printf("Input: %s -> Output: %s\n", input, result)
	}

*/

// ConvertToCompactTimeFormat 将任意格式的时间字符串转换为 "20060102150405" 格式
//时间编码
func Data_sign_encode(input string) (string) {
	// 定义常见的时间格式
	formats := []string{
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"2006.01.02 15:04:05",
		"20060102 15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02",
		time.RFC3339,
		time.RFC1123,
		time.RFC822,
	}

	var parsedTime time.Time
	var err error

	// 尝试用每种格式解析输入字符串
	for _, format := range formats {
		parsedTime, err = time.Parse(format, input)
		if err == nil {
			break
		}
	}

	if err != nil {
		fmt.Println("unable to parse time string with any known format")
		return ""
	}

	// 转换为目标格式
	return parsedTime.Format("20060102150405")
}



//获取时间戳中的指定标志位
// y m d h i s
/*
	//时间专用函数 获取时间戳中的 年 月 日 秒
    //Test
	value:=lib1.Create_Format_time("time_str")
	fmt.Println(value)
	fmt.Println("y",Get_time_str(value1,"y"))
	fmt.Println("m",Get_time_str(value1,"m"))
	fmt.Println("d",Get_time_str(value1,"d"))
	fmt.Println("h",Get_time_str(value1,"h"))
	fmt.Println("i",Get_time_str(value1,"i"))
	fmt.Println("s",Get_time_str(value1,"s"))
 */
func Get_time_str(value,Type string)  string{

	//判断长度是否是标准时间
	if len(value)!=19 {
		fmt.Println("Length_err",len(value),value)
		return value
	}

	value = strings.Replace(value, "/", "-", -1)

	//生成标准时间 ID
	value,_,_,_=Create_time_id(value)

	if len(value)!=14 {
		fmt.Println("Length_err",len(value),value)
		return value
	}

	Type=strings.ToLower(Type)

	//20210901121606

	switch Type {
	//年
	case "y":
		value=value[0:4]
	//月
	case "m":
		value=value[4:6]
	//日
	case "d":
		value=value[6:8]
	//小时
	case "h":
		value=value[8:10]
	//分
	case "i":
		value=value[10:12]
	//秒
	case "s":
		value=value[12:14]
	default:
	}

    return value

}

//随机数函数

//快速随机数 普通生成
// random函数用于生成一个min到max之间的随机整数
func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

//专业随机数 纳秒生成
// GenerateRandomNumber 生成指定范围内的随机整数
// 参数:
// - min: 范围最小值 (包含)
// - max: 范围最大值 (包含)
// 返回: 随机整数和可能的错误
func GenerateRandomNumber(min, max int) (int, error) {
	// 验证输入参数
	if min > max {
		return 0, fmt.Errorf("最小值 %d 不能大于最大值 %d", min, max)
	}

	// 使用当前纳秒时间戳作为随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成 [min, max] 范围内的随机数
	// rand.Intn(n) 生成 [0, n-1] 的随机数，因此调整为 [min, max]
	return min + rand.Intn(max-min+1), nil
}

//高精度 时序 ID
func Res_ID(Type string) string {

	myrand := Random(10000, 99999)

	//返回指定固定长度的数据
	return Create_Format_time("msec_str")+fmt.Sprintf("%d", myrand)+Type
}

// 支持的常见时间格式列表
var supportedFormats = []string{
	"2006-01-02 15:04:05",       // 标准格式：2026-01-07 11:00:00
	"2006/01/02 15:04:05",       // 斜杠分隔：2026/01/07 11:00:00
	"2006.01.02 15:04:05",       // 点分隔：2026.01.07 11:00:00
	"20060102150405",           // 紧凑无分隔符：20260107110000
	"2006-01-02T15:04:05",       // ISO 带 T：2026-01-07T11:00:00
	"2006-01-02",               // 仅日期：2026-01-07
	time.RFC3339,               // 2026-01-07T11:00:00Z07:00
	time.RFC3339Nano,           // 带纳秒：2026-01-07T11:00:00.123456789Z
	time.RFC1123,               // Wed, 07 Jan 2026 11:00:00 MST
	time.RFC1123Z,              // Wed, 07 Jan 2026 11:00:00 +0800
	time.RFC822,                // 07 Jan 26 11:00 MST
	time.RFC822Z,               // 07 Jan 26 11:00 +0800
	"2006-01-02 15:04:05.000",  // 带毫秒
	"2006-01-02T15:04:05Z",     // UTC 时间
	"2006-01-02T15:04:05-07:00",// 带时区偏移
}

//高精度时间判断

/*

//判断目标时间是否大于当前时间制定分钟

		if lib1.IsBeforeMinutes(handle.Data[i].Data_time,3)==false {
			continue
		}

*/

// IsKlineOutdated 判断 Kline 时间是否已经超过当前时间指定分钟数
// 用于检测数据是否太旧（例如：最新 K 线时间距离现在 > 3 分钟 → 认为已过期）
//
// 参数:
//   timeStr - Kline 的时间字符串（一定是过去的时间）
//   minutes - 过期阈值（例如 3 分钟）
//
// 返回:
//   true  - 已经超过指定分钟（数据太旧，需要更新）
//   false - 还在阈值内（数据较新鲜）
func IsKlineOutdated(timeStr string, minutes int) bool {
	var parsedTime time.Time
	var err error

	// 关键修改：使用 time.Local 时区解析
	location := time.Local

	for _, layout := range supportedFormats {
		parsedTime, err = time.ParseInLocation(layout, timeStr, location)
		if err == nil {
			break
		}
	}

	if err != nil {
		fmt.Printf("时间解析失败: %s (错误: %v)，视为已过期\n", timeStr, err)
		return true // 解析失败，保守策略：认为已过期
	}

	now := time.Now()

	elapsed := now.Sub(parsedTime)

	// 可选：打印调试信息（上线前可删除）
	// fmt.Printf("Kline时间: %v | 当前时间: %v | 过去了: %v\n", parsedTime, now, elapsed)

	return elapsed > time.Duration(minutes)*time.Minute
}

// FormatToYMDHM 将任意支持的时间字符串格式化为 "YYYYMMDDHHMM" 格式
// 示例：2026-01-07 11:05:30 → "202601071105"
//
// 参数:
//   timeStr - 任意常见格式的时间字符串
//
// 返回:
//   string - 格式化后的 "年月日时分" 字符串（长度固定为12位）
//   bool   - 是否解析成功（失败返回 false 和空字符串）
func FormatToYMDHM(timeStr string) (string, bool) {
	var parsedTime time.Time
	var err error

	// 依次尝试所有支持的格式进行解析
	for _, layout := range supportedFormats {
		parsedTime, err = time.Parse(layout, timeStr)
		if err == nil {
			break
		}
	}

	if err != nil {
		// 如果所有格式都解析失败
		return "", false
	}

	// 使用固定布局格式化：年(4位) 月(2位) 日(2位) 时(2位) 分(2位)
	// 200601021504 → YYYYMMDDHHMM
	formatted := parsedTime.Format("200601021504")

	return formatted, true
}

//高精度提取日期格式 全自动匹配符合日期的格式

// DateStrToYYYYSlashMMDD 把日期字符串转为 2025/12/19 格式
// 规则：直接取字符串的前8位作为 YYYYMMDD 进行解析
// 支持：20251219、20251219000、202512190000、20251219000000 等
// 失败条件：前8位不是有效日期，或者长度小于8位
func DateStrToYYYYSlashMMDD(s string) (string, bool) {
    s = strings.TrimSpace(s)
    
    // 长度不够8位，直接失败
    if len(s) < 8 {
        return "", false
    }
    
    // 只取前8位
    datePart := s[:8]
    
    t, err := time.Parse("20060102", datePart)
    if err != nil {
        return "", false
    }
    
    return t.Format("2006/01/02"), true
}

// 提取 时间字段进行去重 YYYYMMDD
// 更保守一点的写法（如果非常在意内存分配次数）
func ExtractUniqueSortedDatesConservative(dates []string) []string {
	seen := make(map[string]bool)

	for _, s := range dates {
		if len(s) >= 8 {
			seen[s[:8]] = true
		}
	}

	result := make([]string, 0, len(seen))
	for day := range seen {
		result = append(result, day)
	}

	sort.Strings(result)
	return result
}

// GetMondayDates 过滤出列表中是星期一的日期（格式 YYYYMMDD）
func GetMondayDates(dates []string) []string {
	var mondays []string

	for _, dateStr := range dates {
		if len(dateStr) != 8 {
			continue // 格式不对就跳过
		}

		// 解析 YYYYMMDD 格式
		t, err := time.Parse("20060102", dateStr)
		if err != nil {
			continue // 解析失败就跳过
		}

		// time.Weekday()：0=Sunday, 1=Monday, ..., 6=Saturday
		if t.Weekday() == time.Monday {
			mondays = append(mondays, dateStr)
		}
	}

	return mondays
}


//获取整点秒数
//lib1.SecondsToNextHour(lib1.Create_Format_time("time"))
// SecondsToNextHour 計算距離下一個整點還有多少秒
//   - 無參數 → 使用當前時間
//   - 有參數 → 嘗試解析該時間字串
// 回傳值：永遠只回傳 int (秒數)，錯誤或提示會印在 stdout
func SecondsToNextHour(timeStr ...string) int {
	var now time.Time

	if len(timeStr) == 0 {
		now = time.Now()
	} else {
		input := strings.TrimSpace(timeStr[0])
		if input == "" {
			now = time.Now()
		} else {
			var parsed bool
			for _, layout := range supportedFormats {
				t, err := time.ParseInLocation(layout, input, time.Local)
				if err == nil {
					now = t
					parsed = true
					break
				}
			}
			if !parsed {
				fmt.Printf("錯誤：無法解析時間字串 %q\n", input)
				fmt.Println("  支援的格式範例：")
				fmt.Println("    2026-01-19 15:20:30")
				fmt.Println("    20260119152030")
				fmt.Println("    2026-01-19T15:20:30+08:00")
				fmt.Println("使用當前時間作為 fallback")
				now = time.Now()
			}
		}
	}

	// 核心計算
	_, min, sec := now.Clock()
	secondsPast := min*60 + sec
	secondsLeft := 3600 - secondsPast

	// 只在函數內印出資訊（供除錯或觀察用）
	fmt.Printf("[debug] 基準時間: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("[debug] 距離下一個整點: %d 秒\n", secondsLeft)

	return secondsLeft
}

// IsTimeInRange 判断 target 是否在 start ~ end 范围内（包含边界）
// 所有参数都必须是 14 位数字字符串：yyyymdhhiiss
// 示例：20250315143022 → 2025-03-15 14:30:22
func IsTimeInRange(target, start, end string) (bool, error) {
	const layout = "20060102150405"

	// 长度校验（最简单但很有效的防御）
	if len(target) != 14 || len(start) != 14 || len(end) != 14 {
		return false, fmt.Errorf("时间字符串必须是14位数字，实际长度: target=%d, start=%d, end=%d",
			len(target), len(start), len(end))
	}

	t, err := time.Parse(layout, target)
	if err != nil {
		return false, fmt.Errorf("目标时间格式错误: %s (%v)", target, err)
	}

	s, err := time.Parse(layout, start)
	if err != nil {
		return false, fmt.Errorf("开始时间格式错误: %s (%v)", start, err)
	}

	e, err := time.Parse(layout, end)
	if err != nil {
		return false, fmt.Errorf("结束时间格式错误: %s (%v)", end, err)
	}

	// start <= target <= end
	return !t.Before(s) && !t.After(e), nil
}

//判断是否是周末
// IsWeekend 检查传入的 "2006-01-02 15:04:05" 格式时间字符串（支持单/双位数字）对应的本地日期是否为周末
// 返回：(是否周末, 错误信息)
// 核心特性：
//   ✅ 使用 ParseInLocation + time.Local：将字符串直接解释为系统本地时间（非 UTC 转换）
//   ✅ 布局字符串 "2006-1-2 15:04:05"：完美兼容单/双位月、日、时、分、秒
//   ✅ 明确错误提示：包含无效格式示例
func IsWeekend(datetimeStr string) (bool, error) {
	layout := "2006-1-2 15:04:05" // Go 参考时间布局：1/2 表示月/日支持单双位，15 表示 24 小时制

	// 关键：使用 ParseInLocation 将字符串直接解析为系统本地时区的时间
	// 避免先解析为 UTC 再转换导致的日期偏移（例如：UTC 23:00 可能已是本地次日）
	t, err := time.ParseInLocation(layout, datetimeStr, time.Local)
	if err != nil {
		return false, fmt.Errorf(
			"时间解析失败（格式应为: yyyy-m-d hh:mm:ss，例如 2026-2-19 10:29:30 或 2026-02-19 10:29:30）: %w",
			err,
		)
	}

	// 检查星期（此时 t 已是本地时区，无需额外转换）
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday, nil
}


//高精度时间排序格式化

// TimeFormat 定义支持的时间格式及其是否包含时间部分
type TimeFormat struct {
	Pattern    string // 解析用的格式
	HasTime    bool   // 是否包含时分秒
	OutputFmt  string // 输出用的格式（用于 Format）
}

var supportedFormats_v2 = []TimeFormat{
	{"2006-01-02 15:04:05", true, "20060102150405"},
	{"2006/01/02 15:04:05", true, "20060102150405"},
	{"2006.01.02 15:04:05", true, "20060102150405"},
	{"20060102150405", true, "20060102150405"},
	{"2006-01-02T15:04:05", true, "20060102150405"},
	{time.RFC3339, true, "20060102150405"},
	{time.RFC3339Nano, true, "20060102150405"},
	{time.RFC1123, true, "20060102150405"},
	{time.RFC1123Z, true, "20060102150405"},
	{time.RFC822, true, "20060102150405"},
	{time.RFC822Z, true, "20060102150405"},
	{"2006-01-02 15:04:05.000", true, "20060102150405"},
	{"2006-01-02T15:04:05Z", true, "20060102150405"},
	{"2006-01-02T15:04:05-07:00", true, "20060102150405"},
	{"2006-01-02", false, "20060102"}, // 仅日期
}

type TimeSlice []time.Time

func (t TimeSlice) Len() int           { return len(t) }
func (t TimeSlice) Less(i, j int) bool { return t[i].Before(t[j]) }
func (t TimeSlice) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

// FormatDateList 升级版：根据输入格式类型决定输出格式 排序输出
func FormatDateList(dateStrs []string) ([]string, error) {
	if len(dateStrs) == 0 {
		return []string{}, nil
	}

	// 尝试匹配格式 高精度时序 v2 版本
	var detected *TimeFormat
	for _, tf := range supportedFormats_v2 {
		if _, err := time.Parse(tf.Pattern, dateStrs[0]); err == nil {
			detected = &tf
			break
		}
	}

	if detected == nil {
		return nil, fmt.Errorf("first date '%s' does not match any supported format", dateStrs[0])
	}

	// 验证所有字符串都符合该格式
	times := make([]time.Time, len(dateStrs))
	for i, s := range dateStrs {
		t, err := time.Parse(detected.Pattern, s)
		if err != nil {
			return nil, fmt.Errorf("'%s' doesn't match format '%s': %v", s, detected.Pattern, err)
		}
		times[i] = t
	}

	// 排序（Go 1.5 兼容）
	sort.Sort(TimeSlice(times))

	// 按 detected.OutputFmt 格式化输出
	result := make([]string, len(times))
	for i, t := range times {
		result[i] = t.Format(detected.OutputFmt)
	}

	return result, nil
}