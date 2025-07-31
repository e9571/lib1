package lib1

import (
	"golang.org/x/crypto/sha3"
	"math/big"
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
//在1 中搜索 2 出现次数
func Str_Count(data1,data2 string) int {
	return 	strings.Count(strings.ToLower(data1), strings.ToLower(data2))
}

//无差别字符串数组 int 类型对比 第一个是 , 好数组类型 第二个是数字
func Str_Count_int(data1,data2 string) int {

	list:=Word_Split(data1,",")

	for i:=0;i<len(list);i++ {
		if Parse_int(list[i])==Parse_int(data2) {
			return 1
		}
	}

	return 0
}


/*

    //智能合约 哈希数值函数模拟
   
	// 示例用法
	addr := "0x1234567890abcdef1234567890abcdef12345678" // 模拟以太坊地址
	typeValue := "abcd"                                     // 任意字符串
	result := hashAddressNTokenId(addr, typeValue)
	fmt.Printf("哈希结果 (uint64): %d\n", result)

	// 更多测试用例
	fmt.Println("测试用例 1:", hashAddressNTokenId(addr, "xyz"))
	fmt.Println("测试用例 2:", hashAddressNTokenId(addr, "1234"))
	fmt.Println("测试用例 3:", hashAddressNTokenId("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd", "test"))


*/

func HashAddressNTokenId(addr string, typeValue string) uint64 {
	// 模拟 abi.encodePacked：直接拼接 addr 和 typeValue 的字节
	data := []byte(addr + typeValue)

	// 使用 Keccak256 哈希
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	hashBytes := hash.Sum(nil)

	// 将前 8 字节转换为 uint64（模拟 Solidity uint）
	result := new(big.Int).SetBytes(hashBytes[:8]).Uint64()

	return result
}

//快速编码模式 ASCII 之和 返回值与 HashAddressNTokenId 兼容
func Str_ascii(str string)  uint64{

	var sum int
	for _, char := range str {
		sum += int(char) // 将每个字符的 ASCII 值相加
	}
	//fmt.Println("字符 ASCII 值之和:", sum)
	return uint64(sum)
}