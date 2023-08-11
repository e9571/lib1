package lib1

import 	"sort"

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