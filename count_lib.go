package lib1

import (
	"fmt"
	"math"
	"strings"
)

//数值超高精度计算

//Count_Add
func Count_Add(Real string,Close string) string {

	//修正数据 除数不能为 0
	if Parse_float(Close)==0 {
		return "0"
	}

	Price:=(Parse_float(Real)-Parse_float(Close))/Parse_float(Close)*100

	return Conversion_Num(Price,6)
}

//专用涨跌占比计算
func Count_Amo_Add(Real string,Close string) string {

	//修正数据 除数不能为 0
	if Parse_float(Close)==0 {
		return "0"
	}

	Price:=(Parse_float(Real)-Parse_float(Close))/Parse_float(Close)

	return Conversion_Num(Price,2)
}

func Count_Add_float(Real float64,Close float64) float64 {

	//修正数据 除数不能为 0
	if (Close)==0 {
		return 0
	}

	Price:=((Real)-(Close))/(Close)*100

	return Price
}

//使用相似计算 10%
func Count_SQL_Similar(data_time,Type,High,Lower,Open,Close string) string {

	//sql_str="SELECT ABS(ROUND("+Data_kline_List[i].Close+"*1000-"+Data_kline_List[i].Low+"*1000))<1000 AS number"
	//Num=com.Mysql_Query_Number(sql_str,com.MySQL_global)

	//修正数据 除数不能为 0
	if Parse_float(Open)==0||Parse_float(Open)==0 {
		return "0"
	}



	var price float64
	var proportion float64
	var wave float64

	//精确定位时间点
	if strings.Contains(data_time, "11-15 19:30:00")==true {
		fmt.Println(data_time)
	}

	//使用内部比例自计算
	//生成波幅
	wave=math.Abs(Parse_float(High)-Parse_float(Lower))

	//生成基准比例
	// 5% 比例
	proportion=(wave/100)*15

	switch Type {

	case "max":
		//使用超高精度计算 千分之一
		//proportion=(lib1.Parse_float(High)/10000)*1
		price=math.Abs(Parse_float(High)-Parse_float(Close))  //如果数值相差 5%

		//超过最大比例值
		if price<proportion {
			return "1"
		}


	case "min":

		//proportion=(lib1.Parse_float(Lower)/10000)*1
		price=math.Abs(Parse_float(Lower)-Parse_float(Close))

		//波动值小于比例
		if price<proportion {
			return "1"
		}


	}

	return "0"
}

//豆包代码 两个数字相差是否在 5% 以内
func IsWithinFivePercent(a, b float64) bool {
    diff := math.Abs(a - b)
    maxVal := math.Max(math.Abs(a), math.Abs(b))
    return diff/maxVal < 0.05
}

// 计算两个浮点数之间的距离
func DistanceBetweenTwoNumbers(a, b float64) float64 {
    return math.Abs(a - b)
}

// CalculateDistance 计算两个浮点数之间的绝对距离 Grok3
func CalculateDistance(float1, float2 float64) float64 {
    return math.Abs(float1 - float2)
}