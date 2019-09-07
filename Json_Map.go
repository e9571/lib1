package lib1


import (
    "encoding/json"
    "strconv"
    "strings"
)


//Json to Map and Map to Json 通用转换库

//将 Map 转换成 JSON
func Map_to_json(m map[string]interface{}) string {


   data, err := json.Marshal(m)

if err!=nil {
return  err.Error()
   }

return string(data)
}

//将Json 转换成 Map 只转换一级结构
func Json_to_map(json_str string ) (map[string]string,error){

var result map[string]interface{}

   value_str := make(map[string]string)

   err := json.Unmarshal([]byte(json_str), &result)

if err!=nil {
return value_str,err
   }

for key, value := range result {

switch Typeof_Json(value) {
case "string":
         value_str[key]=value.(string)
case "int":
         value_str[key]=strconv.Itoa(value.(int))
case "float64":
         value_str[key]= strconv.FormatFloat(value.(float64), 'f', -1, 64)
default:
         value_str[key]="0"
}

   }

return value_str,nil

}

//数据反射专用函数 精确解析数据类型
func Typeof_Json(v interface{}) string {

switch t := v.(type) {
case int:
return "int"
case string:
return "string"
case float64:
return "float64"
default:
      _ = t
return "unknown"
}

//return  fmt.Sprintf("%T", v)
}





//非标准 JSON 直接转换  适用于 IOS  无冒号类型 JSON 转换

//泛型转换

func Str_interface_to_json(source string) map[string]string{

//数据分词
    source = strings.Replace(source, "{", "", -1)
   source = strings.Replace(source, "}", "", -1)

   list_str:=lib1.Word_Split(source,",")


   result :=make(map[string]string)

var listStrTmp []string

//数据二次分析
for i:=0;i<len(list_str) ;i++  {

      listStrTmp =lib1.Word_Split(list_str[i],":")

if  len(listStrTmp)==2{
         result[listStrTmp[0]]= listStrTmp[1]
      }
   }

return result
}
