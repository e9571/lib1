package lib1

import (
	"sort"
	"strings"
)

//字符串处理专用函数 处理复杂字符串需求

/*
    示例：地址排序

	var data []string

	data= append(data,"0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	data= append(data,"0x4f1960E29b2cA581a38c5c474e123f420F8092db")


	list:=lib.SortPackage(data,0)

	fmt.Println(list)

 */

//定义序列类型，只要此类型实现了sort.interface{}接口(实现三要素方法)，就可以对此类型排序
type StringList []string

//元素个数
func (t StringList) Len() int {
	return len(t)
}

//比较结果
func (t StringList) Less(i, j int) bool {
	return t[i] < t[j]
}

//交换方式
func (t StringList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

//排序模式
func SortPackage(data []string,Type int) []string{

	var list StringList

	for i:=0;i<len(data);i++ {
		list= append(list,data[i])
	}

	//升序
	if Type==0 {
		//按照定义的规则排序
		sort.Sort(list)
	}else {
		//降序
		sort.Sort(sort.Reverse(list))
	}

	return list

}

//字符串模式
func SortPackageStr(data []string,Type int) (string,string){

	list:=SortPackage(data,Type)

	str_tmp:=""

	for i:=0;i<len(list);i++ {
		str_tmp+=list[i]+"@"
	}

	str_tmp = str_tmp[0 : len(str_tmp)-1]

    return str_tmp,DefaultEncodeMD5(str_tmp)
}

//无差别对比字符串 返回符合数量
func Str_Count(data1,data2 string) int {
	return 	strings.Count(strings.ToLower(data1), strings.ToLower(data2))
}
